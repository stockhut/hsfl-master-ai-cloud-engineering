# Authentication Service

## Building

### Protobuf code generation

```shell
go generate proto/gen.go
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
- authentification/authorization of incoming requests