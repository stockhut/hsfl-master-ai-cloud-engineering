package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestFromFile(t *testing.T) {

	config := configFile{
		Services: map[string]configItem{
			"foo": configItem{
				Route:      "/foo",
				TargetHost: "foo.example.org",
			},
		},
	}

	b, err := yaml.Marshal(config)
	fmt.Println(err)
	fmt.Println(string(b))
}
func TestFromBytes(t *testing.T) {

	input :=
		`services:
  foo:
    route: /foo
    targetHost: foo.example.org
`

	config, err := fromBytes([]byte(input))

	assert.Nil(t, err)

	assert.Len(t, config.Services, 1)
}
