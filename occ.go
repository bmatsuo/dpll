// Copyright 2016 Bryan Matsuo
//
// Use of this software is governed by the MIT license.  A copy of the license
// agreement can be found in the LICENSE file distributed with this software.

package dpll

// clauseOccLists maintains lists of clauses containing each literal.
type clauseOccLists struct {
	_occs   [][]*Clause
	_dirty  []bool
	occs    [][]*Clause
	dirty   []bool
	dirties []Var
}

// newOccLists initializes and returns a new clauseOccLists.
func newClauseOccLists() *clauseOccLists {
	return &clauseOccLists{}
}

// Push adds c to the list of clauses for p.
func (o *clauseOccLists) Push(p Var, c *Clause) {
	o.extend(int(p))
	o.occs[p] = append(o.occs[p], c)
}

func (o *clauseOccLists) extend(n int) {
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

func (o *clauseOccLists) RemoveAll(v Var, free bool) {
	o.extend(int(v))

	if free {
		o.occs[v] = nil
		return
	}

	occs := o.occs[v]
	for i := range occs {
		occs[i] = nil
	}
	o.occs[v] = occs[:0]
}

// Remove immediately removes c from clause list of p.  Instead of calling
// Remove repeatedly the associated clauses can be marked as deleted and their
// literals dirtied using Smudge.  Using Smudge helps amortize the cost of
// removing multiple clauses between calls to Lookup.
func (o *clauseOccLists) Remove(p Var, c *Clause) {
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

func (o *clauseOccLists) Init(p Var) {
	// no-op when using a Go map? that might change..
}

// Occurrences returns the clause list for p.  The lause list may contain
// deleted clauses.
func (o *clauseOccLists) Occurrences(p Var) []*Clause {
	return o.occs[p]
}

// Lookup performs any pending lazy clause removals for p before returning its
// clause list.  The returned clause list will never contain deleted clauses.
func (o *clauseOccLists) Lookup(p Var) []*Clause {
	if o.dirty[p] {
		o.Clean(p)
	}
	return o.occs[p]
}

// Clear purges all clauses if free is true then large internal allocations ar
// released for garbage collection by the runtime.
func (o *clauseOccLists) Clear(free bool) {
	if free {
		o._occs = nil
		o._dirty = nil
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
func (o *clauseOccLists) Clean(p Var) {
	occs := o.occs[p]
	if len(occs) == 0 {
		return
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
func (o *clauseOccLists) Smudge(p Var) {
	if !o.dirty[p] {
		o.dirty[p] = true
		o.dirties = append(o.dirties, p)
	}
}
