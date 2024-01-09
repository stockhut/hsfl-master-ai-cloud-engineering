# Kubernetes

## Prerequisites

You need a working kubernetes cluster.
We used [minikube](https://minikube.sigs.k8s.io/docs/) and [k3s](https://k3s.io/).
Our k3s VM configuration and build instructions for can be found in [../k8s-vm](https://github.com/stockhut/hsfl-master-ai-cloud-engineering/blob/main/k8s-vm)

We expect a working ingress-controller in the cluster.

### With minikube

```shell
minikube addons enable ingress
```

### Secrets

Generate the public and private keys for JWT signing[^1]:
```shell
openssl ecparam -name prime256v1 -genkey -noout -out jwt_private_key.key
openssl ec -in jwt_private_key.key -pubout -out jwt_public_key.key
```

Create kubernetes secrets consumed by the services:
```shell
 kubectl create secret generic jwt-private-key --from-file jwt_private_key.key

 kubectl create secret generic jwt-public-key --from-file jwt_public_key.key
```

## Creating

```shell
kubectl create -f . --recursive
```

## Accessing via minikube:

```shell
minikube tunnel
```

http://127.0.0.1


## Persistence

Postgres data is persisted to `/data` on the host using a volume
