package goalg_test

import (
	"testing"

	"github.com/ericpai/goalg"
	"github.com/stretchr/testify/assert"
)

func TestSparseTable(t *testing.T) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	st := goalg.NewSparseTable(data, func(a ...int) int {
		var res int
		for _, v := range a {
			res += v
		}
		return res
	})
	assert.Equal(t, 28, st.QueryRange(0, 7))
	assert.Equal(t, 4950, st.QueryRange(0, 99))
	assert.Equal(t, 46, st.QueryRange(10, 13))
	assert.Equal(t, 10, st.QueryRange(10, 10))
	assert.Equal(t, 75, st.QueryRange(10, 15))

	st = goalg.NewSparseTable(data, func(a ...int) int {
		res := a[0]
		for _, v := range a {
			if v < res {
				res = v
			}
		}
		return res
	})
	assert.Equal(t, 0, st.QueryRange(0, 7))
	assert.Equal(t, 0, st.QueryRange(0, 99))
	assert.Equal(t, 10, st.QueryRange(10, 13))
	assert.Equal(t, 10, st.QueryRange(10, 10))
	assert.Equal(t, 10, st.QueryRange(10, 15))
	assert.Equal(t, 0, st.QueryRangeFast(0, 7))
	assert.Equal(t, 0, st.QueryRangeFast(0, 99))
	assert.Equal(t, 10, st.QueryRangeFast(10, 13))
	assert.Equal(t, 10, st.QueryRangeFast(10, 10))
	assert.Equal(t, 10, st.QueryRangeFast(10, 15))
}
