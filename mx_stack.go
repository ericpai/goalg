package goalg

import "container/list"

type pair struct {
	key, mx interface{}
}

// MXimumStack is a stack with mximum(maximum or minimum) element in the stack recorded
// Ref: https://cp-algorithms.com/data_structures/stack_queue_modification.html
type MXimumStack struct {
	core *list.List
	less func(i, j interface{}) bool
}

// NewMXimumStack initialize a stack with the comparator lessFunc
func NewMXimumStack(lessFunc func(i, j interface{}) bool) *MXimumStack {
	return &MXimumStack{
		core: list.New(),
		less: lessFunc,
	}
}

// MX returns the mximum element in the stack
func (m *MXimumStack) MX() interface{} {
	if m.core.Len() == 0 {
		return nil
	}
	return m.core.Front().Value.(pair).mx
}

// Push pushes and records the MX record to the stack.
// Time complexity: (O(1))
func (m *MXimumStack) Push(d interface{}) {
	min := d
	if m.core.Len() > 0 && m.less(m.core.Front().Value.(pair).mx, d) {
		min = m.core.Front().Value.(pair).mx
	}
	m.core.PushFront(pair{
		key: d,
		mx:  min,
	})
}

// Pop pops the stack top element.
// Time complexity: (O(1))
func (m *MXimumStack) Pop() interface{} {
	if m.core.Len() == 0 {
		return nil
	}
	d := m.core.Front().Value.(pair)
	m.core.Remove(m.core.Front())
	return d.key
}

// Top returns the stack top element.
// Time complexity: (O(1))
func (m *MXimumStack) Top() interface{} {
	if m.core.Len() == 0 {
		return nil
	}
	return m.core.Front().Value.(pair).key
}

// Len returns the length of the stack.
// Time complexity: (O(1))
func (m *MXimumStack) Len() int {
	return m.core.Len()
}
