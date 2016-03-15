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

// NewLit creates a Lit from v.  If neg is true then the literal is negated.
func NewLit(v Var, pos bool) Lit {
	if pos {
		return Lit(v+v) + 1
	}
	return Lit(v << 1)
}

// NewLitInt creates a Lit from integer x.
func NewLitInt(x int) Lit {
	if x < 0 {
		return NewLit(Var(-x), false)
	}
	return NewLit(Var(x), true)
}

// String returns the string representation of l. For example "1", "-2", or
// "0".
func (l Lit) String() string {
	if l.IsPos() {
		return strconv.Itoa(int(l.Var()))
	}
	return strconv.Itoa(-int(l.Var()))
}

// Var returns the variable in l.  If IsUndef returns is true Var returns an
// undefined variable (its IsUndef method will return true).
func (l Lit) Var() Var {
	return Var(l >> 1)
}

// IsPos returns true iff l is a positive literal.  If IsUndef returns true
// then the return value of IsPos is unspecified.
func (l Lit) IsPos() bool {
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

// Xor returns a literal of l.Var() that is positive iff l.IsPos() XOR b.
func (l Lit) Xor(b bool) Lit {
	return l ^ btol(b)
}

func btol(b bool) Lit {
	if b {
		return 1
	}
	return 0
}
