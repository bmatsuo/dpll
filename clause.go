package dpll

// Clause is a disjunction of literals.
type Clause struct {
	Mark      byte
	Learnt    bool
	Relocated bool
	Size      int32
	Lit       []Lit
	*ClauseExtra
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

// ClauseExtra are extra fields that may appear in a clause.
type ClauseExtra struct {
	Activity    float64
	Abstraction uint32
}

// Subsumes checks if c subsumes c2 and if it can be used to simplify c2 by
// subsumption resolution.  If Subsumes returns true and p is not LitUndef then
// p can be removed from c2.
func (c *Clause) Subsumes(c2 *Clause) (ok bool, p Lit) {
	if c.Size != c2.Size && c.Abstraction != c2.Abstraction {
		return false, LitUndef
	}
	p = LitUndef

	// subsumtion requires the set of literals to be the same except for the
	// presence of p in c and ~p c2 (where p = ~~p for any literal p).
outer:
	for _, lit1 := range c.Lit {
		for _, lit2 := range c.Lit {
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
