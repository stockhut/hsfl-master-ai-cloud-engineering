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
| Key            | Value                                                 |
|----------------|-------------------------------------------------------|
| JWT_PUBLIC_KEY | Path to a jwt public keyfile                          |
| SQLITE_DB_PATH | Path to the SQLite database file                      |
| AUTH_RPC_TARGET| Address of a authentications gRPC service (_IP:HOST_) |

## gRPC

The authentication service exposes a gRPC service.
gRPC and Protobuf code can be imported by other programs from `auth-proto`.
This includes generated code as well as error definitions and converters.

## Responsibilities

- CRUD for recipes
- searching
- tags
- likes from users