// Copyright 2016 Bryan Matsuo
//
// Use of this software is governed by the MIT license.  A copy of the license
// agreement can be found in the LICENSE file distributed with this software.

package dpll

// Clause is a disjunction of literals.
type Clause struct {
	ClauseHeader
	Lit []Lit
}

type clauseExtra struct {
	Clause
	ClauseExtra
}

const smallLimit = 16

type clauseSmall struct {
	Clause
	Lit [smallLimit]Lit
}

type clauseExtraSmall struct {
	Clause
	ClauseExtra // layout after Lit to copy minisat
	Lit         [smallLimit]Lit
}

// NewClause creates a new clause from the given literals.  After calling
// NewClause the slice ps must not be modified in the future.
func NewClause(ps []Lit, extra bool, learnt bool) *Clause {
	if extra {
		if len(ps) < smallLimit {
			cs := &clauseExtraSmall{
				Clause: Clause{
					ClauseHeader: ClauseHeader{
						Learnt: learnt,
					},
				},
			}
			copy(cs.Lit[:], ps)
			cs.Clause.Lit = cs.Lit[:len(ps)]
			cs.Clause.ClauseExtra = &cs.ClauseExtra
			if !learnt {
				cs.Clause.CalcAbstraction()
			}
			return &cs.Clause
		}

		c := &clauseExtra{}
		c.Clause.Learnt = learnt
		c.Lit = make([]Lit, len(ps))
		copy(c.Lit, ps)
		c.Clause.ClauseExtra = &c.ClauseExtra
		if !learnt {
			c.CalcAbstraction()
		}
		return &c.Clause
	}
	if len(ps) < smallLimit {
		cs := &clauseSmall{
			Clause: Clause{
				ClauseHeader: ClauseHeader{
					Learnt: learnt,
				},
			},
		}
		copy(cs.Lit[:], ps)
		cs.Clause.Lit = cs.Lit[:len(ps)]
		return &cs.Clause
	}
	_ps := make([]Lit, len(ps))
	copy(_ps, ps)
	return &Clause{
		ClauseHeader: ClauseHeader{
			Learnt: learnt,
		},
		Lit: _ps,
	}
}

// NewClauseFrom creates a new clause with an inherited ClauseHeader.  The
// extra argument overrides any from ClauseExtra metadata.
func NewClauseFrom(from *Clause, extra bool) *Clause {
	if len(from.Lit) < smallLimit {
		if !extra {
			cs := &clauseSmall{
				Clause: Clause{
					ClauseHeader: from.ClauseHeader,
				},
			}
			cs.Clause.ClauseExtra = nil
			copy(cs.Lit[:], from.Lit)
			cs.Clause.Lit = cs.Lit[:len(from.Lit)]
			return &cs.Clause
		}

		ce := &clauseExtraSmall{
			Clause: Clause{
				ClauseHeader: from.ClauseHeader,
			},
			ClauseExtra: *from.ClauseExtra,
		}
		copy(ce.Lit[:], from.Lit)
		ce.Clause.Lit = ce.Lit[:len(from.Lit)]
		ce.Clause.ClauseHeader.ClauseExtra = &ce.ClauseExtra
		return &ce.Clause
	}

	ps := make([]Lit, len(from.Lit))
	copy(ps, from.Lit)
	if !extra {
		c := &Clause{
			ClauseHeader: from.ClauseHeader,
			Lit:          ps,
		}
		c.ClauseExtra = nil
		return c
	}

	ce := &clauseExtra{
		Clause: Clause{
			ClauseHeader: from.ClauseHeader,
			Lit:          ps,
		},
		ClauseExtra: *from.ClauseExtra,
	}
	ce.Clause.ClauseHeader.ClauseExtra = &ce.ClauseExtra
	return &ce.Clause
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
	for _, lit1 := range c.Lit {
		for _, lit2 := range c2.Lit {
			if lit1 == lit2 {
				continue outer
			}
			if p.IsUndef() && lit1 == lit2.Inverse() {
				p = lit1
				continue outer
			}
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
