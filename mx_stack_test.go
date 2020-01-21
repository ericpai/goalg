package goalg_test

import (
	"testing"

	"github.com/ericpai/goalg"
	"github.com/stretchr/testify/assert"
)

func maxFunc(i, j interface{}) bool {
	return i.(int) > j.(int)
}

func TestMXStack(t *testing.T) {
	mxs := goalg.NewMXimumStack(maxFunc)

	assert.Equal(t, nil, mxs.MX())
	assert.Equal(t, 0, mxs.Len())
	assert.Equal(t, nil, mxs.Top())
	assert.Equal(t, nil, mxs.Pop())

	mxs.Push(1)
	assert.Equal(t, 1, mxs.MX())
	assert.Equal(t, 1, mxs.Len())
	assert.Equal(t, 1, mxs.Top())
	assert.Equal(t, 1, mxs.Pop())

	mxs.Push(1)
	mxs.Push(10)
	assert.Equal(t, 10, mxs.MX())
	assert.Equal(t, 2, mxs.Len())
	assert.Equal(t, 10, mxs.Top())
	assert.Equal(t, 10, mxs.Pop())

	mxs.Push(10)
	mxs.Push(9)
	assert.Equal(t, 10, mxs.MX())
	assert.Equal(t, 3, mxs.Len())
	assert.Equal(t, 9, mxs.Top())
	assert.Equal(t, 9, mxs.Pop())
}
