package load_balancer

import (
	"fmt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/reverse-proxy"
	"net/http"
)

type balancingStrategy interface {
	GetTarget(r *http.Request, replicas []string, f func(host string))
}

func New(replicas []string, strategy balancingStrategy) LoadBalancer {
	return LoadBalancer{
		replicas: replicas,
		strategy: strategy,
	}
}

type LoadBalancer struct {
	replicas []string
	strategy balancingStrategy
}

func (lb LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	lb.strategy.GetTarget(r, lb.replicas, func(host string) {
		fmt.Printf("Target: %s\n", host)
		err := reverse_proxy.Forward(w, r, host)
		if err != nil {
			// TODO: mark target host as unhealthy
			panic(err)
		}
	})
}
