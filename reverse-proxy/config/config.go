package config

import (
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/reverse-proxy"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Services []reverse_proxy.Service
}

type configFile struct {
	Services map[string]configItem `yaml:"services"`
}
type configItem struct {
	Route      string `yaml:"route"`
	TargetHost string `yaml:"targetHost"`
}

func configItemToService(name string, item configItem) reverse_proxy.Service {
	return reverse_proxy.Service{
		Name:       name,
		Route:      item.Route,
		TargetHost: item.TargetHost,
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

	var services = make([]reverse_proxy.Service, 0)

	for name, serviceConfig := range config.Services {
		services = append(services, configItemToService(name, serviceConfig))
	}
	return Config{
		Services: services,
	}, nil
}
