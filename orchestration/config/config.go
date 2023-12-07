package config

import (
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/orchestration"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Services []orchestration.ServiceContainerConfig
}
type configFile struct {
	Services map[string]serviceItem `yaml:"services"`
}

type serviceItem struct {
	Image        string            `yaml:"image"`
	Env          map[string]string `yaml:"env"`
	MinInstances int               `yaml:"minInstances"`
	Mounts       map[string]string `yaml:"mounts"`
}

func serviceItemToContainerConfig(name string, item serviceItem) orchestration.ServiceContainerConfig {
	return orchestration.ServiceContainerConfig{
		Name:        name,
		Image:       item.Image,
		Environment: item.Env,
		Mounts:      item.Mounts,
		MinReplicas: item.MinInstances,
	}
}

func FromFile(path string) (Config, error) {

	content, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	return fromBytes(content)
}

func fromBytes(input []byte) (Config, error) {
	var config configFile
	err := yaml.Unmarshal(input, &config)

	if err != nil {
		return Config{}, err
	}

	var services = make([]orchestration.ServiceContainerConfig, 0)

	for name, serviceConfig := range config.Services {
		services = append(services, serviceItemToContainerConfig(name, serviceConfig))
	}
	return Config{
		Services: services,
	}, nil
}
