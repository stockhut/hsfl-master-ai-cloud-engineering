package orchestration

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/fun"
	"log"
)

type ServiceContainerConfig struct {
	Name        string
	Image       string
	Environment map[string]string
	Mounts      map[string]string
	MinReplicas int
	Port        int
}

type ServiceContainer struct {
	// container id assigned by the docker daemon
	ID          string
	Name        string
	Ip          string
	StoppedChan <-chan container.WaitResponse
	ErrorChan   <-chan error
}

func createContainer(cli *client.Client, sc ServiceContainerConfig) (container.CreateResponse, error) {

	environmentVariables := fun.MapToSlice(sc.Environment, func(name string, value string) string {
		return name + "=" + value
	})

	containerConfig := container.Config{
		Image: sc.Image,
		ExposedPorts: map[nat.Port]struct{}{
			nat.Port("8080/tcp"): {},
		},
		Env: environmentVariables,
	}

	mounts := fun.MapToSlice(sc.Mounts, func(hostDir string, containerDir string) mount.Mount {
		return mount.Mount{
			Type:   "bind",
			Source: hostDir,
			Target: containerDir,
		}
	})

	// TODO: move port config to ServiceContainerConfig
	hostConfig := container.HostConfig{
		Mounts:     mounts,
		AutoRemove: true,
	}

	return cli.ContainerCreate(context.Background(),
		&containerConfig,
		&hostConfig,
		nil, nil, "")
}

func CreateAndStartContainer(cli *client.Client, sc ServiceContainerConfig) (ServiceContainer, error) {

	resp, err := createContainer(cli, sc)
	if err != nil {
		return ServiceContainer{}, err
	}

	err = cli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{})
	if err != nil {
		return ServiceContainer{}, err
	}

	info, err := cli.ContainerInspect(context.Background(), resp.ID)
	if err != nil {
		return ServiceContainer{}, err
	}

	ip := info.NetworkSettings.IPAddress

	statusCh, errCh := cli.ContainerWait(context.Background(), resp.ID, container.WaitConditionNotRunning)

	return ServiceContainer{
		Name:        info.Name,
		ID:          resp.ID,
		Ip:          ip,
		StoppedChan: statusCh,
		ErrorChan:   errCh,
	}, nil
}

func StopAll(cli *client.Client, containers []ServiceContainer) {
	for _, c := range containers {
		err := cli.ContainerStop(context.Background(), c.ID, container.StopOptions{})
		if err != nil {
			log.Printf("⚠️ Failed to stop container: %s:%s\n", c.Name, err)
		} else {
			log.Printf("🛑 Stopped container %s\n", c.Name)
		}
	}
}
