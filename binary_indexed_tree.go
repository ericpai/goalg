package goalg

// BIT is the binary indexed tree/Fenwick tree struct.
//
// It's used to calculate prefix sum in logarithmic time with any data updates at any time.
//
// Ref: https://blog.csdn.net/Yaokai_AssultMaster/article/details/79492190.
type BIT struct {
	data []int // 1-based array
	mod  int
}

// NewBIT inits a BIT instance with original data and calculates the prefix sum % mod.
//
// If mod is 0, no mod action will be executed in the future.
//
// Time complexity: O(n).
func NewBIT(data []int, mod int) *BIT {
	res := make([]int, len(data)+1)
	for i := 0; i < len(data); i++ {
		res[i+1] = data[i]
	}
	for i := 1; i < len(res); i++ {
		j := i + (i & -i)
		if j < len(res) {
			res[j] += res[i]
			if mod > 0 {
				res[j] %= mod
			}
		}
	}
	return &BIT{data: res, mod: mod}
}

// RangeSum returns the range sum of subarray [start, end].
//
// Time complexity: O(logn).
func (b *BIT) RangeSum(start, end int) int {
	res := b.PrefixSum(end) - b.PrefixSum(start-1)
	if b.mod > 0 {
		if res < 0 {
			res += b.mod
		}
		res %= b.mod
	}
	return res
}

// PrefixSum returns the range sum of subarray [0, end].
//
// Time complexity: O(logn).
func (b *BIT) PrefixSum(end int) int {
	var s int
	for end++; end > 0; end -= end & -end {
		s += b.data[end]
		if b.mod > 0 {
			s %= b.mod
		}
	}
	return s
}

// Update updates the element in idx with delta
//
// Time complexity: O(logn).
func (b *BIT) Update(idx, delta int) {
	for idx++; idx < len(b.data); idx += idx & (-idx) {
		b.data[idx] += delta
		if b.mod > 0 {
			b.data[idx] %= b.mod
		}
	}
}
