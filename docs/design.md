## Alien

---

<img src="logo.png" width="200" height="100">

Alien, is an open source system for managing cloud container cluster.It provides basic mechanisms for cce.


## Component

---
- Dependency

| Component  | Version | Deploy     |
|------------|---------|------------|
| Mysql      | 5.6     | Deployment |
| Cert-Manager  | 1.8.0   | Operator |
| Jaeger     | 1.34.1  | Operator   |
| Prometheus | 2.10    | Operator   |

- Micro Service


## Quick Start

---

### 1. Push Image

- mysql:5.6
- quay.io/jaegertracing/jaeger-operator:1.34.1
- gcr.io/kubebuilder/kube-rbac-proxy:v0.8.0

### 2. Deploy Component

- ### MySql


```shell
kubectl apply -f mysql.yaml
```

- ### Cert-Manager

#### online
```shell
# step 1
OS=$(go env GOOS); ARCH=$(go env GOARCH); curl -sSL -o cmctl.tar.gz https://github.com/cert-manager/cert-manager/releases/download/v1.7.2/cmctl-$OS-$ARCH.tar.gz
# step 2
cmctl x install
```

#### offline

[cert-manager.yaml](../deploy/kubernetes/cert-manager/cert-manager.yaml)

```shell
# download image
quay.io/jetstack/cert-manager-webhook:v1.8.0
quay.io/jetstack/cert-manager-controller:v1.8.0
quay.io/jetstack/cert-manager-cainjector:v1.8.0

kubectl apply -f cert-manager.yaml
```


- ### Jaeger-Operator

#### create jaeger operator
[jaeger-operator.yaml](../deploy/kubernetes/observability/jaeger-operator.yaml)

```shell
kubectl create ns observability
kubectl apply -f jaeger-operator.yaml
```

#### create jaeger instance
[jaeger-all-in-one.yaml](../deploy/kubernetes/observability/jaeger-all-in-one.yaml)

```shell
# download image
jaegertracing/all-in-one:1.34.1
jaegertracing/jaeger-agent:1.34.1

kubectl apply -f jaeger-all-in-one.yaml
```

- ### Prometheus



### 3. Install Kompose

- [Installation](https://kompose.io/installation/)

### 4. Deploy Alien

```shell
kompose convert -f deploy/docker-compose/docker-compose.yaml -o deploy/kubernetes

kubectl apply -f .
```