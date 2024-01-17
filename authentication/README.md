# Authentication Service

## Building

### Protobuf code generation
(not needed when using Docker)
```shell
go generate auth-proto/gen.go
```

### SQLC code generation
(not needed when using Docker)
```shell
go generate accounts/_sqlc/gen.go 
```

## Docker Image

```shell
docker build .. --file Dockerfile
```

## Configuration

### Environment variables
| Key             | Value                                                                              |
|-----------------|------------------------------------------------------------------------------------|
| JWT_PRIVATE_KEY | Path to a jwt private keyfile                                                      |
| PG_CONN_STRING | A PostgreSQL connection string, e.g. `postgres://postgres:password@127.0.0.1:5432` |

## Responsibilities

- creation and management of user accounts
- login functionality

## Endpoints
| Action         | Endpoint                         |
|----------------|----------------------------------|
| Login          | `/api/v1/authentication/login`   |
| Create Account | `/api/v1/authentication/account` |