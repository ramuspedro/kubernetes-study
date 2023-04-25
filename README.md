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
```

## Run go

```sh
# run with go
$ go run server.go
```

Run with docker

```sh
# docker build
$ docker build -t ramuspedro/hello-go .

# docker run
$ docker run --rm -p 8080:8080 ramuspedro/hello-go
```

Run kubernetes

```sh
# apply pod.yml
$ kubectl apply -f k8s/pod.yml 

# list pods
$ kubectl get pods

# port forward if port not exposed
$ kubectl port-forward pod/goserver 8080:8080

# delete pod
$ kubectl delete pod goserver

# run replicaset
$ kubectl apply -f k8s/replicaset.yml

# get replicaset
$ kubectl get replicasets

# decribe infomations of pods
$ kubectl decribe pod pod-name

# run deployment
$ kubectl apply -f k8s/deployment.yml
```