# Reverse proxy

## Running

```shell
go run ./cmd/
```

## Configuration

See `config.yml` for an example configuration.
Routes are matched in order, so be sure to configure "/foo" before "/".

### Environment variables

| Key         | Value                   |
|-------------|-------------------------|
| CONFIG_FILE | Path to the config file |