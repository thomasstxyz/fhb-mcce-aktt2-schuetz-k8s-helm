# Demo Go Microservice Delivery with Helm

![](assets/arch-diagram-app.png)

You can find the helm chart in `./helmcharts/app` directory.
It is packaged and pushed to the OCI Repository
`oci://ghcr.io/thomasstxyz/charts/fhb-mcce-aktt2-schuetz-k8s-helm`.

You can find different delivery examples in `./examples/` directory.

- [Delivery with Flux and Helm](examples/flux-helm-delivery/README.md)
- [Delivery with Helm CLI](examples/helm-cli/README.md)
- [Delivery with kubectl](examples/kubectl/README.md)
