package load_balancer

import (
	"fmt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/reverse-proxy"
	"net/http"
	"time"
)

type balancingStrategy interface {
	GetTarget(r *http.Request, replicas []string, f func(host string))
}

func New(replicas []string, strategy balancingStrategy, healthcheckInterval time.Duration) *LoadBalancer {
	return &LoadBalancer{
		replicas:            replicas,
		strategy:            strategy,
		healthcheckInterval: healthcheckInterval,
	}
}

type LoadBalancer struct {
	replicas            []string
	healthyReplicas     []string
	strategy            balancingStrategy
	healthcheckInterval time.Duration
}

// StartHealthchecks regularly checks which replicas are healthy and populate healthyReplicas
func (lb *LoadBalancer) StartHealthchecks() {
	lb.healthyReplicas = healthyHosts(lb.replicas)

	go func() {

		for {
			select {
			case <-time.After(lb.healthcheckInterval):
				lb.healthyReplicas = healthyHosts(lb.replicas)
			}
		}

	}()
}

func healthyHosts(hosts []string) []string {
	healthy := make([]string, 0)

	for _, host := range hosts {
		isHealthy, _ := healthCheck("http://" + host)

		if isHealthy {
			healthy = append(healthy, host)
		}

		fmt.Printf("%s healthy: %v\n", host, isHealthy)

	}

	return healthy
}

func healthCheck(host string) (bool, error) {

	response, err := http.Get(host + "/health")

	if err != nil {
		return false, err
	}

	return response.StatusCode == http.StatusOK, nil
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	lb.strategy.GetTarget(r, lb.healthyReplicas, func(host string) {
		fmt.Printf("Target: %s\n", host)
		err := reverse_proxy.Forward(w, r, host)
		if err != nil {
			// TODO: mark target host as unhealthy
			panic(err)
		}
	})
}
