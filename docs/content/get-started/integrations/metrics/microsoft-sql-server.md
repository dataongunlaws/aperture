---
title: Microsoft SQL Server
description: Integrating Microsoft SQL Server Metrics
keywords:
  - sqlserver
  - otel
  - opentelemetry
  - collector
  - metrics
---

Before proceeding, ensure that you have [built][build] Aperture Agent with the
`sqlserverreceiver` extension enabled, so that [sqlserverreceiver][receiver] is
available.

You can configure [Custom metrics][custom-metrics] for Microsoft SQL Server
using the following configuration in the [Aperture Agent's
config][agent-config]:

```yaml
otel:
  custom_metrics:
    sqlserver:
      per_agent_group: true
      pipeline:
        processors:
          - batch
        receivers:
          - sqlserver
      processors:
        batch:
          send_batch_size: 10
          timeout: 10s
      receivers:
        sqlserver: [sqlserverreceiver configuration here]
```

[build]: /reference/aperturectl/build/agent/agent.md
[receiver]:
  https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/sqlserverreceiver
[custom-metrics]: /reference/configuration/agent.md#custom-metrics-config
[agent-config]: /reference/configuration/agent.md#agent-o-t-e-l-config