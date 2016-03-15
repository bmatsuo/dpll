package dpll

// DPLL is a DPLL satisfiability (SAT) solver.
type DPLL struct {
	Verbosity     int
	VarDecay      float64
	ClauseDecay   float64
	RandomVarFreq float64
	RandomSeed    int64
	LubyRestart   bool
	CCMin         CCMinMode
	PhaseSaving   PhaseSavingLevel
	RandPol       bool // Random polarities for branching heuristics.
	RandInitAct   bool // Initialize variable activities with a small random value.
	MinLearnt     int  // Minimum number to set learnt limit to.
	//GarbageFrac float64 // fraction of wasted memory allowed before garbage collection (???)

	RestartFirst   int     // The initial restart limit.
	RestartIncr    float64 // Factor by which limit increases with each restart.
	LearntFraction float64 // Initial limit for learnt clauses as fraction of original clauses.
	LearntIncr     float64 // Factor by which limit for learnt clauses increase with each restart.

	LearntAdjustConfl int
	LearntAdjustInc   float64

	nsolves        uint64
	nstarts        uint64
	ndecisions     uint64
	nrandDecisions uint64
	npropogations  uint64
	nconflicts     uint64
	ndecVars       uint64
	nclauses       uint64
	nlearnt        uint64
	nclauseLit     uint64
	nlearntLit     uint64
	nmaxLit        uint64
	ntotLit        uint64

	clauses     []*Clause
	learnt      []*Clause
	trail       []Lit // Assignment stack
	trailLim    []int // Seprarating indices for decision levels in trail
	assumptions []Lit // Set of assumptions provided by the user

	activity  []float64
	assigns   []LBool
	polarity  []uint8
	upolarity []LBool
	decision  []uint8
	vardata   []varData
	// watches OccList // 'watches[lit]' is a list of constraints watching 'lit' (will go there if literal becomes true).

	// orderHeap heap.Interface

	ok          bool    // If false the constraints are already unsatisfiable. No part of the solver state may be used.
	claIncr     float64 // Amount to bump next clause with
	varIncr     float64 // Amount to bump next variable with
	qhead       int     // Head of queue as index into trail.
	nsimpAssign int     // Number of top level assignments since last call to Simplify
	nsimpProps  int64   // Number of propagations that must be made before the next call to Simplify
	progress    float64 // Estimate set by search
	removeSat   bool    // Indicates whether a possibly inefficient scan for satisfied clauses should be done in Simplify
	varNext     Var     // Next variable to be created

	releasedVars []Var
	freeVars     []Var

	// temporary buffers
	seen []bool
	//analyzeStack []shrinkStackElem
	analyzeToClear []Lit
	addTmp         []Lit

	maxLearnt         float64
	learntAdjustConfl float64
	learntAdjustCnt   int

	// resource constraints
	conflictBudget    int64
	propagationBudget int64
	asyncInterrupt    bool
}

type varData struct {
	Reason *Clause
	Level  int
}

// New initializes and returns a new SAT solver.
func New() *DPLL {
	d := &DPLL{}
	return d
}

// NewVar adds a new variable. The parameters specify variable mode.
func (d *DPLL) NewVar(upol LBool, dvar bool) Var {
	return 0
}

// SetPolarity declares which polarity the decision heuristic should use for a
// variable. Requires mode 'polarity_user'.
func (d *DPLL) SetPolarity(v Var, upol LBool) {
}

// SetDecision declares if a variable should be eligible for selection in the
// decision heuristic.
func (d *DPLL) SetDecision(v Var, ok bool) {
}

// Value returns the current value of v.
func (d *DPLL) Value(v Var) LBool {
	return LUndef
}

// ValueLit returns the current value of lit.
func (d *DPLL) ValueLit(lit Lit) LBool {
	return LUndef
}

// ValueModel returns the value of v in the last model.  The last call to Solve
// must have returned true.
func (d *DPLL) ValueModel(v Var) LBool {
	return LUndef
}

// ValueLitModel returns the value of lit in the last model.  The last call to
// Solve must have returned true.
func (d *DPLL) ValueLitModel(lit Lit) LBool {
	return LUndef
}

// NumAssign returns the number of assigned literals.
func (d *DPLL) NumAssign() int {
	return 0
}

// NumClause returns the number of original clauses.
func (d *DPLL) NumClause() int {
	return 0
}

// NumLearn returns the number of learnt clauses.
func (d *DPLL) NumLearn() int {
	return 0
}

// NumVar returns the number of variables.
func (d *DPLL) NumVar() int {
	return 0
}

// NumVarFree returns the number of free variables.
func (d *DPLL) NumVarFree() int {
	return 0
}

// AddClause adds a CNF clause containing the given literals.
func (d *DPLL) AddClause(lit ...Lit) bool {
	return false
}

// Simplify removes already satisfied clauses.
func (d *DPLL) Simplify() bool {
	return false
}

// Solve searches for a model that respects the given assupmtions.
func (d *DPLL) Solve(assump ...Lit) bool {
	return false
}

// SolveLimited behaves like Solve but respects resource contraints (?).
func (d *DPLL) SolveLimited(assump ...Lit) LBool {
	return LUndef
}

// Okay returns false if the solver is in a conflicting state.
func (d *DPLL) Okay() bool {
	return false
}

// Implies returns true if the given assumptions imply the given assignments.
func (d *DPLL) Implies(assumps []Lit, assign []Lit) bool {
	return false
}

// Model returns assignments found in the last call to Solve.  If Solve could
// not find a model (perhaps under assupmtions) Model returns nil.
func (d *DPLL) Model() []LBool {
	return nil
}

// Conflict returns the final, non-empty clause expressed in assumptions if
// Solve could not find a model.  If the last call to Solve found a model then
// Conflict returns nil.
func (d *DPLL) Conflict() []Lit {
	return nil
}

// CCMinMode controls conflict clause minimization.
type CCMinMode int

// Available conflict clause minimization modes.
const (
	CCMinNone CCMinMode = iota
	CCMinBasic
	CCMinDeep
)

// PhaseSavingLevel controls the amount of phase saving.
type PhaseSavingLevel int

// Avalaible levels of phase saving.
const (
	PhaseSavingNone PhaseSavingLevel = iota
	PhaseSavingLimited
	PhaseSavingFull
)
