---
title: vCenter
description: Integrating vCenter Metrics
keywords:
  - vcenter
  - otel
  - opentelemetry
  - collector
  - metrics
---

Before proceeding, ensure that you have [built][build] Aperture Agent with the
`vcenterreceiver` extension enabled, so that [vcenterreceiver][receiver] is
available.

You can configure [Custom metrics][custom-metrics] for vCenter using the
following configuration in the [Aperture Agent's config][agent-config]:

```yaml
otel:
  custom_metrics:
    vcenter:
      per_agent_group: true
      pipeline:
        processors:
          - batch
        receivers:
          - vcenter
      processors:
        batch:
          send_batch_size: 10
          timeout: 10s
      receivers:
        vcenter: [vcenterreceiver configuration here]
```

[build]: /reference/aperturectl/build/agent/agent.md
[receiver]:
  https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/vcenterreceiver
[custom-metrics]: /reference/configuration/agent.md#custom-metrics-config
[agent-config]: /reference/configuration/agent.md#agent-o-t-e-l-config