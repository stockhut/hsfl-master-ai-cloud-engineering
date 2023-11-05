package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/stockhut/hsfl-master-ai-cloud-engineering/reverse-proxy"
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

	router := http.NewServeMux()
	router.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {

			handled := false
			for _, service := range services {
				if strings.HasPrefix(r.URL.Path, service.Route) {
					handled = true

					logger.Printf("%s => %s (%s)\n", r.URL, service.Name, service.TargetHost)
					err := reverse_proxy.Forward(w, r, service)
					if err != nil {
						w.WriteHeader(http.StatusInternalServerError)
						logger.Printf("Failed to forward request: %s", err)
					}
					break
				}
			}

			if !handled {
				logger.Printf("No matching service for %v\n", r.URL)
			}
		})

	addr := fmt.Sprintf("%s:%d", Host, Port)
	log.Fatal(http.ListenAndServe(addr, router))
}
