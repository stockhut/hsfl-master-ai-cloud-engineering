# Recipe Service

To build this manually without docker, you also need to generate the protobuf code for the Authentication service.
## Docker Image

```shell
docker build .. --file Dockerfile
```

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


## Profiling

A [pprof](https://pkg.go.dev/net/http/pprof@go1.21.6) endpoint is available at port 6060

## Responsibilities

- CRUD for recipes
- searching
- tags
- likes from users

## Endpoints

| Action                 | Endpoint                    | Note                 |
|------------------------|----------------------------------------------------|
| Create Recipe          | `/api/v1/recipe`            | Needs valid JWT token|
| Get recipe(s) by author| `/api/v1/recipe/by/<author>`| Needs valid JWT token|
| Get recipe by ID       | `/api/v1/recipe/<id>`       | Needs valid JWT token|
| Get own recipe(s)      | `/api/v1/recipe/by/<author>`| Needs valid JWT token|
| Check health           | `/health`                   |                      |
