// Copyright 2016 Bryan Matsuo
//
// Use of this software is governed by the MIT license.  A copy of the license
// agreement can be found in the LICENSE file distributed with this software.

package dpll

type clauseQueue struct {
	q     []*Clause
	start int
	end   int
}

func newClauseQueue(cap int) *clauseQueue {
	if cap < 1 {
		cap = 1
	}
	return &clauseQueue{
		q: make([]*Clause, cap),
	}
}

func (q *clauseQueue) clear() {
	for i := range q.q {
		q.q[i] = nil
	}
	q.q = q.q[:1]
	q.start = 0
	q.end = 0
}

func (q *clauseQueue) Len() int {
	if q.end < q.start {
		return len(q.q) - q.end + q.start
	}
	return q.end - q.start
}

func (q *clauseQueue) Clause(i int) *Clause {
	if i < 0 || i >= q.Len() {
		panic("index out of range")
	}
	return q.q[(q.start+i)%len(q.q)]
}

func (q *clauseQueue) Front() *Clause {
	if q.start == q.end {
		panic("index out of range")
	}
	return q.q[q.start]
}

func (q *clauseQueue) Pop() *Clause {
	c := q.Front()
	q.q[q.start] = nil
	q.start = (q.start + 1) % len(q.q)
	return c
}

func (q *clauseQueue) Insert(c *Clause) {
	if c == nil {
		panic("nil clause")
	}
	// q.end is always unused
	q.q[q.end] = c
	q.end = (q.end + 1) % len(q.q)

	if q.end == q.start {
		old := q.q
		q.q = make([]*Clause, (len(old)*3+1)>>1)
		n := copy(q.q, old[q.start:])
		n2 := copy(q.q[n:], old[:q.end])
		q.start = 0
		q.end = n + n2
	}
}
