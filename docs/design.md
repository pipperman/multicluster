## Alien

---

<img src="logo.png" width="200" height="100">

Alien, is an open source system for managing cloud container cluster.It provides basic mechanisms for cce.


## Component

---

| Component  | Version | Deploy     |
|------------|---------|------------|
| Mysql      | 5.6     | Deployment |
| Jaeger     | 1.34.1  | Operator   |
| Prometheus | 2.10    | Operator   |


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
kubectl apply -f jaeger-all-in-one.yaml
```

- ### Prometheus



### 3. Deploy Alien