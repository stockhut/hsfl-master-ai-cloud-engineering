# Dev setup

## JWT keys

The authentication service needs a private key to sign JWTs.
The keys are provided by mounting this directory as `/keys`.

Generate the private key:
```shell
openssl ecparam -name prime256v1 -genkey -noout -out jwt_private_key.key
```