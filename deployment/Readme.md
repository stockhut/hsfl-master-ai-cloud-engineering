```shell
minikube addons enable ingress
```

```shell
minikube tunnel
```

http://127.0.0.1

# Secrets

```shell
 kubectl create secret generic jwt-private-key --from-file jwt_private_key.key

 kubectl create secret generic jwt-public-key --from-file jwt_public_key.key

```

# Persistence

Postgres data is persisted to `/data` on the host using a volume
