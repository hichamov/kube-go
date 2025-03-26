# Kubernetes client-go

In order to interact with a kubernetes cluster with a Go application, we need the official client-go library.

There are additional libraries to use:

* client-go: exposes all the interfaces to interact with k8s api
* api: exposes the kubernetes objects types
* api-machinery: utility methods to develop APIs similar to kubernetes.

The module we will import is 'k8s.io/client-go*'
