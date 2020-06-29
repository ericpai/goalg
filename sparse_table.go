package goalg

// SparseTable is fast in multiple range maximum/minimum queries in an immutable array.
//
// Ref: https://cp-algorithms.com/data_structures/sparse-table.html.
type SparseTable struct {
	t  [][]int
	f  func(d ...int) int
	n  int
	lg []int
}

// NewSparseTable creates a sparse table with the range rule l.
//
// NOTICE: Parameter f must be an idempotent function: https://en.wikipedia.org/wiki/Idempotence. if you want to do range queries in O(1). If not, wrong query results will be returned.
//
// Time complexity: O(nlogn).
func NewSparseTable(data []int, f func(d ...int) int) *SparseTable {
	n := len(data)
	t := make([][]int, n+1)
	for i := 0; i < n; i++ {
		t[i] = append(t[i], data[i])
	}
	for i := 1; 1<<i <= n && 1<<i > 0; i++ {
		for j := 0; j+(1<<i) <= n && j+(1<<i) > 0; j++ {
			t[j] = append(t[j], f(t[j][i-1], t[j+(1<<(i-1))][i-1]))
		}
	}
	lg := make([]int, n+1)
	lg[1] = 0
	for i := 2; i <= n; i++ {
		lg[i] = lg[i/2] + 1
	}
	return &SparseTable{
		t:  t,
		f:  f,
		n:  n,
		lg: lg,
	}
}

// QueryRangeFast queries the range [left, right] with function f.
//
// NOTICE: Only available if f is an idempotent function: https://en.wikipedia.org/wiki/Idempotence.
//
// Time complexity: O(1).
func (s *SparseTable) QueryRangeFast(left, right int) int {
	l := right - left + 1
	return s.f(s.t[left][s.lg[l]], s.t[right-(1<<s.lg[l])+1][s.lg[l]])
}

// QueryRange queries the range [left, right] with function f.
//
// Time complexity: O(logn).
func (s *SparseTable) QueryRange(left, right int) int {
	l := right - left + 1
	res := s.t[left][s.lg[l]]
	for left += 1 << s.lg[l]; left <= right; left += 1 << s.lg[l] {
		l = right - left + 1
		res = s.f(res, s.t[left][s.lg[l]])
	}
	return res
}
