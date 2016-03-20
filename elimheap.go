// Copyright 2016 Bryan Matsuo
//
// Use of this software is governed by the MIT license.  A copy of the license
// agreement can be found in the LICENSE file distributed with this software.

package dpll

import "container/heap"

// elimQueue is a heap with special access methods.
type elimQueue minCostHeap

func newElimQueue(numocc *map[Lit]int) *elimQueue {
	return (*elimQueue)(newMinCostHeap(numocc))
}

func (q *elimQueue) h() *minCostHeap {
	return (*minCostHeap)(q)
}

func (q *elimQueue) Len() int {
	return q.h().Len()
}

func (q *elimQueue) Init() {
	q.h().Init()
}

// RemoveMin pops the least expensive variable from q and returns it.
func (q *elimQueue) RemoveMin() (v Var, ok bool) {
	return q.h().RemoveMin()
}

func (q *elimQueue) Contains(v Var) bool {
	return q.h().Contains(v)
}

func (q *elimQueue) Update(v Var) {
	q.h().Update(v)
}

func (q *elimQueue) Fix(v Var) {
	q.h().Fix(v)
}

func (q *elimQueue) Rebuild(vs []Var) {
	q.h().Rebuild(vs)
}

func (q *elimQueue) Push(v Var) {
	q.h().Push(v)
}

// minCostHeap is a heap.Interface that prioritizes variables by minimizing
// literal occurance uniformity, maximizing bias between negated and positive
// literal occurrences for a variable.
type minCostHeap struct {
	numocc *map[Lit]int
	index  []int
	vars   []Var
}

var _ heap.Interface = &minCostHeap{}

func newMinCostHeap(numocc *map[Lit]int) *minCostHeap {
	return &minCostHeap{numocc: numocc}
}

func (h *minCostHeap) Len() int {
	return len(h.vars)
}

func (h *minCostHeap) costSign(v Var, neg bool) uint64 {
	return uint64((*h.numocc)[Literal(v, neg)])
}

func (h *minCostHeap) cost(v Var) uint64 {
	return h.costSign(v, false) * h.costSign(v, true)
}

func (h *minCostHeap) Less(i, j int) bool {
	return h.cost(h.vars[i]) < h.cost(h.vars[j])
}

func (h *minCostHeap) Swap(i, j int) {
	h.index[h.vars[i]] = j
	h.index[h.vars[j]] = i
	h.vars[i], h.vars[j] = h.vars[j], h.vars[i]
}

// Push assumes that x is not already in h
func (h *minCostHeap) Push(x interface{}) {
	v := x.(Var)
	h.extend(v)
	h.index[v] = len(h.vars)
	h.vars = append(h.vars, v)
}

func (h *minCostHeap) Pop() interface{} {
	x := h.vars[len(h.vars)-1]
	h.vars = h.vars[:len(h.vars)-1]
	h.index[x] = -1
	return x
}

func (h *minCostHeap) extend(v Var) {
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

func (h *minCostHeap) RemoveMin() (v Var, ok bool) {
	if len(h.vars) == 0 {
		return 0, false
	}
	return heap.Pop(h).(Var), true
}

func (h *minCostHeap) Contains(v Var) bool {
	if int(v) < len(h.index) {
		return h.index[v] >= 0
	}
	return false
}

func (h *minCostHeap) Update(v Var) {
	if h.Contains(v) {
		h.Fix(v)
	} else {
		h.Push(v)
	}
}

func (h *minCostHeap) Fix(v Var) {
	heap.Fix(h, h.index[v])
}

func (h *minCostHeap) Init() {
	heap.Init(h)
}

func (h *minCostHeap) Rebuild(vs []Var) {
	h.vars = append(h.vars[:0], vs...)
	h.index = h.index[:0]
	for i := range vs {
		h.extend(vs[i])
		h.index[vs[i]] = i
	}
	heap.Init(h)
}
