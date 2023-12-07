# Loadtest

## Running

```shell
go run ./cmd
```

## Configuration

See `load.config.json` for an example configuration.

| Attribute | Meaning                                                                         |
|-----------|---------------------------------------------------------------------------------|
| users     | How many concurrent users should be simulated                                   |
| rampup    | Time over which the number of users increases until `users` is reached, (in ms) |
| duration  | Total duration of the test (in ms)                                              |
| targets   | List of URLs to target. For each request a random entry is picked               |

