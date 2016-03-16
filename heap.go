package dpll

import "container/heap"

// maxActiveHeap is a heap.Interface that prioritizes variables by max
// activity.
type maxActiveHeap struct {
	act   *[]float64
	index []int
	vars  []Var
}

var _ heap.Interface = &maxActiveHeap{}

func newMaxActiveHeap(act *[]float64) *maxActiveHeap {
	return &maxActiveHeap{act: act}
}

func (h *maxActiveHeap) Len() int {
	return len(h.vars)
}

func (h *maxActiveHeap) Less(i, j int) bool {
	return (*h.act)[h.vars[i]] > (*h.act)[h.vars[j]]
}

func (h *maxActiveHeap) Swap(i, j int) {
	h.index[h.vars[i]] = j
	h.index[h.vars[j]] = i
	h.vars[i], h.vars[j] = h.vars[j], h.vars[i]
}

// Push assumes that x is not already in h
func (h *maxActiveHeap) Push(x interface{}) {
	v := x.(Var)
	h.extend(v)
	h.index[v] = len(h.vars)
	h.vars = append(h.vars, v)
}

func (h *maxActiveHeap) Pop() interface{} {
	x := h.vars[len(h.vars)-1]
	h.vars = h.vars[:len(h.vars)-1]
	h.index[x] = -1
	return x
}

func (h *maxActiveHeap) extend(v Var) {
	if int(v) >= len(h.index) {
		n := 2 * int(v)
		index := make([]int, n)
		for i := len(h.index); i < n; i++ {
			index[i] = -1
		}
		copy(index, h.index)
		h.index = index
	}
}

func (h *maxActiveHeap) RemoveMin() Var {
	return heap.Pop(h).(Var)
}

func (h *maxActiveHeap) Contains(v Var) bool {
	if int(v) < len(h.index) {
		return h.index[v] >= 0
	}
	return false
}

func (h *maxActiveHeap) Decrease(v Var) {
	heap.Fix(h, h.index[v])
}

func (h *maxActiveHeap) Rebuild(vs []Var) {
	h.vars = append(h.vars[:0], vs...)
	h.index = h.index[:0]
	for i := range vs {
		h.extend(vs[i])
		h.index[vs[i]] = i
	}
	heap.Init(h)
}
