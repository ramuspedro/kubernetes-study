# Tutorial Kubernetes

## Kind commands

Using kind (https://kind.sigs.k8s.io/docs/user/quick-start/#installing-from-release-binaries)

```shell
# creating clusters from file
$ kind create cluster --config=./k8s/kind.yml --name=fullcycle
```

## Kubectl

```shell
# using kind context clusters
$ kubectl cluster-info --context kind-fullcycle

# list clusters from that cluster
$ kubectl get nodes

# list clusters available
$ kubectl config get-clusters

# use context
$ kubectl config use-context kind-fullcycle
``