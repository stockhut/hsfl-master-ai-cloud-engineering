package fun

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCount(t *testing.T) {
	t.Run("Count counts items matching pred", func(t *testing.T) {
		t.Run("Empty slice", func(t *testing.T) {

			items := []bool{}

			assert.Equal(t, 0, Count(items, func(b bool) bool {
				return b
			}))
		})

		t.Run("Non empty slice", func(t *testing.T) {

			items := []bool{true, false, false, true}

			assert.Equal(t, 2, Count(items, func(b bool) bool {
				return b
			}))
		})
	})
}
