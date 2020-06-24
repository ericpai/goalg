package goalg

// DSU is the short for Disjoint Set Union data structure.
//
// Ref: https://cp-algorithms.com/data_structures/disjoint_set_union.html
type DSU struct {
	parent *DSU
	size   int
}

// NewDSU initializes a distinct DSU with rank 1.
func NewDSU() *DSU {
	p := &DSU{size: 1}
	p.parent = p
	return p
}

// Union unions two DSU together by their size.
//
// Time complexity: (O(1))
func (d *DSU) Union(u *DSU) {
	dp := d.Find()
	up := u.Find()
	if dp != up {
		if dp.size < up.size {
			dp, up = up, dp
		}
		up.parent = dp
		dp.size += up.size
	}
}

// Find returns the root DSU of the current one with path compression.
//
// Time complexity: (O(1))
func (d *DSU) Find() *DSU {
	if d.parent == d {
		return d
	}
	res := d.parent.Find()
	d.parent = res
	return res
}
