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
| Key            | Value                                 |
|----------------|---------------------------------------|
| JWT_PUBLIC_KEY | Path to a jwt public keyfile          |
| PG_CONN_STRING | A PostgreSQL connection string, e.g. `postgres://postgres:password@127.0.0.1:5432` |

## Development

For development purposes, there is a docker compose file with PostgreSQL in `dev-db/`

```shell
cd dev-db
docker compose up -d
```

## Responsibilities

- CRUD for recipes
- searching
- tags
- likes from users