package goalg

// Heap , aka PriorityQueue
//
// The top of the heap is the least element sorted by lessFunc function
type Heap struct {
	data []interface{}
	less func(i, j interface{}) bool
}

// NewHeap initials a minimum heap with data and the comparaing function lessFunc.
//
// Time complexity: O(n).
func NewHeap(data []interface{}, lessFunc func(i, j interface{}) bool) *Heap {

	h := &Heap{
		data: make([]interface{}, len(data)),
		less: lessFunc,
	}
	for i := 0; i < h.Len(); i++ {
		h.data[i] = data[i]
	}
	for i := (len(data) - 1) >> 1; i >= 0; i-- {
		h.down(i)
	}
	return h
}

// Len returns the length of elements in the heap.
//
// Time complexity: O(1).
func (h *Heap) Len() int {
	return len(h.data)
}

// Push adds e to the heap and heapify itself automatically
//
// Time complexity: O(logn).
func (h *Heap) Push(e interface{}) {
	h.data = append(h.data, e)
	var p int
	for idx := h.Len() - 1; idx > 0; idx = p {
		p = (idx - 1) >> 1
		if h.less(h.data[p], h.data[idx]) {
			break
		}
		h.data[idx], h.data[p] = h.data[p], h.data[idx]
	}
}

// Pop removes the top element from the heap and heapify itself automatically.
//
// It h.Len() == 0, nil will be returned.
//
// Time complexity: O(logn).
func (h *Heap) Pop() interface{} {
	if h.Len() == 0 {
		return nil
	}
	h.data[0], h.data[h.Len()-1] = h.data[h.Len()-1], h.data[0]
	ret := h.data[h.Len()-1]
	h.data[h.Len()-1] = nil
	h.data = h.data[:h.Len()-1]
	h.down(0)
	return ret
}

// Top gets the top element from the heap.
//
// It h.Len() == 0, nil will be returned.
//
// Time complexity: O(1).
func (h *Heap) Top() interface{} {
	if h.Len() == 0 {
		return nil
	}
	return h.data[0]
}

// down ensures the heap properity of the sub heap from idx.
//
// Time complexity: O(height(heap) - depth(idx))
func (h *Heap) down(idx int) {
	var minIdx, right, left int
	for {
		minIdx = idx
		left = (idx << 1) + 1
		right = (idx << 1) + 2
		if left < len(h.data) && left >= 0 && h.less(h.data[left], h.data[minIdx]) {
			minIdx = left
		}
		if right < len(h.data) && right >= 0 && h.less(h.data[right], h.data[minIdx]) {
			minIdx = right
		}
		if minIdx == idx {
			break
		}
		h.data[idx], h.data[minIdx] = h.data[minIdx], h.data[idx]
		idx = minIdx
	}
}
