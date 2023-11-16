package environment

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetRequiredEnvVar(t *testing.T) {

	t.Run("should return environment variable", func(t *testing.T) {
		t.Setenv("TEST_VARIABLE", "FOO")

		v := GetRequiredEnvVar("TEST_VARIABLE")

		assert.Equal(t, "FOO", v)
	})

}
