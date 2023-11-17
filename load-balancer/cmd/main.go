package main

import (
	"fmt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/load-balancer"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/load-balancer/strategies/round_robin"
	"net/http"
	"time"
)

func main() {

	fmt.Println("Starting load balancer")

	replicas := []string{
		"localhost:8081",
		//"localhost:9001",
		//"localhost:9002",
		"localhost:9003",
	}

	lb := load_balancer.New(replicas, round_robin.New(), 10*time.Second)
	lb.StartHealthchecks()

	panic(http.ListenAndServe(":5001", lb))

}
