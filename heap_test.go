package goalg_test

import (
	"testing"

	"github.com/ericpai/goalg"
	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	h := goalg.NewHeap([]interface{}{5, 4, 3, 2, 1}, func(i, j interface{}) bool {
		return i.(int) < j.(int)
	})
	assert.Equal(t, 5, h.Len())
	assert.Equal(t, 1, h.Top())
	assert.Equal(t, 1, h.Pop())
	assert.Equal(t, 4, h.Len())
	assert.Equal(t, 2, h.Top())
	assert.Equal(t, 2, h.Pop())
	h.Push(1)
	assert.Equal(t, 4, h.Len())
	assert.Equal(t, 1, h.Top())

	h.Push(6)
	assert.Equal(t, 5, h.Len())
	assert.Equal(t, 1, h.Top())
	assert.Equal(t, 1, h.Pop())

	assert.Equal(t, 4, h.Len())
	assert.Equal(t, 3, h.Top())
	assert.Equal(t, 3, h.Pop())

	assert.Equal(t, 3, h.Len())
	assert.Equal(t, 4, h.Top())
	assert.Equal(t, 4, h.Pop())

	assert.Equal(t, 2, h.Len())
	assert.Equal(t, 5, h.Top())
	assert.Equal(t, 5, h.Pop())

	assert.Equal(t, 1, h.Len())
	assert.Equal(t, 6, h.Top())
	assert.Equal(t, 6, h.Pop())

	assert.Equal(t, 0, h.Len())
	assert.Equal(t, nil, h.Top())
	assert.Equal(t, nil, h.Pop())
}
