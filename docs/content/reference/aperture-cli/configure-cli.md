---
title: Configuration
keywords:
  - cli
sidebar_position: 2
---

Configure aperturectl to point to your Aperture Cloud endpoint: Save the
following as `~/.aperturectl/config`:

```toml
[controller]
url = "ORGANIZATION_NAME.app.fluxninja.com:443"
project_name = "PROJECT_NAME"
api_key = "PERSONAL_API_KEY"
```

Replace `ORGANIZATION_NAME` with the Aperture Cloud organization name and
`PERSONAL_API_KEY` with the Personal API key linked to the user. If a Personal
API key has not been created, generate a new one through the Aperture Cloud UI.
Refer to [Personal API Keys][api-keys] for additional information.

:::info

See also [aperturectl configuration file format reference][aperturectl-config].

:::

:::note Self-hosted Aperture Controller

With a [self-hosted][self-hosted] Aperture Controller, if the Controller is at
the cluster pointed at by `~/.kube/config` or `KUBECONFIG`, no configuration
file nor flags are needed at all. Otherwise, you need the `--controller` flag.

:::

[self-hosted]: /self-hosting/self-hosting.md
[aperturectl-config]: /reference/configuration/aperturectl.md
[api-keys]: /reference/aperture-cli/personal-api-keys.md