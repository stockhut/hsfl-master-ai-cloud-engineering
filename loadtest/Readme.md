# Loadtest

This tool can be used to loadtest a webserver.
The number of requests per second (RPS) can be configured in multiple steps, where the RPS is hold for a specified duration.
Between the phases, the number of RPS is interpolated linearly.

Multiple target URLs with different methods and optional payloads can be specified.

## Running

```shell
go run ./cmd
```

## Configuration

See `loadtest.yaml` for an example configuration.

`rampup` and `duration` can be specified using Go duration strings (see https://pkg.go.dev/time#ParseDuration)

| Attribute        | Meaning                                                                                                                                                                                                        |
|------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| host             | The target host (_IP:PORT_)                                                                                                                                                                                    |
| responseStats    | When `true`, the server response is awaited and http status code and average response times are logged. This setting might cause some requests to fail if the OS runs out of free ports for outgoing requests. |
| headers          | A map of http header to set. The `Host` header is added automatically if there is no value specified manually. The `Content-Length` header is set automatically.                                               |
| phases           | A list of phases to execute in order, each item consisting of the following attributes                                                                                                                         |
| phases.rps       | Requests per Second                                                                                                                                                                                            |
| phases.rampup    | Time over which the number of requests increases until `rps` is reached (e.g. `30s`)                                                                                                                           |
| phases.duration  | Total duration of the test (e.g. `30s`)                                                                                                                                                                        |
| targets          | List of URLs to target. For each request a random entry is picked                                                                                                                                              |
| targets.method   | HTTP Method                                                                                                                                                                                                    |
| targets.path     | Target Path                                                                                                                                                                                                    |
| targets.body     | Optional HTTP Body. Takes precedence over `bodyFile`                                                                                                                                                           |
| targets.bodyFile | Optional file to be used as HTTP body                                                                                                                                                                          |