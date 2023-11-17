# Load balancer

## Thread safety

Since requests are handled concurrently and some state is accessed across thread boundaries, there are some efforts made to avoid race conditions.


### Load balancer

The slice holding all healthy replicas is guarded by a `sync.RWMutex`.
### Round-robin

Request counting to determine which host has to handle the request is implemented using an `atomic.Int32`.

### Least connections

The map that keeps track of hosts and their connections is guarded by a `sync.RWMutex`.