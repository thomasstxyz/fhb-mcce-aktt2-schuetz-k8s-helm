# Delivery with Flux and Helm

First, ensure you have the flux CLI installed locally.
Then export your GitHub user and API Token as environment variables.

```bash
export GITHUB_TOKEN=<your-token>
export GITHUB_USER=<your-username>
```

Check flux prerequisites for your cluster.

```bash
flux check --pre
```

Bootstrap Flux onto your cluster and git repo.

```bash
flux bootstrap github \
  --owner=$GITHUB_USER \
  --repository=fhb-mcce-aktt2-schuetz-k8s-helm \
  --branch=main \
  --path=./clusters/my-cluster \
  --personal
```

Now you can create a file in the path `./clusters/my-cluster` which Flux watches.

```bash
touch clusters/my-cluster/apps.yaml
```

With the following contents:

```yaml
---
apiVersion: kustomize.toolkit.fluxcd.io/v1beta2
kind: Kustomization
metadata:
  name: apps
  namespace: flux-system
spec:
  interval: 10m0s
  path: ./examples/flux-helm-delivery
  prune: true
  sourceRef:
    kind: GitRepository
    name: flux-system
```

This will create a Flux `Kustomization`,
which watches the path `./examples/flux-helm-delivery`
and applies manifests it finds there to your cluster.

This path contains our `HelmRepository` and `HelmRelease` manifests.
After successful reconciliation, you should see the App being deployed.

Now you may port-forward the frontend service:

```bash
kubectl -n default port-forward svc/frontend-app 8080
```

and have a look via a web browser on your local machine at [http://127.0.0.1:8080/](http://127.0.0.1:8080/)
