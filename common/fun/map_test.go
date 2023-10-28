package fun

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {

	a := []int{1, 2, 3, 4}

	b := Map(a, func(a int) int {
		return a + a
	})

	assert.Equal(t, []int{2, 4, 6, 8}, b)
}
