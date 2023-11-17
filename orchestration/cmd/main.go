package main

import (
	"fmt"
	"github.com/docker/docker/client"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/orchestration"
	"os"
	"os/signal"
)

func main() {

	serviceConfigs := []orchestration.ServiceContainerConfig{
		{
			Name:  "recipe",
			Image: "recipe:latest",
			Environment: map[string]string{
				"JWT_PUBLIC_KEY": "/keys/jwt_public_key.key",
			},
			Mounts: map[string]string{
				"/home/f/Projects/Hochschule/hsfl-master-ai-cloud-engineering/authentication": "/keys",
			},
			MinReplicas: 2,
		},
	}

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers := make([]orchestration.ServiceContainer, 0)

	defer func() {
		p := recover()
		if p != nil {
			fmt.Printf("Panic: %s\n", p)
			fmt.Print("Recovered, stopping containers")
			orchestration.StopAll(cli, containers)
		}
	}()

	for _, serviceConfig := range serviceConfigs {

		for i := 0; i < serviceConfig.MinReplicas; i++ {

			serviceContainer, err := orchestration.CreateAndStartContainer(cli, serviceConfig)
			if err != nil {
				fmt.Printf("Failed to start service container: %s\n", err)
			}
			fmt.Printf("started container %v\n", serviceContainer)
			containers = append(containers, serviceContainer)
		}
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Kill, os.Interrupt)
	go func() {

		sig := <-sigs
		fmt.Printf("Received signal %s\n", sig)
		orchestration.StopAll(cli, containers)
		os.Exit(1)
	}()
}
