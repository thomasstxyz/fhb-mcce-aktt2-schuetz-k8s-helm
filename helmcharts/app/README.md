# app

Installs the [app](https://github.com/thomasstxyz/fhb-mcce-aktt2-schuetz-k8s-helm) helm chart.

## Prerequisites

- Kubernetes 1.16+
- Helm 3+

## Install Helm Chart

```console
helm install [RELEASE_NAME] oci://ghcr.io/thomasstxyz/charts/fhb-mcce-aktt2-schuetz-k8s-helm/app --version 0.1.5
```

_See [helm install](https://helm.sh/docs/helm/helm_install/) for command documentation._

## Uninstall Helm Chart

```console
helm uninstall [RELEASE_NAME]
```

This removes all the Kubernetes components associated with the chart and deletes the release.

_See [helm uninstall](https://helm.sh/docs/helm/helm_uninstall/) for command documentation._
