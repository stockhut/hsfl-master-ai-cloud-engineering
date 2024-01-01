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

Generate the public and private keys for JWT signing[^1]:
```shell
openssl ecparam -name prime256v1 -genkey -noout -out jwt_private_key.key
openssl ec -in jwt_private_key.key -pubout -out jwt_public_key.key
```

### Reverse proxy

The reverse proxy is configured via `reverse-proxy.yml`

[^1]: https://notes.salrahman.com/generate-es256-es384-es512-private-keys/