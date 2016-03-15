package dpll

// DPLL is a DPLL satisfiability (SAT) solver.
type DPLL struct {
}

// New initializes and returns a new SAT solver.
func New() *DPLL {
	d := &DPLL{}
	return d
}
