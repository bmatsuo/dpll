// Copyright 2016 Bryan Matsuo
//
// Use of this software is governed by the MIT license.  A copy of the license
// agreement can be found in the LICENSE file distributed with this software.

package dpll

import "strconv"

// Helpful constants
const (
	VarUndef = Var(0)
	LitUndef = Lit(0)

	// unlike miniSAT a lit_Error analogue is unnecessary because go supports
	// multiple return types.
)

// Var is a propositional variable.  Variables begin at 1.
type Var uint

// IsUndef returns true if v is undefined.
func (v Var) IsUndef() bool {
	return v == 0
}

// Lit is a propositional literal that encodes possible negation with a
// variable.
type Lit uint

// Literal creates a Lit from v.  If neg is true then the literal is negated.
func Literal(v Var, neg bool) Lit {
	if neg {
		return Lit(v+v) + 1
	}
	return Lit(v << 1)
}

// LiteralInt creates a Lit from integer x.
func LiteralInt(x int) Lit {
	if x < 0 {
		return Literal(Var(-x), true)
	}
	return Literal(Var(x), false)
}

// String returns the string representation of l. For example "1", "-2", or
// "0".
func (l Lit) String() string {
	if l.IsNeg() {
		return strconv.Itoa(-int(l.Var()))
	}
	return strconv.Itoa(int(l.Var()))
}

// Var returns the variable in l.  If IsUndef returns is true Var returns an
// undefined variable (its IsUndef method will return true).
func (l Lit) Var() Var {
	return Var(l >> 1)
}

// IsNeg returns true iff l is a negated literal.  If IsUndef returns true
// then the return value of IsNeg is unspecified.
func (l Lit) IsNeg() bool {
	return l&1 == 1
}

// IsUndef returns true if l is undefined.
func (l Lit) IsUndef() bool {
	return l.Var().IsUndef()
}

// Inverse returns a negated literal.  If lit is negated the result is a double
// negated (positive) literal.
func (l Lit) Inverse() Lit {
	return l ^ 1
}

// Xor returns a literal of l.Var() that is negated iff l.IsNeg() XOR b.
func (l Lit) Xor(b bool) Lit {
	return l ^ btol(b)
}

func btol(b bool) Lit {
	if b {
		return 1
	}
	return 0
}

type litSlice []Lit

func (s litSlice) Len() int           { return len(s) }
func (s litSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s litSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type litSet map[Lit]struct{}

func newLitSet() litSet {
	return make(litSet)
}

func (s litSet) insert(p Lit) {
	s[p] = struct{}{}
}

func (s litSet) slice() []Lit {
	ps := make([]Lit, 0, len(s))
	for p := range s {
		ps = append(ps, p)
	}
	return ps
}
