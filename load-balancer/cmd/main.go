package main

import (
	"fmt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/load-balancer"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/load-balancer/strategies/least_connections"
	"net/http"
)

func main() {

	fmt.Println("Starting load balancer")

	replicas := []string{
		"localhost:9001",
		"localhost:9002",
		"localhost:9003",
	}

	lb := load_balancer.New(replicas, least_connections.New())

	panic(http.ListenAndServe(":5001", lb))

}
