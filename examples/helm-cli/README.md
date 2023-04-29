# Delivery with Helm CLI

You may install the application via the Helm CLI,
by following the below steps.

## Prerequisites

- Kubernetes 1.16+
- Helm 3+

## Install

Install the backend services:

```bash
helm install backend1 oci://ghcr.io/thomasstxyz/charts/fhb-mcce-aktt2-schuetz-k8s-helm/app --version 0.1.5
helm install backend2 oci://ghcr.io/thomasstxyz/charts/fhb-mcce-aktt2-schuetz-k8s-helm/app --version 0.1.5
helm install backend3 oci://ghcr.io/thomasstxyz/charts/fhb-mcce-aktt2-schuetz-k8s-helm/app --version 0.1.5
```

Install the frontend service:
> Note: Escape the commas with a backslash on a Bash shell!

```bash
helm install frontend oci://ghcr.io/thomasstxyz/charts/fhb-mcce-aktt2-schuetz-k8s-helm/app --version 0.1.5 --set backendServices="backend1-app:8080\,backend2-app:8080\,backend3-app:8080"
```

Now you may port-forward the frontend service:

```bash
kubectl -n default port-forward svc/frontend-app 8080
```

and have a look via a web browser on your local machine at [http://127.0.0.1:8080/](http://127.0.0.1:8080/)

## Uninstall

```bash
helm uninstall frontend
helm uninstall backend1
helm uninstall backend2
helm uninstall backend3
```

