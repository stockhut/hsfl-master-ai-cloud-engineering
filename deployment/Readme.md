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

```