# Recipe Service

### Testing

Install `mockgen`:
```shell
 go install go.uber.org/mock/mockgen@latest
```

Before testing, generate the mocks:

```shell
go generate _mocks/gen.go
```
## Configuration

### Environment variables
| Key            | Value                                                                              |
|----------------|------------------------------------------------------------------------------------|
| JWT_PUBLIC_KEY | Path to a jwt public keyfile                                                       |
| AUTH_RPC_TARGET| Address of a authentications gRPC service (_IP:PORT_)                              |
| PG_CONN_STRING | A PostgreSQL connection string, e.g. `postgres://postgres:password@127.0.0.1:5432` |

## Development

For development purposes, there is a docker compose file with PostgreSQL in `dev-db/`

```shell
cd dev-db
docker compose up -d
```

## gRPC

The authentication service exposes a gRPC service.
gRPC and Protobuf code can be imported by other programs from `auth-proto`.
This includes generated code as well as error definitions and converters.

## Responsibilities

- CRUD for recipes
- searching
- tags
- likes from users