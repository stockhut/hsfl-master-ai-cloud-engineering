package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/fun"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/load-balancer"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/load-balancer/config"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/load-balancer/strategies/round_robin"
	"log"
	"net/http"
	"time"
)

func main() {

	fmt.Println("Starting load balancer")

	configFile := "config.yaml"
	cfg, err := config.FromFile(configFile)
	if err != nil {
		log.Fatalf("Failed to read config file %s: %s", configFile, err)
	}

	ctx := context.Background()

	var replicas []string

	replicas = append(replicas, cfg.Hosts...)

	for _, image := range cfg.ContainerImages {
		running, err := getRunningReplicas(ctx, image)
		if err != nil {
			fmt.Printf("Failed to get running replicas for image %s: %s\n", image, err)
		}
		fmt.Printf("Found %d running instances for %s\n", len(running), image)

		containerHosts := fun.Map(running, func(ip string) string {
			return ip + ":" + cfg.ContainerPort
		})
		replicas = append(replicas, containerHosts...)
	}

	fmt.Println("Replicas:")
	for _, r := range replicas {
		fmt.Println(r)
	}
	fmt.Println("---")

	lb := load_balancer.New(replicas, round_robin.New(), 10*time.Second)
	lb.StartHealthchecks()

	panic(http.ListenAndServe(cfg.Listen, lb))

}

// / getRunningReplicas returns a slice of all container ip addresses running the given image.
// / (image includes name and version)
func getRunningReplicas(ctx context.Context, image string) ([]string, error) {

	fmt.Printf("looking for %s containers\n", image)

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	replicas := make([]string, 0)
	for _, container := range containers {

		if container.Image == image {
			ip, err := getContainerIP(ctx, cli, container.ID)
			if err != nil {
				fmt.Printf("Failed to get IP for container %s: %s", container.ID, err)
			}
			replicas = append(replicas, ip)
		}
	}

	return replicas, nil
}

func getContainerIP(ctx context.Context, cli *client.Client, id string) (string, error) {
	info, err := cli.ContainerInspect(ctx, id)

	if err != nil {
		return "", err
	}

	return info.NetworkSettings.IPAddress, nil
}
