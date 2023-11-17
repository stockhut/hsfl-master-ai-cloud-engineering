package round_robin

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

type RoundRobin struct {
	idx atomic.Int32
}

func New() *RoundRobin {
	return &RoundRobin{
		idx: atomic.Int32{},
	}
}

func (r *RoundRobin) GetTarget(_ *http.Request, replicas []string, f func(host string)) {
	i := r.idx.Add(1)
	host := replicas[i%int32(len(replicas))]

	fmt.Printf("Picked %s from healthy list: %v\n", host, replicas)
	f(host)
}
