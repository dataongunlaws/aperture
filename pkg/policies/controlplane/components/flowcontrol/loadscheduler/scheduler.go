package loadscheduler

import (
	"context"
	"fmt"
	"math"
	"path"
	"time"

	prometheusmodel "github.com/prometheus/common/model"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/fx"
	"go.uber.org/multierr"
	"google.golang.org/protobuf/proto"

	policylangv1 "github.com/fluxninja/aperture/api/gen/proto/go/aperture/policy/language/v1"
	policysyncv1 "github.com/fluxninja/aperture/api/gen/proto/go/aperture/policy/sync/v1"
	"github.com/fluxninja/aperture/pkg/config"
	etcdclient "github.com/fluxninja/aperture/pkg/etcd/client"
	etcdwriter "github.com/fluxninja/aperture/pkg/etcd/writer"
	"github.com/fluxninja/aperture/pkg/metrics"
	"github.com/fluxninja/aperture/pkg/notifiers"
	"github.com/fluxninja/aperture/pkg/policies/controlplane/components/query/promql"
	"github.com/fluxninja/aperture/pkg/policies/controlplane/iface"
	"github.com/fluxninja/aperture/pkg/policies/controlplane/runtime"
	"github.com/fluxninja/aperture/pkg/policies/paths"
)

var (
	tokenRateQueryInterval  = time.Second * 1
	autoTokensQueryInterval = time.Second * 10
)

// Scheduler is part of the LoadScheduler component stack.
type Scheduler struct {
	policyReadAPI       iface.Policy
	acceptedQuery       *promql.ScalarQuery
	incomingQuery       *promql.ScalarQuery
	tokensQuery         *promql.TaggedQuery
	tokensByWorkload    *policysyncv1.TokensDecision
	writer              *etcdwriter.Writer
	componentID         string
	autoTokensEtcdPaths []string
}

// Name implements runtime.Component.
func (*Scheduler) Name() string { return "Scheduler" }

// Type implements runtime.Component.
func (*Scheduler) Type() runtime.ComponentType { return runtime.ComponentTypeSource }

// ShortDescription implements runtime.Component.
func (s *Scheduler) ShortDescription() string {
	return fmt.Sprintf("%d agent groups", len(s.autoTokensEtcdPaths))
}

// IsActuator implements runtime.Component.
func (*Scheduler) IsActuator() bool { return false }

// NewSchedulerAndOptions creates scheduler and its fx options.
func NewSchedulerAndOptions(
	schedulerProto *policylangv1.LoadScheduler_Scheduler,
	componentID string,
	policyReadAPI iface.Policy,
	agentGroups []string,
) (runtime.Component, fx.Option, error) {
	var options []fx.Option

	var etcdPaths []string

	for _, agentGroup := range agentGroups {
		etcdPaths = append(etcdPaths, path.Join(paths.AutoTokenResultsPath,
			paths.AgentComponentKey(agentGroup, policyReadAPI.GetPolicyName(), componentID)))
	}

	scheduler := &Scheduler{
		policyReadAPI: policyReadAPI,
		tokensByWorkload: &policysyncv1.TokensDecision{
			TokensByWorkloadIndex: make(map[string]uint64),
		},
		componentID:         componentID,
		autoTokensEtcdPaths: etcdPaths,
	}
	options = append(options, fx.Invoke(scheduler.setupWriter))

	// Prepare parameters for prometheus queries
	policyParams := fmt.Sprintf("%s=\"%s\",%s=\"%s\",%s=\"%s\"",
		metrics.PolicyNameLabel,
		policyReadAPI.GetPolicyName(),
		metrics.PolicyHashLabel,
		policyReadAPI.GetPolicyHash(),
		metrics.ComponentIDLabel,
		componentID,
	)

	acceptedQuery, acceptedQueryOptions, acceptedQueryErr := promql.NewScalarQueryAndOptions(
		fmt.Sprintf("sum(rate(%s{%s}[10s]))",
			metrics.AcceptedTokensMetricName,
			policyParams),
		tokenRateQueryInterval,
		componentID,
		policyReadAPI,
		"AcceptedTokenRate",
	)
	if acceptedQueryErr != nil {
		return nil, fx.Options(), acceptedQueryErr
	}
	scheduler.acceptedQuery = acceptedQuery
	options = append(options, acceptedQueryOptions)

	incomingQuery, incomingQueryOptions, incomingQueryErr := promql.NewScalarQueryAndOptions(
		fmt.Sprintf("sum(rate(%s{%s}[10s]))",
			metrics.IncomingTokensMetricName,
			policyParams),
		tokenRateQueryInterval,
		componentID,
		policyReadAPI,
		"IncomingTokenRate",
	)
	if incomingQueryErr != nil {
		return nil, nil, incomingQueryErr
	}
	scheduler.incomingQuery = incomingQuery
	options = append(options, incomingQueryOptions)

	if schedulerProto.Parameters == nil {
		return nil, nil, fmt.Errorf("scheduler parameters are nil")
	}
	if schedulerProto.Parameters.AutoTokens {
		tokensQuery, tokensQueryOptions, tokensQueryErr := promql.NewTaggedQueryAndOptions(
			fmt.Sprintf("sum by (%s) (increase(%s{%s}[30m])) / sum by (%s) (increase(%s{%s}[30m]))",
				metrics.WorkloadIndexLabel,
				metrics.WorkloadLatencySumMetricName,
				policyParams,
				metrics.WorkloadIndexLabel,
				metrics.WorkloadLatencyCountMetricName,
				policyParams),
			autoTokensQueryInterval,
			componentID,
			policyReadAPI,
			"Tokens",
		)
		if tokensQueryErr != nil {
			return nil, nil, tokensQueryErr
		}
		scheduler.tokensQuery = tokensQuery
		options = append(options, tokensQueryOptions)
	}

	return scheduler, fx.Options(options...), nil
}

func (s *Scheduler) setupWriter(etcdClient *etcdclient.Client, lifecycle fx.Lifecycle) error {
	logger := s.policyReadAPI.GetStatusRegistry().GetLogger()
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			s.writer = etcdwriter.NewWriter(etcdClient, true)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			var merr error
			for _, etcdPath := range s.autoTokensEtcdPaths {
				_, err := etcdClient.KV.Delete(clientv3.WithRequireLeader(ctx), etcdPath)
				if err != nil {
					logger.Error().Err(err).Msg("Failed to delete tokens decision config")
					merr = multierr.Append(merr, err)
				}
			}
			s.writer.Close()
			return merr
		},
	})
	return nil
}

// Execute implements runtime.Component.Execute.
func (s *Scheduler) Execute(inPortReadings runtime.PortToReading, tickInfo runtime.TickInfo) (runtime.PortToReading, error) {
	logger := s.policyReadAPI.GetStatusRegistry().GetLogger()
	var errMulti error

	if s.tokensQuery != nil {
		taggedResult, err := s.tokensQuery.ExecuteTaggedQuery(tickInfo)
		promValue := taggedResult.Value
		if err != nil {
			if err != promql.ErrNoQueriesReturned {
				logger.Error().Err(err).Msg("could not read tokens query from prometheus")
				errMulti = multierr.Append(errMulti, err)
			}
		} else if promValue != nil {
			if vector, ok := promValue.(prometheusmodel.Vector); ok {
				tokensDecision := &policysyncv1.TokensDecision{
					TokensByWorkloadIndex: make(map[string]uint64),
				}
				for _, sample := range vector {
					for k, v := range sample.Metric {
						if k == metrics.WorkloadIndexLabel {
							// if sample.Value is NaN, continue
							if math.IsNaN(float64(sample.Value)) {
								continue
							}
							workloadIndex := string(v)
							sampleValue := uint64(sample.Value)
							tokensDecision.TokensByWorkloadIndex[workloadIndex] = sampleValue
							break
						}
					}
				}
				err = s.publishQueryTokens(tokensDecision)
				if err != nil {
					errMulti = multierr.Append(errMulti, err)
					logger.Error().Err(err).Msg("failed to publish tokens")
				}
			} else {
				err = fmt.Errorf("tokens query returned a non-vector value")
				errMulti = multierr.Append(errMulti, err)
				logger.Error().Err(err).Msg("Failed to parse tokens")
			}
		}
	}

	var acceptedReading, incomingReading runtime.Reading

	outPortReadings := make(runtime.PortToReading)

	acceptedScalarResult, err := s.acceptedQuery.ExecuteScalarQuery(tickInfo)
	acceptedValue := acceptedScalarResult.Value
	if err != nil {
		acceptedReading = runtime.InvalidReading()
		if err != promql.ErrNoQueriesReturned {
			errMulti = multierr.Append(errMulti, err)
		}
	} else {
		acceptedReading = runtime.NewReading(acceptedValue)
	}
	outPortReadings["accepted_token_rate"] = []runtime.Reading{acceptedReading}

	incomingScalarResult, err := s.incomingQuery.ExecuteScalarQuery(tickInfo)
	incomingValue := incomingScalarResult.Value
	if err != nil {
		incomingReading = runtime.InvalidReading()
		if err != promql.ErrNoQueriesReturned {
			errMulti = multierr.Append(errMulti, err)
		}
	} else {
		incomingReading = runtime.NewReading(incomingValue)
	}
	outPortReadings["incoming_token_rate"] = []runtime.Reading{incomingReading}

	return outPortReadings, errMulti
}

// DynamicConfigUpdate is a no-op for this component.
func (s *Scheduler) DynamicConfigUpdate(event notifiers.Event, unmarshaller config.Unmarshaller) {}

func (s *Scheduler) publishQueryTokens(tokens *policysyncv1.TokensDecision) error {
	logger := s.policyReadAPI.GetStatusRegistry().GetLogger()
	s.tokensByWorkload = tokens
	policyName := s.policyReadAPI.GetPolicyName()
	policyHash := s.policyReadAPI.GetPolicyHash()

	wrapper := &policysyncv1.TokensDecisionWrapper{
		TokensDecision: tokens,
		CommonAttributes: &policysyncv1.CommonAttributes{
			PolicyName:  policyName,
			PolicyHash:  policyHash,
			ComponentId: s.componentID,
		},
	}
	dat, err := proto.Marshal(wrapper)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to marshal tokens")
		return err
	}
	for _, etcdPath := range s.autoTokensEtcdPaths {
		s.writer.Write(etcdPath, dat)
	}
	return nil
}