local flow_selector_defaults = {
  service_selector: {
    service: '__REQUIRED_FIELD__',
  },
  flow_matcher: {
    control_point: '__REQUIRED_FIELD__',
  },
};

local promql_driver_defaults = {
  query_string: '__REQUIRED_FIELD__',
  forward: {
    threshold: '__REQUIRED_FIELD__',
    operator: '__REQUIRED_FIELD__',
  },
  backward: {
    threshold: '__REQUIRED_FIELD__',
    operator: '__REQUIRED_FIELD__',
  },
  reset: {
    threshold: '__REQUIRED_FIELD__',
    operator: '__REQUIRED_FIELD__',
  },
};

local average_latency_driver_defaults = {
  flow_selector: flow_selector_defaults,
  forward: {
    threshold: '__REQUIRED_FIELD__',
  },
  backward: {
    threshold: '__REQUIRED_FIELD__',
  },
  reset: {
    threshold: '__REQUIRED_FIELD__',
  },
};

local percentile_latency_driver_defaults = {
  flux_meter: {
    flow_selector: flow_selector_defaults,
    static_buckets: {
      buckets: [5.0, 10.0, 25.0, 50.0, 100.0, 250.0, 500.0, 1000.0, 2500.0, 5000.0, 10000.0],
    },
  },
  forward: {
    threshold: '__REQUIRED_FIELD__',
  },
  backward: {
    threshold: '__REQUIRED_FIELD__',
  },
  reset: {
    threshold: '__REQUIRED_FIELD__',
  },
  percentile: 95,
};

local ema_latency_driver_defaults = {
  flow_selector: flow_selector_defaults,
  ema: {
    ema_window: '1500s',
    warmup_window: '60s',
  },
  forward: {
    latency_tolerance_multiplier: 1.05,
  },
  backward: {
    latency_tolerance_multiplier: 1.05,
  },
  reset: {
    latency_tolerance_multiplier: 1.25,
  },
};

local rollout_policy_base_defaults = {

  load_shaper: {
    flow_regulator_parameters: {
      flow_selector: flow_selector_defaults,
      label_key: '',
    },
    steps: [
      {
        target_accept_percentage: '__REQUIRED_FIELD__',
        duration: '__REQUIRED_FIELD__',
      },
    ],
  },

  drivers: {},

  components: [],

  resources: {
    flow_control: {
      classifiers: [],
    },
  },

  evaluation_interval: '1s',
};

local rollout_policy_defaults = rollout_policy_base_defaults {

  drivers: {
    promql_drivers: [
      promql_driver_defaults,
    ],
    average_latency_drivers: [
      average_latency_driver_defaults,
    ],
    percentile_latency_drivers: [
      percentile_latency_driver_defaults,
    ],
    ema_latency_drivers: [
      ema_latency_driver_defaults,
    ],
  },

};


{
  /**
  * @param (common.policy_name: string required) Name of the policy.
  */
  common: {
    policy_name: '__REQUIRED_FIELD__',
  },
  /**
  * @param (policy: rollout_policy required) Parameters for the Feature Rollout policy.
  */
  policy: rollout_policy_base_defaults,
  /**
  * @schema (promql_driver.query_string: string required) The Prometheus query to be run. Must return a scalar or a vector with a single element.
  * @schema (promql_driver.forward.threshold: float64) The threshold for the forward criteria.
  * @schema (promql_driver.forward.operator: string) The operator for the forward criteria. oneof: `gt | lt | gte | lte | eq | neq`
  * @schema (promql_driver.backward.threshold: float64) The threshold for the backward criteria.
  * @schema (promql_driver.backward.operator: string) The operator for the backward criteria. oneof: `gt | lt | gte | lte | eq | neq`
  * @schema (promql_driver.reset.threshold: float64) The threshold for the reset criteria.
  * @schema (promql_driver.reset.operator: string) The operator for the reset criteria. oneof: `gt | lt | gte | lte | eq | neq`
  */
  promql_driver: promql_driver_defaults,
  /**
  * @schema (average_latency_driver.flow_selector: aperture.spec.v1.FlowSelector required) Identify the service and flows whose latency needs to be measured.
  * @schema (average_latency_driver.forward.threshold: float64) The threshold for the forward criteria.
  * @schema (average_latency_driver.backward.threshold: float64) The threshold for the backward criteria.
  * @schema (average_latency_driver.reset.threshold: float64) The threshold for the reset criteria.
  */
  average_latency_driver: average_latency_driver_defaults,
  /**
  * @schema (percentile_latency_driver.flux_meter: aperture.spec.v1.FluxMeter required) FluxMeter specifies the flows whose latency needs to be measured and parameters for the histogram metrics.
  * @schema (percentile_latency_driver.forward.threshold: float64) The threshold for the forward criteria.
  * @schema (percentile_latency_driver.backward.threshold: float64) The threshold for the backward criteria.
  * @schema (percentile_latency_driver.reset.threshold: float64) The threshold for the reset criteria.
  * @schema (percentile_latency_driver.percentile: float64) The percentile to be used for latency measurement.
  */
  percentile_latency_driver: percentile_latency_driver_defaults,
  /**
  * @schema (ema_latency_driver.flow_selector: aperture.spec.v1.FlowSelector required) Identify the service and flows whose latency needs to be measured.
  * @schema (ema_latency_driver.forward.latency_tolerance_multiplier: float64) The threshold for the forward criteria.
  * @schema (ema_latency_driver.backward.latency_tolerance_multiplier: float64) The threshold for the backward criteria.
  * @schema (ema_latency_driver.reset.latency_tolerance_multiplier: float64) The threshold for the reset criteria.
  * @schema (ema_latency_driver.ema: aperture.spec.v1.EMAParameters required) The parameters for the exponential moving average.
  */
  ema_latency_driver: ema_latency_driver_defaults,
  /**
  * @schema (rollout_policy.load_shaper: aperture.spec.v1.LoadShaperParameters required) Identify the service and flows of the feature that needs to be rolled out. And specify feature rollout steps.
  * @schema (rollout_policy.components: []aperture.spec.v1.Component) List of additional circuit components.
  * @schema (rollout_policy.resources: aperture.spec.v1.Resources) List of additional resources.
  * @schema (rollout_policy.evaluation_interval: string) The interval between successive evaluations of the Circuit.
  * @schema (rollout_policy.drivers.promql_drivers: []promql_driver) List of promql drivers that compare results of a Prometheus query against forward, backward and reset thresholds.
  * @schema (rollout_policy.drivers.average_latency_drivers: []average_latency_driver) List of drivers that compare average latency against forward, backward and reset thresholds.
  * @schema (rollout_policy.drivers.percentile_latency_drivers: []percentile_latency_driver) List of drivers that compare percentile latency against forward, backward and reset thresholds.
  * @schema (rollout_policy.drivers.ema_latency_drivers: []ema_latency_driver) List of drivers that compare trend latency against forward, backward and reset thresholds.
  */
  rollout_policy: rollout_policy_defaults,
  /**
  * @param (dashboard.refresh_interval: string) Refresh interval for dashboard panels.
  * @param (dashboard.time_from: string) From time of dashboard.
  * @param (dashboard.time_to: string) To time of dashboard.
  */
  dashboard: {
    refresh_interval: '5s',
    time_from: 'now-15m',
    time_to: 'now',
    /**
    * @param (dashboard.datasource.name: string) Datasource name.
    * @param (dashboard.datasource.filter_regex: string) Datasource filter regex.
    */
    datasource: {
      name: '$datasource',
      filter_regex: '',
    },
  },
  rollout_policy_base: rollout_policy_base_defaults,
}