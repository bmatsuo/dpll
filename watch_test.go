// Copyright 2016 Bryan Matsuo
//
// Use of this software is governed by the MIT license.  A copy of the license
// agreement can be found in the LICENSE file distributed with this software.

package dpll

import (
	"reflect"
	"testing"
)

func TestOccList_Remove(t *testing.T) {
	lits1 := []Lit{
		Literal(1, true),
		Literal(2, false),
	}
	lits2 := []Lit{
		Literal(1, true),
		Literal(3, true),
	}
	c1 := NewClause(lits1, false, false)
	c2 := NewClause(lits2, false, false)
	w1 := watcher{c1, lits1[1]}
	w2 := watcher{c1, lits1[1].Inverse()}
	w3 := watcher{c2, lits2[1]}
	w4 := watcher{c2, lits2[1].Inverse()}

	o := newOccLists()
	o.Push(lits1[0], w1)
	o.Push(lits1[0], w2)
	o.Push(lits2[0], w3)
	o.Push(lits2[0], w4)

	ws := o.Occurrences(lits1[0])
	if !reflect.DeepEqual(ws, []watcher{w1, w2, w3, w4}) {
		t.Errorf("watchers: %v (!= %v)", ws, []watcher{w1, w2, w3, w4})
	}

	// remove w1 and w2
	o.Remove(lits1[0], w1)
	o.Remove(lits1[0], w2)
	ws = o.Occurrences(lits1[0])
	if !reflect.DeepEqual(ws, []watcher{w3, w4}) {
		t.Errorf("watchers: %v (!= %v)", ws, []watcher{w3, w4})
	}
}

func TestOccList_Clean(t *testing.T) {
	lits1 := []Lit{
		Literal(1, true),
		Literal(2, false),
	}
	lits2 := []Lit{
		Literal(1, true),
		Literal(3, true),
	}
	c1 := NewClause(lits1, false, false)
	c2 := NewClause(lits2, false, false)
	w1 := watcher{c1, lits1[1]}
	w2 := watcher{c1, lits1[1].Inverse()}
	w3 := watcher{c2, lits2[1]}
	w4 := watcher{c2, lits2[1].Inverse()}

	o := newOccLists()
	o.Push(lits1[0], w1)
	o.Push(lits1[0], w2)
	o.Push(lits2[0], w3)
	o.Push(lits2[0], w4)

	ws := o.Lookup(lits1[0])
	if !reflect.DeepEqual(ws, []watcher{w1, w2, w3, w4}) {
		t.Errorf("watchers: %v (!= %v)", ws, []watcher{w1, w2, w3, w4})
	}

	// smudge without deleting anything -- just wasting cpu cycles
	o.Smudge(lits1[0])
	ws = o.Lookup(lits1[0])
	if !reflect.DeepEqual(ws, []watcher{w1, w2, w3, w4}) {
		t.Errorf("watchers: %v (!= %v)", ws, []watcher{w1, w2, w3, w4})
	}

	// mark and smudge will clean w1 and w2 at the next lookup
	c1.Mark = MarkDel
	o.Smudge(lits1[0])
	ws = o.Lookup(lits1[0])
	if !reflect.DeepEqual(ws, []watcher{w3, w4}) {
		t.Errorf("watchers: %v (!= %v)", ws, []watcher{w3, w4})
	}
}

func TestWatcher_IsDeleted(t *testing.T) {
	lits := []Lit{
		Literal(1, false),
		Literal(1, true),
	}
	c1 := NewClause(lits, false, false)
	c2 := NewClause(lits, false, false)
	c2.Mark = MarkDel

	w1 := watcher{c1, lits[0]}
	w2 := watcher{c2, lits[0]}

	if w1.IsDeleted() {
		t.Errorf("deleted: %v", w1.c.Mark)
	}
	if !w2.IsDeleted() {
		t.Errorf("not deleted: %v", w2.c.Mark)
	}
}

func TestWatcher_Equal(t *testing.T) {
	l1 := []Lit{
		Literal(1, false),
		Literal(1, true),
	}
	l2 := []Lit{
		Literal(1, false),
		Literal(2, true),
	}
	c1 := NewClause(l1, false, false)
	c2 := NewClause(l2, false, false)

	w1p0 := watcher{c1, l1[0]}
	w1n0 := watcher{c1, l1[0].Inverse()}
	w2p0 := watcher{c2, l2[0]}
	w2p1 := watcher{c2, l2[1]}

	if !w1p0.Equal(&w1n0) {
		t.Errorf("%v != %v (%v == %v)", w1p0, w1n0, w1p0, w1n0)
	}

	if !w2p0.Equal(&w2p1) {
		t.Errorf("%v != %v (%v == %v)", w2p0, w2p1, w2p0, w2p1)
	}

	if w1p0.Equal(&w2p1) {
		t.Errorf("%v != %v (%v == %v)", w1p0, w2p1, w1p0, w2p1)
	}
}
