# Cluster UI

[Cluster UI by AppsCode](https://github.com/bytebuilders) - ByteBuilders Cluster UI

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/billing-ui --version=v2023.03.23
$ helm upgrade -i billing-ui appscode/billing-ui -n ace-system --create-namespace --version=v2023.03.23
```

## Introduction

This chart deploys a Cluster UI on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.21+

## Installing the Chart

To install/upgrade the chart with the release name `billing-ui`:

```bash
$ helm upgrade -i billing-ui appscode/billing-ui -n ace-system --create-namespace --version=v2023.03.23
```

The command deploys a Cluster UI on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `billing-ui`:

```bash
$ helm uninstall billing-ui -n ace-system
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `billing-ui` chart and their default values.

|                 Parameter                  |                                                             Description                                                              |         Default         |
|--------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------|-------------------------|
| replicaCount                               |                                                                                                                                      | <code>1</code>          |
| registryFQDN                               | Docker registry fqdn used to pull app related images. Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image} | <code>ghcr.io</code>    |
| image.registry                             | Docker registry used to pull app container image                                                                                     | <code>appscode</code>   |
| image.repository                           | App container image                                                                                                                  | <code>billing-ui</code> |
| image.tag                                  | Overrides the image tag whose default is the chart appVersion.                                                                       | <code>""</code>         |
| image.pullPolicy                           |                                                                                                                                      | <code>Always</code>     |
| imagePullSecrets                           |                                                                                                                                      | <code>[]</code>         |
| nameOverride                               |                                                                                                                                      | <code>""</code>         |
| fullnameOverride                           |                                                                                                                                      | <code>""</code>         |
| serviceAccount.name                        |                                                                                                                                      | <code>""</code>         |
| podAnnotations                             |                                                                                                                                      | <code>{}</code>         |
| podSecurityContext                         |                                                                                                                                      | <code>{}</code>         |
| securityContext                            |                                                                                                                                      | <code>{}</code>         |
| service.type                               |                                                                                                                                      | <code>ClusterIP</code>  |
| service.port                               |                                                                                                                                      | <code>80</code>         |
| resources                                  |                                                                                                                                      | <code>{}</code>         |
| autoscaling.enabled                        |                                                                                                                                      | <code>false</code>      |
| autoscaling.minReplicas                    |                                                                                                                                      | <code>1</code>          |
| autoscaling.maxReplicas                    |                                                                                                                                      | <code>100</code>        |
| autoscaling.targetCPUUtilizationPercentage |                                                                                                                                      | <code>80</code>         |
| nodeSelector                               |                                                                                                                                      | <code>{}</code>         |
| tolerations                                |                                                                                                                                      | <code>[]</code>         |
| affinity                                   |                                                                                                                                      | <code>{}</code>         |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i billing-ui appscode/billing-ui -n ace-system --create-namespace --version=v2023.03.23 --set replicaCount=1
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i billing-ui appscode/billing-ui -n ace-system --create-namespace --version=v2023.03.23 --values values.yaml
```