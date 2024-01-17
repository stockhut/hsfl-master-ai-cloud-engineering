package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Listen          string   `yaml:"listen"`
	Hosts           []string `yaml:"hosts"`
	ContainerPort   string   `yaml:"containerPort"`
	ContainerImages []string `yaml:"images"`
}

func FromFile(path string) (Config, error) {

	content, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	return fromBytes(content)
}

func fromBytes(input []byte) (Config, error) {
	var config Config
	err := yaml.Unmarshal(input, &config)

	if err != nil {
		return Config{}, err
	}

	return config, nil
}
