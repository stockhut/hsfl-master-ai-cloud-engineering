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
| Key            | Value                            |
|----------------|----------------------------------|
| JWT_PUBLIC_KEY | Path to a jwt public keyfile     |
| SQLITE_DB_PATH | Path to the SQLite database file |

## Responsibilities

- CRUD for recipes
- searching
- tags
- likes from users