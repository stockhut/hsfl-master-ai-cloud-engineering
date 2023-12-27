# Dev setup

```shell
docker compose up
```

http://localhost:5000

This docker compose setup contains all services, the reverse proxy and a postgres instance.

## Configuration

### JWT keys

The authentication service needs a private key to sign JWTs.
The keys are provided by mounting this directory as `/keys`.

Generate the private key:
```shell
openssl ecparam -name prime256v1 -genkey -noout -out jwt_private_key.key
```

### Reverse proxy

The reverse proxy is configured via `reverse-proxy.yml`