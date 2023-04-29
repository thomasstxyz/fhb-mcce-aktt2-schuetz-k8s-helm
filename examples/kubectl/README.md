# Delivery with kubectl

It is used for client-server communication. Helm is installed in the system namespace kube-system, so it can perform operations on a Kubernetes cluster. To install the server part on the cluster and add local settings, run the command:

```
helm init



1. How to apply the resource to our Cluster:

```
kubectl apply -f  helmrepository.yaml
```
2. When running following Command you can see HelmChart
```
kubectl get helmchart
```
3. To see artefacts and  other parameters from HelmChart you can run following:

```
kubectl describe helmchart podinfo

