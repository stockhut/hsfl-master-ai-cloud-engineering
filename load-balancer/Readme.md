# Load balancer

## Configuration

See `config.yaml` for an example.

Load balancing targets can be defined statically by their IP/hostname or discovered via docker image names.

| Key           | Meaning                                                          |
|---------------|------------------------------------------------------------------|
| listen        | The address to listen on(_IP:PORT_)                              |
| hosts         | A list of load balancing targets (_IP:PORT_)                     |
| images        | A list of container image names used to discover running targets |
| containerPort | The port to use with discovered docker containers                |

## Running

```shell
go run ./cmd
```
## Testing

### Generating mocks
```shell
go generate strategies/ip-hash/_mocks/gen.go
```

## Thread safety

Since requests are handled concurrently and some state is accessed across thread boundaries, there are some efforts made to avoid race conditions.


### Load balancer

The slice holding all healthy replicas is guarded by a `sync.RWMutex`.
### Round-robin

Request counting to determine which host has to handle the request is implemented using an `atomic.Int32`.

### Least connections

The map that keeps track of hosts and their connections is guarded by a `sync.RWMutex`.