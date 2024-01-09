# Loadtest

## Running

```shell
go run ./cmd
```

## Configuration

See `loadtest.yaml` for an example configuration.

| Attribute     | Meaning                                                                                                                                                                                                        |
|---------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| host          | The target host (_IP:PORT_)                                                                                                                                                                                    |
| responseStats | When `true`, the server response is awaited and http status code and average response times are logged. This setting might cause some requests to fail if the OS runs out of free ports for outgoing requests. |
| headers       | A map of http header to set. The `Host` header is added automatically if there is no value specified manually                                                                                                  |
| phases        | A list of phases to execute in order, each item consisting of the following attributes                                                                                                                         |
| rampup        | Time over which the number of users increases until `users` is reached, (in ms)                                                                                                                                |
| duration      | Total duration of the test (in ms)                                                                                                                                                                             |
| targets       | List of URLs to target. For each request a random entry is picked                                                                                                                                              |

