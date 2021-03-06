---
title: Connector Service sub-chart
---

To configure the Connector Service sub-chart, override the default values of its [`values.yaml`](https://github.com/kyma-project/kyma/blob/main/resources/application-connector/charts/connector-service/values.yaml) file. This document describes parameters that you can configure.

>**TIP:** See how to [change Kyma settings](../../04-operation-guides/operations/03-change-kyma-config-values.md).

## Configurable parameters

This table lists the configurable parameters, their descriptions, and default values:

| Parameter | Description | Default value |
|-----------|-------------|---------------|
| **deployment.args.tokenLength**| Specifies the number of characters in a registration token. | `64` |
| **deployment.args.appTokenExpirationMinutes** | Specifies the period of time after which a token for an Application expires. It is provided in minutes. | `5` |
| **deployment.args.runtimeTokenExpirationMinutes** | Specifies the period of time after which a token for a Runtime expires. It is provided in minutes. | `10` |
| **deployment.args.appValidityTime** | Specifies the period of time during which certificates that the service issues for an Application are valid. It is provided in days. | `92d` |
| **deployment.args.runtimeValidityTime** | Specifies the period of time during which certificates that the service issues for a Runtime are valid. It is provided in days. | `92d` |
| **deployment.args.central** | Determines whether Connector Service works in the central mode. | `false` |
| **deployment.args.requestLogging** | Enables logging of incoming requests.| `false ` |
| **deployment.envvars.country** | Specifies a country expected in a Certificate Signing Request. It is provided as a two-letter country code. | `DE` |
| **deployment.envvars.organization** | Specifies an organization expected in a Certificate Signing Request. | `Organization` |
| **deployment.envvars.organizationalunit** | Specifies an organizational unit expected in a Certificate Signing Request. | `OrgUnit` |
| **deployment.envvars.locality** | Specifies a locality expected in a Certificate Signing Request. | `Waldorf` |
| **deployment.envvars.province** | Specifies a province expected in a Certificate Signing Request. | `Waldorf` |