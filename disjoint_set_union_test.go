package goalg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDSU(t *testing.T) {
	d1 := NewDSU()
	d2 := NewDSU()
	d3 := NewDSU()
	d4 := NewDSU()
	d5 := NewDSU()
	assert.Same(t, d1, d1.Find())
	assert.Equal(t, 1, d1.size)
	assert.True(t, d1 != d2)
	d1.Union(d2)
	assert.Same(t, d1, d1.Find())
	assert.Same(t, d1, d2.Find())
	assert.Equal(t, 2, d1.size)
	assert.Equal(t, 1, d2.size)

	d3.Union(d4)
	d3.Union(d5)
	assert.Equal(t, 3, d3.size)
	assert.Same(t, d3, d5.parent)
	assert.Same(t, d3, d4.parent)
	d2.Union(d3)
	assert.Equal(t, 5, d3.size)
	assert.Equal(t, 2, d1.size)
}
