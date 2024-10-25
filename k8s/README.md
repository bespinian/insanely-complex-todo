# Kubernetes Deployment

This folder contains plain YML files for a Kubernetes setup of our todo app. It is basically the
same architecture as the [Docker Compose setup](../).

You can install this in a local Kubernetes cluster using Docker Desktop. In this case you also need
to install Traefik as Ingress controller. If you install this in some other cluster, you might
already have an ingress controller and don't need Traefik. That's up to you.

## Prerequisites

You need a Kubernetes cluster of some sort. If you don't have one at hand, just activate the one
in Docker Desktop. The internet tells you how.

If you have no ingress controller in your cluster, install Traefik as follows:

```sh
kubectl create ns traefik
kubectl -n traefik apply -f k8s/traefik
```

## Install ICT

The provided YAML files assume you have a local cluster and therefore access to it via localhost.
If you're working with a "real" remote cluster, you might need to change the
[Ingress Resource](./apps/ingress.yaml).

```sh
kubectl create namespace todo
kubectl --namespace todo apply -f k8s/services
kubectl --namespace todo apply -f k8s/apps
```

In a remote cluster, you should be done and be able to open the app in your browser.

In a local cluster, we now need to open a port-forward to Traefik. Then we can access the app on
localhost as usual.

```sh
kubectl --namespace traefik port-forward services/traefik-web 8000:80
open localhost:8000
```
