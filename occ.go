// Copyright 2016 Bryan Matsuo
//
// Use of this software is governed by the MIT license.  A copy of the license
// agreement can be found in the LICENSE file distributed with this software.

package dpll

// occLists maintains occurrance lists with lazy deletion
type occLists struct {
	occs    map[Lit][]watcher
	dirty   map[Lit]bool
	dirties []Lit
}

func newOccLists() *occLists {
	return &occLists{
		occs:  make(map[Lit][]watcher),
		dirty: make(map[Lit]bool),
	}
}

func (o *occLists) Push(p Lit, w watcher) {
	o.occs[p] = append(o.occs[p], w)
}

func (o *occLists) Remove(p Lit, w watcher) {
	occs := o.occs[p]
	if len(occs) == 0 {
		return
	}
	var j int
	for i := range occs {
		if !occs[i].Equal(&w) {
			occs[j] = occs[i]
			j++
		}
	}
	o.occs[p] = occs[:j]
}

func (o *occLists) Init(p Lit) {
	// no-op when using a Go map? that might change..
}

func (o *occLists) Occurances(p Lit) []watcher {
	return o.occs[p]
}

func (o *occLists) Lookup(p Lit) []watcher {
	if o.dirty[p] {
		o.Clean(p)
	}
	return o.occs[p]
}

func (o *occLists) Clear(free bool) {
	if free {
		o.occs = map[Lit][]watcher{}
		o.dirty = map[Lit]bool{}
		o.dirties = nil
	} else {
		// replacing maps with empty maps is way easier
		o.occs = make(map[Lit][]watcher, len(o.occs))
		o.dirty = make(map[Lit]bool, len(o.dirty))
		o.dirties = o.dirties[:0]
	}
}

func (o *occLists) CleanAll() {
	for _, p := range o.dirties {
		// dirties may contain duplicates, o.dirty should be checked always.
		if o.dirty[p] {
			o.Clean(p)
		}
	}
	o.dirties = o.dirties[:0]
}

func (o *occLists) Clean(p Lit) {
	occs := o.occs[p]
	if len(occs) == 0 {
		panic("no occurances")
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
	if !o.dirty[p] {
		o.dirty[p] = true
		o.dirties = append(o.dirties, p)
	}
}

type watcher struct {
	c       *Clause
	blocker Lit
}

func (w *watcher) IsDeleted() bool {
	return w.c.Mark == 1
}

func (w *watcher) Equal(w2 *watcher) bool {
	return w.c == w2.c
}
