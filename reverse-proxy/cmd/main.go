package main

import (
	"fmt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/reverse-proxy"
	"log"
	"net/http"
	"os"
)

const Port int = 5000
const Host string = "localhost"

func main() {

	services := []reverse_proxy.Service{
		{
			Name:       "recipe service",
			Route:      "/api/v1/recipe",
			TargetHost: "localhost:8081",
		},
		{
			Name:       "web-service",
			Route:      "/",
			TargetHost: "localhost:3000",
		},
	}

	logger := log.New(os.Stdout, "", 0)

	proxy := reverse_proxy.New(logger, services)

	addr := fmt.Sprintf("%s:%d", Host, Port)
	log.Fatal(http.ListenAndServe(addr, proxy))
}
