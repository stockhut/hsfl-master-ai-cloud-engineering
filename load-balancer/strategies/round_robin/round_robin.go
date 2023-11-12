package round_robin

import "net/http"

type RoundRobin struct {
	idx int
}

func New() *RoundRobin {
	return &RoundRobin{
		idx: 0,
	}
}

func (r *RoundRobin) GetTarget(_ *http.Request, replicas []string, f func(host string)) {
	host := replicas[r.idx%len(replicas)]
	r.idx += 1

	f(host)
}
