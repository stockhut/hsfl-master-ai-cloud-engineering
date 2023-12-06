package main

import (
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/loadtest/config"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/loadtest/worker"
	"log"
	"time"
)

func main() {

	configPath := "load.config.json"
	conf, err := config.FromFS(configPath)
	if err != nil {
		log.Fatalf("Failed to load config file %s: %s", configPath, err)
	}

	numWorkers := conf.Users

	job := worker.JobFunc[int](func() int {
		return 2
	})
	p := worker.NewRampedPool[int](numWorkers, job)

	go p.Start()

	timeout := time.After(time.Duration(conf.Duration) * time.Millisecond)

	<-timeout
	p.Stop()

}
