package load_balancer

import (
	"fmt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/reverse-proxy"
	"net/http"
	"slices"
	"sync"
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
		healthyLock:         &sync.RWMutex{},
	}
}

type LoadBalancer struct {
	replicas        []string
	healthyReplicas []string
	// healthyLock guards access to the healthyReplicas slice
	healthyLock         *sync.RWMutex
	strategy            balancingStrategy
	healthcheckInterval time.Duration
}

// StartHealthchecks regularly checks which replicas are healthy and populate healthyReplicas
func (lb *LoadBalancer) StartHealthchecks() {

	lb.healthyLock.Lock()
	lb.healthyReplicas = healthyHosts(lb.replicas)
	lb.healthyLock.Unlock()

	go func() {

		for {
			select {
			case <-time.After(lb.healthcheckInterval):
				lb.healthyLock.Lock()
				lb.healthyReplicas = healthyHosts(lb.replicas)
				lb.healthyLock.Unlock()
			}
		}

	}()

}

// markUnhealthy marks the given host as unhealthy. Acquires healthyLock
func (lb *LoadBalancer) markUnhealthy(host string) {

	lb.healthyLock.Lock()
	for i, h := range lb.healthyReplicas {
		if h == host {
			lb.healthyReplicas = append(lb.healthyReplicas[:i], lb.healthyReplicas[i+1:]...)
		}
	}
	lb.healthyLock.Unlock()
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

	// create a copy of the slice to avoid race conditions with healthcheck updates on lb.healthyReplicas
	lb.healthyLock.RLock()
	h := slices.Clone(lb.healthyReplicas)
	lb.healthyLock.RUnlock()

	lb.strategy.GetTarget(r, h, func(host string) {
		fmt.Printf("Target: %s\n", host)
		err := reverse_proxy.Forward(w, r, host)
		if err != nil {
			lb.markUnhealthy(host)
		}
	})
}
