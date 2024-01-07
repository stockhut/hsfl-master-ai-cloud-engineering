# Kubernetes

## Prerequisites

### Prepare minikube

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
