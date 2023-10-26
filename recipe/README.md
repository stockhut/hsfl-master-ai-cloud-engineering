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

## Responsibilities

- CRUD for recipes
- searching
- tags
- likes from users