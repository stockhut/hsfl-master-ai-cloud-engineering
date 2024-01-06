package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type LoadTest struct {
	ResponseStats bool     `yaml:"responseStats"`
	Host          string   `yaml:"host"`
	Targets       []string `yaml:"targets"`
	Headers       Headers  `yaml:"headers"`
	Phases        []Phase  `yaml:"phases"`
}

type Phase struct {
	Rps      int           `yaml:"rps"`
	Rampup   time.Duration `yaml:"rampup"`
	Duration time.Duration `yaml:"duration"`
}

type Headers map[string]string

func FromFile(path string) (LoadTest, error) {

	content, err := os.ReadFile(path)
	if err != nil {
		return LoadTest{}, err
	}

	return fromBytes(content)
}

func fromBytes(input []byte) (LoadTest, error) {
	var config LoadTest
	err := yaml.Unmarshal(input, &config)

	if err != nil {
		return LoadTest{}, err
	}

	return config, err
}
