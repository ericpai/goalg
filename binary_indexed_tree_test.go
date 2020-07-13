package goalg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBIT(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	t.Run("case=no-mod", func(t *testing.T) {
		bit := NewBIT(data, 0)
		assert.Equal(t, 1, bit.PrefixSum(0))
		assert.Equal(t, 28, bit.PrefixSum(6))
		assert.Equal(t, 105, bit.PrefixSum(13))

		assert.Equal(t, 104, bit.RangeSum(1, 13))
		assert.Equal(t, 20, bit.RangeSum(1, 5))
		assert.Equal(t, 34, bit.RangeSum(6, 9))
		assert.Equal(t, 8, bit.RangeSum(7, 7))
		assert.Equal(t, 36, bit.RangeSum(0, 7))

		bit.Update(7, 10) // data := []int{1, 2, 3, 4, 5, 6, 7, 18, 9, 10, 11, 12, 13, 14}

		assert.Equal(t, 1, bit.PrefixSum(0))
		assert.Equal(t, 28, bit.PrefixSum(6))
		assert.Equal(t, 115, bit.PrefixSum(13))

		assert.Equal(t, 114, bit.RangeSum(1, 13))
		assert.Equal(t, 20, bit.RangeSum(1, 5))
		assert.Equal(t, 44, bit.RangeSum(6, 9))
		assert.Equal(t, 18, bit.RangeSum(7, 7))
		assert.Equal(t, 46, bit.RangeSum(0, 7))
	})

	t.Run("case=with-mod", func(t *testing.T) {
		bit := NewBIT(data, 13)
		assert.Equal(t, 1, bit.PrefixSum(0))
		assert.Equal(t, 2, bit.PrefixSum(6))
		assert.Equal(t, 1, bit.PrefixSum(13))

		assert.Equal(t, 0, bit.RangeSum(1, 13))
		assert.Equal(t, 7, bit.RangeSum(1, 5))
		assert.Equal(t, 8, bit.RangeSum(6, 9))
		assert.Equal(t, 8, bit.RangeSum(7, 7))
		assert.Equal(t, 10, bit.RangeSum(0, 7))

		bit.Update(7, -10) // data := []int{1, 2, 3, 4, 5, 6, 7, -2, 9, 10, 11, 12, 13, 14}

		assert.Equal(t, 1, bit.PrefixSum(0))
		assert.Equal(t, 2, bit.PrefixSum(6))
		assert.Equal(t, 4, bit.PrefixSum(13))

		assert.Equal(t, 3, bit.RangeSum(1, 13))
		assert.Equal(t, 7, bit.RangeSum(1, 5))
		assert.Equal(t, 11, bit.RangeSum(6, 9))
		assert.Equal(t, 11, bit.RangeSum(7, 7))
		assert.Equal(t, 0, bit.RangeSum(0, 7))
	})
}
