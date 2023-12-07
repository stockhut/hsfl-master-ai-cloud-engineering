package main

import (
	"fmt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/loadtest/config"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/loadtest/worker"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	configPath := "load.config.json"
	conf, err := config.FromFS(configPath)
	if err != nil {
		log.Fatalf("Failed to load config file %s: %s", configPath, err)
	}

	numWorkers := conf.Users

	jobFactory := worker.JobFactoryFunc[any](func() worker.Job[any] {
		return worker.JobFunc[any](func() any {
			target := randomItemFromSlice(conf.Targets)

			_, err := http.Get(target)
			if err != nil {
				log.Printf("Request to target %s failed: %s", target, err)
			}
			//time.Sleep(1 * time.Second)
			fmt.Println(target)

			return true
		})
	})
	p := worker.NewRampedPool[any](numWorkers, jobFactory)

	go p.Start()

	timeout := time.After(time.Duration(conf.Duration) * time.Millisecond)

	<-timeout
	p.Stop()

}

func randomItemFromSlice[T any](ts []T) T {
	i := rand.Intn(len(ts))
	return ts[i]
}
