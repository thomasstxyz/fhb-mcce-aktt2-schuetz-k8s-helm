# Delivery with kubectl

## Install

```bash
kubectl apply -f examples/kubectl/manifest.yaml
```

Now you may port-forward the frontend service:

```bash
kubectl -n default port-forward svc/frontend-app 8080
```

and have a look via a web browser on your local machine at [http://127.0.0.1:8080/](http://127.0.0.1:8080/)

## Uninstall

```bash
kubectl delete -f examples/kubectl/manifest.yaml
```

## Notes

> The `manifest.yaml` was previously built like this:
>
>     helm template backend1 ./helmcharts/app > examples/kubectl/manifest.yaml
>     helm template backend2 ./helmcharts/app >> examples/kubectl/manifest.yaml
>     helm template backend3 ./helmcharts/app >> examples/kubectl/manifest.yaml
>     helm template frontend ./helmcharts/app --version 0.1.5 --set backendServices="backend1-app:8080\,backend2-app:8080\,backend3-app:8080" >> examples/kubectl/manifest.yaml