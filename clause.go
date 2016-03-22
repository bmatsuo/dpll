// Copyright 2016 Bryan Matsuo
//
// Use of this software is governed by the MIT license.  A copy of the license
// agreement can be found in the LICENSE file distributed with this software.

//go:generate bash build_clauselocal.sh

package dpll

// Clause is a disjunction of literals.
type Clause struct {
	ClauseHeader
	Lit []Lit
}

// clauseExtra used during clause construction when needed and forces data
// locality to limit indirection overhead (by avoiding most cache misses).
type clauseExtra struct {
	Clause
	ClauseExtra
}

// NewClause creates a new clause containing a copy of ps.
func NewClause(ps []Lit, extra bool, learnt bool) *Clause {
	c := newClause(ps, extra)
	c.Learnt = learnt
	if extra && !learnt {
		c.CalcAbstraction()
	}
	return c
}

// NewClauseFrom creates a new clause with an inherited ClauseHeader.  The
// extra argument overrides any from ClauseExtra metadata.
func NewClauseFrom(from *Clause, extra bool) *Clause {
	c := newClause(from.Lit, extra)
	if extra {
		ce := c.ClauseExtra
		*ce = *from.ClauseExtra
		c.ClauseHeader = from.ClauseHeader
		c.ClauseExtra = ce
	} else {
		c.ClauseHeader = from.ClauseHeader
		c.ClauseHeader.ClauseExtra = nil
	}
	return c
}

// newClause constructs a new clause using a copy of ps.  if extra is true the
// returned clause will have a non-nil ClauseExtra field.  newClause attempts
// to ensure that all allocated memory is in a continuous structure to maximize
// data locality and cache hits.
func newClause(ps []Lit, extra bool) *Clause {
	if extra {
		if len(ps) < len(mkClauseExtraLocal) {
			// allocate data as a continuous chunk to maximize cache hits
			return mkClauseExtraLocal[len(ps)](ps)
		}
		_ps := make([]Lit, len(ps))
		copy(_ps, ps)
		ce := &clauseExtra{
			Clause: Clause{Lit: _ps},
		}
		ce.Clause.ClauseExtra = &ce.ClauseExtra
		return &ce.Clause
	}
	if len(ps) < len(mkClauseLocal) {
		// allocate data as a continuous chunk to maximize cache hits
		return mkClauseLocal[len(ps)](ps)
	}
	_ps := make([]Lit, len(ps))
	copy(_ps, ps)
	return &Clause{Lit: _ps}
}

/*
// It is not clear yet if this abstraction can improve performance.

type clauseAllocator struct {
	ForceExtra bool
}

func (ca *clauseAllocator) New(ps []Lit, learnt bool) *Clause {
	extra := learnt || ca.ForceExtra
	return NewClause(ps, extra, learnt)
}

func (ca *clauseAllocator) From(c *Clause) *Clause {
	extra := c.Learnt || ca.ForceExtra
	return NewClauseFrom(c, extra)
}
*/

// CalcAbstraction computes and stores an abstraction of variables in c like a
// checksum.
func (c *Clause) CalcAbstraction() {
	var abs uint32
	for _, lit := range c.Lit {
		abs |= 1 << ((lit.Var() - 1) & 31)
	}
	c.Abstraction = abs
}

// Len returns the number of literals in c.
func (c *Clause) Len() int {
	return len(c.Lit)
}

// ClauseExtra are extra fields that may appear in a clause.
type ClauseExtra struct {
	Activity    float64
	Abstraction uint32
}

// Subsumes checks if c subsumes c2 and if it can be used to simplify c2 by
// subsumption resolution.  If Subsumes returns true and p is not LitUndef then
// p can be removed from c2.
func (c *Clause) Subsumes(c2 *Clause) (ok bool, p Lit) {
	if c.Len() != c2.Len() && c.Abstraction&^c2.Abstraction != 0 {
		return false, LitUndef
	}
	p = LitUndef

	// subsumtion requires the set of literals to be the same except for the
	// presence of p in c and ~p c2 (where p = ~~p for any literal p).
outer:
	for i := range c.Lit {
		for j := range c2.Lit {
			if !c.Lit[i].SharesVar(c2.Lit[j]) {
				continue
			}
			if c.Lit[i] == c2.Lit[j] {
				continue outer
			}
			if !p.IsUndef() {
				break
			}
			p = c.Lit[i]
			continue outer
		}
		return false, LitUndef
	}

	return true, p
}

// Strengthen removes p from the list of literals in c.
func (c *Clause) Strengthen(p Lit) bool {
	for i, lit := range c.Lit {
		if lit == p {
			copy(c.Lit[i:], c.Lit[i+1:])
			c.Lit = c.Lit[:len(c.Lit)-1]

			c.CalcAbstraction()
			return true
		}
	}
	return false
}

// ClauseHeader contains Clause metadata that can be inherited from other
// clauses.
type ClauseHeader struct {
	Mark      Mark
	Learnt    bool
	Relocated bool // this seems unnecessary
	*ClauseExtra
}

// Mark is a small scratch space that can be used to flag clauses for
// application dependent purposes.  The DPLL type in this package sets the Mark
// to signal that a clause should be deleted.
type Mark uint8

// HasAny returns true if and only if m contains any of the bits marked in mm.
func (m Mark) HasAny(mm Mark) bool {
	return m&mm != 0
}

// HasAll returns true if and only if m contains all of the bits marked in mm.
func (m Mark) HasAll(mm Mark) bool {
	return m&mm != mm
}

// Marks that are used by the package's DPLL solver.  Use of any Mark constants
// in not required in a custom solver.
const (
	MarkDel Mark = 1 // The marked clause should be deleted
)
