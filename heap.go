// Copyright 2016 Bryan Matsuo
//
// Use of this software is governed by the MIT license.  A copy of the license
// agreement can be found in the LICENSE file distributed with this software.

package dpll

import "container/heap"

// activityQueue is a heap with special access methods.
type activityQueue maxActiveHeap

func newActivityQueue(activity *[]float64) *activityQueue {
	return (*activityQueue)(newMaxActiveHeap(activity))
}

func (q *activityQueue) Len() int {
	return (*maxActiveHeap)(q).Len()
}

// this is super confusing. the literature needs to be consulted to figure out
// if max activity or min activity is really desired.
func (q *activityQueue) RemoveMin() Var {
	return heap.Pop((*maxActiveHeap)(q)).(Var)
}

func (q *activityQueue) Contains(v Var) bool {
	return (*maxActiveHeap)(q).Contains(v)
}

func (q *activityQueue) Decrease(v Var) {
	heap.Fix((*maxActiveHeap)(q), q.index[v])
}

func (q *activityQueue) Rebuild(vs []Var) {
	(*maxActiveHeap)(q).Rebuild(vs)
}

func (q *activityQueue) Push(v Var) {
	(*maxActiveHeap)(q).Push(v)
}

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
