// Copyright 2016 Bryan Matsuo
//
// Use of this software is governed by the MIT license.  A copy of the license
// agreement can be found in the LICENSE file distributed with this software.

package dpll

// clauseOccLists maintains lists of clauses containing each literal.
type clauseOccLists struct {
	occs    map[Lit][]*Clause
	dirty   map[Lit]bool
	dirties []Lit
}

// newOccLists initializes and returns a new clauseOccLists.
func newClauseOccLists() *clauseOccLists {
	return &clauseOccLists{
		occs:  make(map[Lit][]*Clause),
		dirty: make(map[Lit]bool),
	}
}

// Push adds c to the list of clauses for p.
func (o *clauseOccLists) Push(p Lit, c *Clause) {
	o.occs[p] = append(o.occs[p], c)
}

// Remove immediately removes c from clause list of p.  Instead of calling
// Remove repeatedly the associated clauses can be marked as deleted and their
// literals dirtied using Smudge.  Using Smudge helps amortize the cost of
// removing multiple clauses between calls to Lookup.
func (o *clauseOccLists) Remove(p Lit, c *Clause) {
	occs := o.occs[p]
	if len(occs) == 0 {
		return
	}
	var j int
	for i := range occs {
		if occs[i] != c {
			occs[j] = occs[i]
			j++
		}
	}
	o.occs[p] = occs[:j]
}

func (o *clauseOccLists) Init(p Lit) {
	// no-op when using a Go map? that might change..
}

// Occurrences returns the clause list for p.  The lause list may contain
// deleted clauses.
func (o *clauseOccLists) Occurrences(p Lit) []*Clause {
	return o.occs[p]
}

// Lookup performs any pending lazy clause removals for p before returning its
// clause list.  The returned clause list will never contain deleted clauses.
func (o *clauseOccLists) Lookup(p Lit) []*Clause {
	if o.dirty[p] {
		o.Clean(p)
	}
	return o.occs[p]
}

// Clear purges all clauses if free is true then large internal allocations ar
// released for garbage collection by the runtime.
func (o *clauseOccLists) Clear(free bool) {
	if free {
		o.occs = map[Lit][]*Clause{}
		o.dirty = map[Lit]bool{}
		o.dirties = nil
	} else {
		// replacing maps with empty maps is way easier
		o.occs = make(map[Lit][]*Clause, len(o.occs))
		o.dirty = make(map[Lit]bool, len(o.dirty))
		o.dirties = o.dirties[:0]
	}
}

// CleanAll processes pending clause removals for all literals.
func (o *clauseOccLists) CleanAll() {
	for _, p := range o.dirties {
		// dirties may contain duplicates, o.dirty should be checked always.
		if o.dirty[p] {
			o.Clean(p)
		}
	}
	o.dirties = o.dirties[:0]
}

// Clean processes pending clause removals for p.
func (o *clauseOccLists) Clean(p Lit) {
	occs := o.occs[p]
	if len(occs) == 0 {
		panic("no occurrences")
	}
	var j int
	for i := range occs {
		if !occs[i].Mark.HasAny(MarkDel) {
			occs[j] = occs[i]
			j++
		}
	}
	o.occs[p] = occs[:j]
	o.dirty[p] = false
}

// Smudge marks p as dirty
func (o *clauseOccLists) Smudge(p Lit) {
	if !o.dirty[p] {
		o.dirty[p] = true
		o.dirties = append(o.dirties, p)
	}
}
