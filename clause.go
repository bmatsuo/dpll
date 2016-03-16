package dpll

// ClauseHeader contains Clause metadata that can be inherited from other
// clauses.
type ClauseHeader struct {
	Mark      byte
	Learnt    bool
	Relocated bool
	*ClauseExtra
}

// Clause is a disjunction of literals.
type Clause struct {
	ClauseHeader
	Lit []Lit
}

// NewClause creates a new clause from the given literals.  After calling
// NewClause the slice ps must not be modified in the future.
func NewClause(ps []Lit, extra bool, learnt bool) *Clause {
	c := &Clause{}
	c.Lit = ps
	c.Learnt = learnt
	if extra {
		c.ClauseExtra = &ClauseExtra{}
		if learnt {
			c.Activity = 0
		} else {
			c.CalcAbstraction()
		}
	}
	return c
}

// NewClauseFrom creates a new clause with an inherited ClauseHeader.  The
// extra argument overrides any from ClauseExtra metadata.
func NewClauseFrom(from *Clause, extra bool) *Clause {
	c := &Clause{}
	c.ClauseHeader = from.ClauseHeader
	if !extra {
		c.ClauseExtra = nil
	} else if from.Learnt {
		c.ClauseExtra = &ClauseExtra{Activity: from.Activity}
	} else {
		c.ClauseExtra = &ClauseExtra{Abstraction: from.Abstraction}
	}
	c.Lit = make([]Lit, len(from.Lit))
	copy(c.Lit, from.Lit)
	return c
}

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

// CalcAbstraction computes and stores an abstraction of variables in c like a
// checksum.
func (c *Clause) CalcAbstraction() {
	var abs uint32
	for _, lit := range c.Lit {
		abs |= 1 << (uint32(lit.Var()) & 31)
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
	if c.Len() != c2.Len() && c.Abstraction != c2.Abstraction {
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
			if p == LitUndef && lit1 == lit2.Inverse() {
				p = lit1
				continue outer
			}
		}
		return false, LitUndef
	}

	return true, p
}

// Strengthen removes p from the list of literals in c.
func (c *Clause) Strengthen(p Lit) {
	for i, lit := range c.Lit {
		if lit == p {
			copy(c.Lit[i:], c.Lit[i+1:])
			c.Lit = c.Lit[:len(c.Lit)-1]
			c.CalcAbstraction()
			return
		}
	}
}
