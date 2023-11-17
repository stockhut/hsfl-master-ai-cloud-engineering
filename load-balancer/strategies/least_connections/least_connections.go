package least_connections

import (
	"net/http"
	"sync"
)

type LeastConnections struct {
	connectionCount map[string]uint32
	m               *sync.RWMutex
}

func New() *LeastConnections {
	return &LeastConnections{
		m:               &sync.RWMutex{},
		connectionCount: make(map[string]uint32),
	}
}

func (l *LeastConnections) GetTarget(_ *http.Request, replicas []string, f func(host string)) {

	//x := rand.Int() / 10000000
	l.m.RLock()

	//fmt.Printf("%d Start Picking: connection count: \n%v\n", x, l.connectionCount)

	minConnections := uint32(9999999)
	var minHost string
	for _, host := range replicas {
		c := l.connectionCount[host]

		if c == 0 {
			minHost = host
			break
		}
		if c < minConnections {
			minConnections = c
			minHost = host
		}
	}
	l.m.RUnlock()

	l.m.Lock()
	l.connectionCount[minHost] += 1
	//fmt.Printf("%d Picked: %s\n", x, minHost)
	l.m.Unlock()

	f(minHost)

	l.m.Lock()
	l.connectionCount[minHost] -= 1
	l.m.Unlock()

	//fmt.Printf("%d End. New connection count: \n%v\n", x, l.connectionCount)

}
