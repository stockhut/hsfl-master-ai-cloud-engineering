# Reverse proxy

## Running

```shell
go run ./cmd/
```

## Configuration

See `config.yml` for an example configuration.
Routes are matched in order, so be sure to configure "/foo" before "/".