package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/orchestration"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/orchestration/config"
	"io"
	"log"
	"os"
	"os/signal"
)

func main() {

	configFilePath := "config.yml"
	c, err := config.FromFile(configFilePath)
	if err != nil {
		log.Fatalf("⚠️ Failed to load config file %s: %s\n", configFilePath, err)
	}

	log.Printf("🔧 Loaded config file %s", configFilePath)
	log.Printf("🔧 %d services configured", len(c.Services))

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers := make([]orchestration.ServiceContainer, 0)

	defer func() {
		p := recover()
		if p != nil {
			fmt.Printf("⚠️ Panic: %s\n", p)
			fmt.Print(" ⚠️ Recovered, stopping containers")
			orchestration.StopAll(cli, containers)
		}
	}()

	for _, serviceConfig := range c.Services {

		log.Printf("🏗️ Pulling image %s\n", serviceConfig.Image)
		out, err := cli.ImagePull(context.TODO(), serviceConfig.Image, types.ImagePullOptions{})
		if err != nil {
			log.Printf("⚠️ %s", err)
		} else {
			// wait until image is pulled
			io.ReadAll(out)
		}
		for i := 0; i < serviceConfig.MinReplicas; i++ {

			serviceContainer, err := orchestration.CreateAndStartContainer(cli, serviceConfig)
			if err != nil {
				log.Printf("⚠️ Failed to start service container: %s\n", err)
			}
			log.Printf("🚀 Started container '%s' with IP %s\n", serviceContainer.Name, serviceContainer.Ip)
			containers = append(containers, serviceContainer)
		}
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Kill, os.Interrupt)
	func() {

		sig := <-sigs
		log.Printf("🛑 Received %s signal\n", sig)
		orchestration.StopAll(cli, containers)
		os.Exit(1)
	}()
}
