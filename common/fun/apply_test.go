package fun

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApply(t *testing.T) {

	t.Run("should call functions from left to right", func(t *testing.T) {

		a := func(s string) string {
			return "a" + s
		}

		b := func(s string) string {
			return "b" + s
		}

		c := func(s string) string {
			return "c" + s
		}

		abc := Apply(a, b, c)

		result := abc("")

		expected := a(b(c("")))

		assert.Equal(t, "abc", expected)
		assert.Equal(t, expected, result)
	})

	t.Run("should be lazy", func(t *testing.T) {

		called := false

		a := func(t any) any {
			called = true
			return t
		}

		b := func(t any) any {
			called = true
			return t
		}

		// notice this does not call the resulting function
		Apply(a, b)

		assert.Equal(t, false, called)
	})

}
