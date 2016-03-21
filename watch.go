// Copyright 2016 Bryan Matsuo
//
// Use of this software is governed by the MIT license.  A copy of the license
// agreement can be found in the LICENSE file distributed with this software.

package dpll

// occLists maintains Lit watcher lists used during propagation of assignment
// to derive models/conflicts.
type occLists struct {
	_occs   [][]watcher
	_dirty  []bool
	occs    [][]watcher
	dirty   []bool
	dirties []Lit
}

// newOccLists initializes and returns a new occLists.
func newOccLists() *occLists {
	return &occLists{}
}

// Push adds w to the list of watchers for p.
func (o *occLists) Push(p Lit, w watcher) {
	o.extend(int(p))
	o.occs[p] = append(o.occs[p], w)
}

func (o *occLists) extend(n int) {
	if n >= len(o.occs) {
		if n >= len(o._occs) {
			for i := len(o._occs) - 1; i < n; i++ {
				o._occs = append(o._occs, nil)
			}
			for i := len(o._dirty) - 1; i < n; i++ {
				o._dirty = append(o._dirty, false)
			}
		}
		o.occs = o._occs[:n+1]
		o.dirty = o._dirty[:n+1]
	}
}

func (o *occLists) get(p Lit) ([]watcher, bool) {
	if int(p) < len(o.occs) {
		ws := o.occs[p]
		return ws, ws == nil
	}
	return nil, false
}

func (o *occLists) RemoveAll(p Lit, free bool) {
	if free {
		o.occs = nil
		return
	}

	ws, ok := o.get(p)
	if !ok {
		return
	}
	for i := range ws {
		ws[i] = watcher{}
	}
	o.occs[p] = ws[:0]
}

// Remove immediately removes w from the watcher list of p.  Instead of calling
// Remove repeatedly the associated clauses can be marked as deleted and their
// literals dirtied using Smudge.  Using Smudge helps amortize the cost of
// removing multiple watchers between calls to Lookup.
func (o *occLists) Remove(p Lit, w watcher) {
	o.extend(int(p))
	occs, _ := o.get(p)
	var j int
	for i := range occs {
		if !occs[i].Equal(&w) {
			occs[j] = occs[i]
			j++
		}
	}
	for i := j; i < len(occs); i++ {
		occs[i] = watcher{}
	}
	o.occs[p] = occs[:j]
}

func (o *occLists) Init(p Lit) {
	// no-op when using a Go map? that might change..
}

// Occurrences returns the watcher list for p.  The watcher list may contain
// deleted clauses.
func (o *occLists) Occurrences(p Lit) []watcher {
	ws, _ := o.get(p)
	return ws
}

// Lookup performs any pending lazy watcher removals for p before returning its
// watcher list.  The returned watcher list will never contain deleted clauses.
func (o *occLists) Lookup(p Lit) []watcher {
	o.extend(int(p))
	if o.dirty[p] {
		o.Clean(p)
	}
	return o.occs[p]
}

// Clear purges all watchers if free is true then large internal allocations ar
// released for garbage collection by the runtime.
func (o *occLists) Clear(free bool) {
	if free {
		o.occs = nil
		o.dirty = nil
		o.dirties = nil
	} else {
		// replacing maps with empty maps is way easier
		for i := range o._occs {
			o._occs[i] = nil
		}
		o.occs = o.occs[:0]
		o.dirty = o.dirty[:0]
		o.dirties = o.dirties[:0]
	}
}

// CleanAll processes pending watcher removals for all literals.
func (o *occLists) CleanAll() {
	for _, p := range o.dirties {
		// dirties may contain duplicates, o.dirty should be checked always.
		o.extend(int(p))
		if o.dirty[p] {
			o.Clean(p)
		}
	}
	o.dirties = o.dirties[:0]
}

// Clean processes pending watcher removals for p.
func (o *occLists) Clean(p Lit) {
	occs, _ := o.get(p)
	if len(occs) == 0 {
		return
	}
	var j int
	for i := range occs {
		if !occs[i].IsDeleted() {
			occs[j] = occs[i]
			j++
		}
	}
	o.occs[p] = occs[:j]
	o.dirty[p] = false
}

// Smudge marks p as dirty
func (o *occLists) Smudge(p Lit) {
	o.extend(int(p))
	if !o.dirty[p] {
		o.dirty[p] = true
		o.dirties = append(o.dirties, p)
	}
}

// watcher is a Clause which is blocked in model search by Lit blocker.
type watcher struct {
	c       *Clause
	blocker Lit
}

// IsDeleted returns true if the watcher's clause is marked with MarkDel.
//
// TODO: It would be nice if this could be a little more extensible.  It forces
// use of MarkDel currently.
func (w *watcher) IsDeleted() bool {
	return w.c.Mark == MarkDel
}

// Equal returns true if w and w2 represent the same clause and is intended to
// compare watchers within the same literal watcher list.
func (w *watcher) Equal(w2 *watcher) bool {
	return w.c == w2.c
}
