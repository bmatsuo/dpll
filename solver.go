// Copyright 2016 Bryan Matsuo
//
// Use of this software is governed by the MIT license.  A copy of the license
// agreement can be found in the LICENSE file distributed with this software.

package dpll

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"sort"
	"time"
)

// Opt declares options for a DPLL solver.
type Opt struct {
	Verbosity     int              // Log verbosity
	VarDecay      float64          // Activity decay factor
	ClauseDecay   float64          // Activity decay factor
	RandVarFreq   float64          // Frequency with which the decision heuristic tries to choose a random variable
	RandSeed      int64            // Control pseudorandom number sequence
	NoLubyRestart bool             // Do not use the Luby restart sequence
	CCMin         CCMinMode        // Control conflict clause minimization
	PhaseSaving   PhaseSavingLevel // Control phase saving
	RandPol       bool             // Random polarities for branching heuristics.
	RandInitAct   bool             // Initialize variable activities with a small random value.
	MinLearnt     int              // Minimum number to set learnt limit to.
	GarbageFrac   float64          // fraction of wasted memory allowed before garbage collection (???)

	RestartFirst   int     // The initial restart limit.
	RestartIncr    float64 // Factor by which limit increases with each restart.
	LearntFraction float64 // Initial limit for learnt clauses as fraction of original clauses.
	LearntIncr     float64 // Factor by which limit for learnt clauses increase with each restart.

	LearntAdjustConfl int
	LearntAdjustIncr  float64
}

var optDefault = &Opt{
	VarDecay:     0.95,
	ClauseDecay:  0.999,
	RandSeed:     0x1234C0DE,
	CCMin:        CCMinDeep,
	PhaseSaving:  PhaseSavingFull,
	RestartFirst: 100,
	RestartIncr:  2,
	GarbageFrac:  0.20,

	LearntFraction: 1.0 / 3.0,
	LearntIncr:     1.1,

	LearntAdjustConfl: 100,
	LearntAdjustIncr:  1.5,
}

func mergeOptDefault(o *Opt) *Opt {
	return mergeOpt(optDefault, o)
}

func mergeOpt(o1, o2 *Opt) *Opt {
	o := &Opt{}
	*o = *o1
	if o2 == nil {
		return o
	}
	if o2.Verbosity != 0 {
		o.Verbosity = o2.Verbosity
	}
	if o2.VarDecay != 0 {
		o.VarDecay = o2.VarDecay
	}
	if o2.ClauseDecay != 0 {
		o.ClauseDecay = o2.ClauseDecay
	}

	if o2.RandVarFreq != 0 {
		o.RandVarFreq = o2.RandVarFreq
	}
	if o2.RandSeed != 0 {
		o.RandSeed = o2.RandSeed
	}

	if o2.NoLubyRestart {
		o.NoLubyRestart = o2.NoLubyRestart
	}

	if o2.CCMin != 0 {
		o.CCMin = o2.CCMin
	}
	if o2.PhaseSaving != 0 {
		o.PhaseSaving = o2.PhaseSaving
	}

	if o2.RandPol {
		o.RandPol = o2.RandPol
	}
	if o2.RandInitAct {
		o.RandInitAct = o2.RandInitAct
	}
	if o2.MinLearnt != 0 {
		o.MinLearnt = o2.MinLearnt
	}

	if o2.RestartFirst != 0 {
		o.RestartFirst = o2.RestartFirst
	}
	if o2.RestartIncr != 0 {
		o.RestartIncr = o2.RestartIncr
	}

	if o2.LearntFraction != 0 {
		o.LearntFraction = o2.LearntFraction
	}
	if o2.LearntIncr != 0 {
		o.LearntIncr = o2.LearntIncr
	}

	if o2.LearntAdjustConfl != 0 {
		o.LearntAdjustConfl = o2.LearntAdjustConfl
	}
	if o2.LearntAdjustIncr != 0 {
		o.LearntAdjustIncr = o2.LearntAdjustIncr
	}

	return o
}

// DPLL is a DPLL satisfiability (SAT) solver.  DPLL is a synchronous structure
// and its variables/methods must not be used concurrently from multiple
// goroutines.
type DPLL struct {
	Opt

	useExtra         bool
	addClauseFn      func(p ...Lit) bool
	garbageCollectFn func()

	startTime time.Time

	model    []LBool
	conflict []Lit

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
	ngarbLit       uint64
	nmaxLit        uint64
	ntotLit        uint64

	clauses     []*Clause
	learnt      []*Clause
	trail       []Lit // Assignment stack
	trailLim    []int // Seprarating indices for decision levels in trail
	assumptions []Lit // Set of assumptions provided by the user

	activity  []float64
	assigns   []LBool
	polarity  []bool
	upolarity []LBool
	decision  []bool
	vardata   []varData
	watches   *occLists // 'watches[lit]' is a list of constraints watching 'lit' (will go there if literal becomes true).

	orderHeap *activityQueue

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
	seen           []Seen
	analyzeStack   []shrinkLit
	analyzeToClear []Lit
	addTmp         []Lit

	maxLearnt         float64
	learntAdjustConfl float64
	learntAdjustCnt   int

	// resource constraints
	conflictBudget    int64
	propagationBudget int64
	asyncInterrupt    bool

	r *rand.Rand
}

var _ Solver = (*DPLL)(nil)

// New initializes and returns a new SAT solver.
func New(opt *Opt) *DPLL {
	d := &DPLL{}
	d.Opt = *mergeOptDefault(opt)
	d.watches = newOccLists()
	d.orderHeap = newActivityQueue(&d.activity)
	d.ok = true
	d.claIncr = 1
	d.varIncr = 1
	d.nsimpAssign = -1
	d.removeSat = true
	d.conflictBudget = -1
	d.propagationBudget = -1
	d.assigns = []LBool{LUndef}
	d.vardata = []varData{{nil, 0}}
	d.activity = []float64{0}
	d.seen = []Seen{0}
	d.polarity = []bool{true}
	d.upolarity = []LBool{LUndef}
	d.decision = []bool{false}
	d.initRand()
	return d
}

func (d *DPLL) clearStack() {
	d.analyzeStack = d.analyzeStack[:0]
}

func (d *DPLL) pushShrinkLit(i uint32, p Lit) {
	d.analyzeStack = append(d.analyzeStack, shrinkLit{i, p})
}

func (d *DPLL) initRand() {
	seed := d.RandSeed
	if seed == 0 {
		seed = 0x1234C0DE
	}
	src := rand.NewSource(seed)
	d.r = rand.New(src)
}

func (d *DPLL) randf64() float64 {
	return d.r.Float64()
}

func (d *DPLL) randn(n int) int {
	return d.r.Intn(n)
}

// NewVar adds a new variable. The parameters specify variable mode.
func (d *DPLL) NewVar(upol LBool, dvar bool) Var {
	var v Var
	if len(d.freeVars) > 0 {
		v = d.freeVars[len(d.freeVars)-1]
		d.freeVars = d.freeVars[:len(d.freeVars)-1]
		d.assigns[v] = LUndef
		d.vardata[v] = varData{nil, 0}
		d.activity[v] = d.varInitActivity()
		d.seen[v] = SeenUndef
		d.polarity[v] = true
		d.upolarity[v] = upol
	} else {
		d.varNext++
		v = d.varNext
		d.assigns = append(d.assigns, LUndef)
		d.vardata = append(d.vardata, varData{nil, 0})
		d.activity = append(d.activity, d.varInitActivity())
		d.seen = append(d.seen, SeenUndef)
		d.polarity = append(d.polarity, true)
		d.upolarity = append(d.upolarity, upol)
		d.decision = append(d.decision, false)
	}

	d.watches.Init(Literal(v, false))
	d.watches.Init(Literal(v, true))

	d.SetDecision(v, dvar)

	return v
}

// ReleaseVar makes lit true.  The caller promises to never refer to variable
// after ReleaseVar is called.
func (d *DPLL) ReleaseVar(lit Lit) {
	if d.Value(lit.Var()).IsUndef() {
		d.addClauseAlias(lit)
		d.releasedVars = append(d.releasedVars, lit.Var())
	}
}

func (d *DPLL) varInitActivity() float64 {
	if d.RandInitAct {
		return d.randf64() * 1e-5
	}
	return 0
}

// SetPolarity declares which polarity the decision heuristic should use for a
// variable. Requires mode 'polarity_user'.
func (d *DPLL) SetPolarity(v Var, upol LBool) {
	d.upolarity[v] = upol
}

// SetDecision declares if a variable should be eligible for selection in the
// decision heuristic.
func (d *DPLL) SetDecision(v Var, ok bool) {
	if ok && !d.decision[v] {
		d.ndecVars++
	} else if !ok && d.decision[v] {
		d.ndecVars--
	}

	d.decision[v] = ok
	d.insertVarOrder(v)
}

func (d *DPLL) isSeen(v Var) bool {
	return d.seen[v].IsSeen()
}

func (d *DPLL) reason(v Var) *Clause {
	return d.vardata[v].Reason
}

func (d *DPLL) level(v Var) int {
	return d.vardata[v].Level
}

func (d *DPLL) insertVarOrder(v Var) {
	if !d.orderHeap.Contains(v) && d.decision[v] {
		d.orderHeap.Push(v)
	}
}

func (d *DPLL) varDecayActivity() {
	d.varIncr /= d.VarDecay
}

func (d *DPLL) varBumpActivity(v Var) {
	d.varBumpActivityIncr(v, d.varIncr)
}

func (d *DPLL) varBumpActivityIncr(v Var, inc float64) {
	d.activity[v] += inc
	if d.activity[v] > 1e100 {
		for i := range d.activity {
			d.activity[i] *= 1e-100
		}
		d.varIncr *= 1e-100
	}

	if d.orderHeap.Contains(v) {
		d.orderHeap.Decrease(v)
	}
}

func (d *DPLL) claDecayActivity() {
	d.claIncr /= d.ClauseDecay
}

func (d *DPLL) claBumpActivity(c *Clause) {
	d.claBumpActivityIncr(c, d.claIncr)
}

func (d *DPLL) claBumpActivityIncr(c *Clause, inc float64) {
	c.Activity += inc
	if c.Activity > 1e20 {
		for i := range d.learnt {
			d.learnt[i].Activity *= 1e-20
		}
		d.claIncr *= 1e-20
	}
}

// Value returns the current value of v.
func (d *DPLL) Value(v Var) LBool {
	return d.assigns[v]
}

// ValueLit returns the current value of lit.
func (d *DPLL) ValueLit(lit Lit) LBool {
	if d.assigns[lit.Var()].IsUndef() {
		return LUndef
	}
	if lit.IsNeg() {
		return LiftBool(d.assigns[lit.Var()].IsFalse())
	}
	return LiftBool(d.assigns[lit.Var()].IsTrue())
}

// ValueModel returns the value of v in the last model.  The last call to Solve
// must have returned true.
func (d *DPLL) ValueModel(v Var) LBool {
	return d.model[v]
}

// ValueLitModel returns the value of lit in the last model.  The last call to
// Solve must have returned true.
func (d *DPLL) ValueLitModel(lit Lit) LBool {
	return d.model[lit.Var()].Xor(lit.IsNeg())
}

// NumAssign returns the number of assigned literals.
func (d *DPLL) NumAssign() int {
	return len(d.trail)
}

// NumClause returns the number of original clauses.
func (d *DPLL) NumClause() int {
	return int(d.nclauses)
}

// NumLearn returns the number of learnt clauses.
func (d *DPLL) NumLearn() int {
	return int(d.nlearnt)
}

// NumVar returns the number of variables.
func (d *DPLL) NumVar() int {
	return int(d.varNext)
}

// NumVarFree returns the number of free variables.
//
// BUG:
// Computation is not quite correct.  Try to calculate right instead of
// adapting it like below.
func (d *DPLL) NumVarFree() int {
	if len(d.trailLim) == 0 {
		return int(d.ndecVars) - len(d.trail)
	}
	return int(d.ndecVars) - d.trailLim[0]
}

func (d *DPLL) enqueue(p Lit, from *Clause) bool {
	if !d.ValueLit(p).IsUndef() {
		return d.ValueLit(p).IsTrue()
	}
	d.uncheckedEnqueue(p, from)
	return true
}

func (d *DPLL) uncheckedEnqueue(p Lit, from *Clause) {
	if !d.ValueLit(p).IsUndef() {
		panic("enqueued literal has a value")
	}
	if d.Verbosity >= 3 {
		log.Printf("ASSIGN %v @ %d", p, d.decisionLevel())
	}
	d.assigns[p.Var()] = LiftBool(!p.IsNeg())
	d.vardata[p.Var()] = varData{from, d.decisionLevel()}
	d.trail = append(d.trail, p)
}

func isRemoved(c *Clause) bool {
	return c.Mark == MarkDel
}

func (d *DPLL) locked(c *Clause) bool {
	if !d.ValueLit(c.Lit[0]).IsTrue() {
		return false
	}
	reason := d.reason(c.Lit[0].Var())
	if reason == nil {
		return false
	}

	// BUG: what does the line below mean?
	// d.ca.lea(reason(var(c[0]))) == &c;
	return true
}

func (d *DPLL) newDecisionLevel() {
	d.trailLim = append(d.trailLim, len(d.trail))
}

func (d *DPLL) decisionLevel() int {
	return len(d.trailLim)
}

func (d *DPLL) abstractLevel(v Var) uint32 {
	return 1 << (uint(d.level(v)) & 31)
}

func (d *DPLL) setConflBudget(n int64) {
	d.conflictBudget = int64(d.nconflicts) + n
}

func (d *DPLL) setPropBudget(n int64) {
	d.propagationBudget = int64(d.npropogations) + n
}

func (d *DPLL) interrupt() {
	d.asyncInterrupt = true
}

func (d *DPLL) clearInterrupt() {
	d.asyncInterrupt = false
}

func (d *DPLL) budgetOff() {
	d.conflictBudget = -1
	d.propagationBudget = -1
}

func (d *DPLL) withinBudget() bool {
	return !d.asyncInterrupt &&
		(d.conflictBudget < 0 || d.nconflicts < uint64(d.conflictBudget)) &&
		(d.propagationBudget < 0 || d.npropogations < uint64(d.propagationBudget))
}

func (d *DPLL) addClauseAlias(c ...Lit) bool {
	if d.addClauseFn != nil {
		return d.addClauseFn(c...)
	}
	return d.AddClause(c...)
}

// AddClause adds a CNF clause containing the given literals.  The literals in
// c will be sorted.
func (d *DPLL) AddClause(c ...Lit) bool {
	return d.addClause(c)
}

// AddClauseCopy adds a CNF clause containing a sorted copy the given literals.
func (d *DPLL) AddClauseCopy(c []Lit) bool {
	return d.addClause(c)
}

func (d *DPLL) newClause(ps []Lit, learnt bool) *Clause {
	return NewClause(ps, d.useExtra, learnt)
}

func (d *DPLL) newClauseFrom(c *Clause) *Clause {
	return NewClauseFrom(c, d.useExtra)
}

func (d *DPLL) removeClause(c *Clause) {
	d.detachClause(c, false)
	if d.locked(c) {
		d.vardata[c.Lit[0].Var()].Reason = nil
	}
	c.Mark = MarkDel
	// no need to free
}

func (d *DPLL) detachClause(c *Clause, strict bool) {
	if c.Len() <= 1 {
		panic("small clause")
	}

	// stict or lazy detatching
	if strict {
		d.watches.Remove(c.Lit[0].Inverse(), watcher{c, c.Lit[1]})
		d.watches.Remove(c.Lit[1].Inverse(), watcher{c, c.Lit[0]})
	} else {
		d.watches.Smudge(c.Lit[0].Inverse())
		d.watches.Smudge(c.Lit[1].Inverse())
	}

	d.ngarbLit += uint64(c.Len())
	if c.Learnt {
		d.nlearnt--
		d.nlearntLit -= uint64(c.Len())
	} else {
		d.nclauses--
		d.nclauseLit -= uint64(c.Len())
	}
}

func (d *DPLL) checkGarbage() {
	d.checkGarbageFrac(d.GarbageFrac, false)
}

func (d *DPLL) checkGarbageFrac(frac float64, force bool) {
	if force || float64(d.ngarbLit)/float64(d.nclauseLit+d.nlearntLit+d.ngarbLit) > frac {
		d.garbageCollectAlias()
	}
}

func (d *DPLL) garbageCollectAlias() {
	if d.garbageCollectFn != nil {
		d.garbageCollectFn()
	} else {
		d.garbageCollect()
	}
}

func (d *DPLL) garbageCollect() {
	numremove := d.relocAll()
	d.ngarbLit = 0
	if d.Verbosity >= 2 {
		log.Printf("===============================================================================")
		log.Printf("| Garbage collection:   %12d literals freed                                   |", numremove)
	}
}

func (d *DPLL) relocAll() (numremove int) {
	// TODO: all watchers
	d.watches.CleanAll()

	for i := range d.trail {
		v := d.trail[i].Var()

		// is it safe to call locked dangling reason check?
		if d.reason(v) != nil && (d.reason(v).Relocated || d.locked(d.reason(v))) {
			if isRemoved(d.reason(v)) {
				panic("reason removed")
			}
		}
	}

	// learnt clauses
	learnt := d.learnt[:0]
	for i := range d.learnt {
		if !isRemoved(d.learnt[i]) {
			learnt = append(learnt, d.learnt[i])
		} else {
			numremove++
		}
	}
	d.learnt = learnt

	// original clauses
	clauses := d.clauses[:0]
	for i := range d.clauses {
		if !isRemoved(d.clauses[i]) {
			clauses = append(clauses, d.clauses[i])
		} else {
			numremove++
		}
	}
	d.clauses = clauses

	return numremove
}

func (d *DPLL) satisfied(c *Clause) bool {
	for i := range c.Lit {
		if d.ValueLit(c.Lit[i]).IsTrue() {
			return true
		}
	}
	return false
}

func (d *DPLL) rebuildOrderHeap() {
	var vs []Var
	maxvar := Var(d.NumVar())
	_ = d.decision[maxvar]
	for v := Var(1); v <= maxvar; v++ {
		if d.decision[v] && d.Value(v).IsUndef() {
			vs = append(vs, v)
		}
	}
	d.orderHeap.Rebuild(vs)
}

func (d *DPLL) removeSatisfied(ptr *[]*Clause) {
	cs := *ptr

	var i, j int
	for ; i < len(cs); i++ {
		c := cs[i]
		if d.satisfied(c) {
			d.removeClause(c)
		} else {
			if !d.ValueLit(c.Lit[0]).IsUndef() || !d.ValueLit(c.Lit[1]).IsUndef() {
				panic(fmt.Sprintf("literals have value(s) %v %v", d.ValueLit(c.Lit[0]), d.ValueLit(c.Lit[1])))
			}

			// trim clause
			for k := 2; k < c.Len(); k++ {
				if d.ValueLit(c.Lit[k]).IsFalse() {
					c.Lit[k] = c.Lit[c.Len()-1]
					c.Lit = c.Lit[:c.Len()-1]
					k--
				}
			}
			cs[j] = c
			j++
		}
	}
	(*ptr) = cs[:j]
}

func (d *DPLL) reduceDB() {
	lowerLimit := d.claIncr / float64(len(d.learnt))
	sort.Sort(clausesByActivity(d.learnt))

	var i, j int
	for ; i < len(d.learnt); i++ {
		c := d.learnt[i]
		// don't delete binary or locked clauses.
		if c.Len() > 2 && !d.locked(c) && (i < len(d.learnt)/2 || c.Activity < lowerLimit) {
			d.removeClause(c)
		} else {
			d.learnt[j] = c
			j++
		}
	}
	for i := j; i < len(d.learnt); i++ {
		d.learnt[i] = nil
	}
	d.learnt = d.learnt[:j]
	d.checkGarbage()
}

type clausesByActivity []*Clause

func (cs clausesByActivity) Len() int {
	return len(cs)
}

func (cs clausesByActivity) Less(i, j int) bool {
	return cs[i].Len() > 2 && cs[j].Len() > 2 || cs[i].Activity < cs[j].Activity
}

func (cs clausesByActivity) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

// propagate all equeued facts.  If a conflict arises the conflicting clause is
// returned.  The propagation queue will be empty after propagate returns.
func (d *DPLL) propagate() *Clause {
	var conflict *Clause
	var numprops int

	for d.qhead < len(d.trail) {
		p := d.trail[d.qhead]
		pnot := p.Inverse()
		d.qhead++
		ws := d.watches.Lookup(p)
		numprops++

		var j int
	NEXTCLAUSE:
		for i := 0; i < len(ws); {
			block := ws[i].blocker
			if d.ValueLit(block).IsTrue() {
				ws[j] = ws[i]
				j++
				i++
				continue
			}

			// make sure the false literal is at index 1
			c := ws[i].c
			if c.Lit[0] == pnot {
				c.Lit[0], c.Lit[1] = c.Lit[1], pnot
			} else if c.Lit[1] != pnot {
				panic(fmt.Sprintf("index 1 %v (!= %v)", c.Lit[1], pnot))
			}
			i++

			// if first watch is true then clause is already satisfied
			first := c.Lit[0]
			if first != block && d.ValueLit(first).IsTrue() {
				ws[j] = watcher{c, first}
				j++
				continue
			}

			// look for new watch
			for k := 2; k < c.Len(); k++ {
				if !d.ValueLit(c.Lit[k]).IsFalse() {
					c.Lit[1], c.Lit[k] = c.Lit[k], pnot
					if c.Lit[1].Inverse() == p {
						ws = d.watches.Occurrences(p)
					} else {
						d.watches.Push(c.Lit[1].Inverse(), watcher{c, first})
					}
					continue NEXTCLAUSE
				}
			}

			// didn't find watch -- clause is unit under assignment
			ws[j] = watcher{c, first}
			j++
			if d.ValueLit(first).IsFalse() {
				conflict = c
				d.qhead = len(d.trail)
				j += copy(ws[j:], ws[i:])
				i = len(ws)
			} else {
				d.uncheckedEnqueue(first, c)
			}
		}
		d.watches.occs[p] = ws[:j]

	}

	d.npropogations += uint64(numprops)
	d.nsimpProps += int64(numprops)

	return conflict
}

// Simplify the clause database according to the current top-level assignment.
// Currently removal of satisfied clauses is all the simplification provided
// but additional logic may be added in forked implementations.
func (d *DPLL) Simplify() bool {
	if d.decisionLevel() != 0 {
		panic("non-root decision level")
	}

	if !d.ok || d.propagate() != nil {
		d.ok = false
		return false
	}

	if d.NumAssign() == d.nsimpAssign || d.nsimpProps > 0 {
		return true
	}

	// remove satisfied clauses
	d.removeSatisfied(&d.learnt)
	if !d.removeSat {
		// TODO
	} else {
		d.removeSatisfied(&d.clauses)

		// remove all released variables from the trail
		for i := range d.releasedVars {
			if d.isSeen(d.releasedVars[i]) {
				panic(fmt.Sprintf("seen released var: %d", d.releasedVars[i]))
			}
			d.seen[d.releasedVars[i]] = SeenSource
		}

		var j int
		for i := range d.trail {
			if !d.isSeen(d.trail[i].Var()) {
				d.trail[j] = d.trail[i]
				j++
			}
		}
		d.trail = d.trail[:j]
		d.qhead = len(d.trail)

		for i := range d.releasedVars {
			d.seen[d.releasedVars[i]] = SeenUndef
		}

		// released vars are now ready to be reused
		d.freeVars = append(d.freeVars, d.releasedVars...)
		d.releasedVars = d.releasedVars[:0]
	}
	d.checkGarbage()
	d.rebuildOrderHeap()

	d.nsimpAssign = d.NumAssign()
	d.nsimpProps = int64(d.nclauseLit + d.nlearntLit)

	return true
}

// Solve searches for a model that respects the given assupmtions.
func (d *DPLL) Solve(assump ...Lit) bool {
	d.budgetOff()
	d.assumptions = assump
	return d.solve().IsTrue()
}

// SolveLimited behaves like Solve but respects resource contraints (?).
func (d *DPLL) SolveLimited(assump ...Lit) LBool {
	d.assumptions = assump
	return d.solve()
}

// solve searches for a model that respects the d.assumptions
func (d *DPLL) solve() LBool {
	d.startTime = time.Now()
	if !d.ok {
		return LFalse
	}

	d.model = nil
	d.conflict = nil

	defer d.checkGarbageFrac(0, true)

	d.nsolves++
	d.maxLearnt = float64(d.NumClause()) * d.LearntFraction
	if d.maxLearnt < float64(d.MinLearnt) {
		d.maxLearnt = float64(d.MinLearnt)
	}

	d.learntAdjustConfl = float64(d.LearntAdjustConfl)
	d.learntAdjustCnt = d.LearntAdjustConfl
	status := LUndef

	if d.Verbosity >= 1 {
		log.Printf("============================[ Search Statistics ]==============================")
		log.Printf("| Conflicts |          ORIGINAL         |          LEARNT          | Progress |")
		log.Printf("|           |    Vars  Clauses Literals |    Limit  Clauses Lit/Cl |          |")
		log.Printf("===============================================================================")
	}

	// do search
	for currRestarts := 0; status.IsUndef(); currRestarts++ {
		var restBase float64
		if d.NoLubyRestart {
			restBase = math.Pow(d.RestartIncr, float64(currRestarts))
		} else {
			restBase = Luby(d.RestartIncr, currRestarts)
		}
		status = d.search(int(restBase * float64(d.RestartFirst)))
		if !d.withinBudget() {
			break
		}
	}

	if d.Verbosity >= 1 {
		log.Printf("===============================================================================")
	}

	if status.IsTrue() {
		d.model = make([]LBool, d.NumVar()+1)
		copy(d.model, d.assigns)
	} else if status.IsFalse() && len(d.conflict) == 0 {
		d.ok = false
	}

	d.cancelUntil(0)
	return status
}

// search finds a model with at most maxconflict conflicts.  If maxconflict is
// negative then search tolerates any number of conflicts.
//
// search returns LTrue if all variables are decision variables, which implies
// the clause set is satisfiable.  Search returns LFalse if the clause is
// unsatisfiable and LUndef if the maxconflict limit was reached.
func (d *DPLL) search(maxconflict int) LBool {
	if !d.ok {
		panic("not okay")
	}

	var numconflict int

	d.nstarts++

	for {
		conflict := d.propagate()
		if conflict != nil {
			d.nconflicts++
			numconflict++
			if d.decisionLevel() == 0 {
				return LFalse
			}
			learnt, btlevel := d.analyze(conflict)
			d.cancelUntil(btlevel)

			if len(learnt) == 1 {
				d.uncheckedEnqueue(learnt[0], nil)
			} else {
				c := d.newClause(learnt, true)
				d.learnt = append(d.learnt, c)
				d.attachClause(c)
				d.claBumpActivity(c)
				d.uncheckedEnqueue(learnt[0], c)
			}

			d.varDecayActivity()
			d.claDecayActivity()

			d.learntAdjustCnt--
			if d.learntAdjustCnt == 0 {
				d.learntAdjustConfl *= d.LearntAdjustIncr
				d.learntAdjustCnt = int(d.learntAdjustConfl)
				d.maxLearnt *= d.LearntIncr

				if d.Verbosity >= 1 {
					decadj := len(d.trail)
					if len(d.trailLim) > 0 {
						decadj = d.trailLim[0]
					}
					log.Printf("| %9d | %7d %8d %8d | %8d %8d %6.0f | %6.3f %% |",
						d.nconflicts,
						d.ndecVars-uint64(decadj), d.NumClause(), d.nclauseLit,
						int(d.maxLearnt), d.NumLearn(), float64(d.nlearntLit)/float64(d.NumLearn()),
						d.progressEstimate()*100,
					)
				}
			}
		} else { // no conflict; c == nil
			if (maxconflict >= 0 && numconflict >= maxconflict) || !d.withinBudget() {
				// too many conflicts
				d.progress = d.progressEstimate()
				d.cancelUntil(0)
				return LUndef
			}

			// simplify the set of problem clauses
			if d.decisionLevel() == 0 && !d.Simplify() {
				return LFalse
			}

			// reduce the set of learnt clauses
			if float64(len(d.learnt)-d.NumAssign()) >= d.maxLearnt {
				d.reduceDB()
			}

			// peform the user provided assumptions
			var next Lit
			for d.decisionLevel() < len(d.assumptions) {
				p := d.assumptions[d.decisionLevel()]
				if d.ValueLit(p).IsTrue() {
					// dummy decision level
					d.newDecisionLevel()
				} else if d.ValueLit(p).IsFalse() {
					d.conflict = d.analyzeFinal(p).slice()
					return LFalse
				} else {
					next = p
					break
				}
			}

			if next.IsUndef() {
				d.ndecisions++
				next = d.pickBranchLit()
				if next.IsUndef() {
					// model found
					return LTrue
				}
			}

			d.newDecisionLevel()
			d.uncheckedEnqueue(next, nil)
		}
	}
}

func (d *DPLL) progressEstimate() float64 {
	progress := 0.0
	F := 1.0 / float64(d.NumVar())

	for i := 0; i <= d.decisionLevel(); i++ {
		beg := 0
		if i > 0 {
			beg = d.trailLim[i-1]
		}
		end := len(d.trail)
		if i != d.decisionLevel() {
			end = d.trailLim[i]
		}
		progress += math.Pow(F, float64(i)) * float64(end-beg)
	}

	return progress / float64(d.NumVar())
}

// PrintStats prints statistics about solving meant to be called after solving
// has terminated.
func (d *DPLL) PrintStats() {
	runsec := d.runtime()
	var memused float64 // TODO
	log.Printf("restarts              : %d", d.nstarts)
	log.Printf("conflicts             : %-12d   (%.0f / sec)", d.nconflicts, float64(d.nconflicts)/runsec)
	log.Printf("decisions             : %-12d   (%.0f / sec) (%4.2f %% random)", d.ndecisions, float64(d.ndecisions)/runsec, float64(d.nrandDecisions)*100.0)
	log.Printf("propagations          : %-12d   (%.0f / sec)", d.npropogations, float64(d.npropogations)/runsec)
	log.Printf("conflict literals     : %-12d   (%4.2f %% deleted)", d.ntotLit, float64(d.nmaxLit-d.ntotLit)*100.0/float64(d.nmaxLit))
	if memused != 0 {
		log.Printf("memory used           : %.2f MB", memused)
	}
	log.Printf("runtime               : %0.3g sec", runsec)
}

// runtime returns the time since in number of seconds
func (d *DPLL) runtime() float64 {
	dur := time.Since(d.startTime)
	// avoid rounding errors from large floating point arithmetic by computing
	// seconds and fractional seconds separately.
	return float64(dur/time.Second) + float64(dur%time.Second)/float64(time.Second)
}

// Okay returns false if the solver is in a conflicting state.
func (d *DPLL) Okay() bool {
	return d.ok
}

// Implies returns any assignments implied by the given assumptions.  If
// assumptions are plausible then the second return argument of Implies is
// true.  Otherwise Implies returns false if the assumptions result in a
// contradiction.
func (d *DPLL) Implies(assumps []Lit) (assign []Lit, ok bool) {
	d.trailLim = append(d.trailLim, len(d.trail))
	for _, a := range assumps {
		if d.ValueLit(a).IsFalse() {
			d.cancelUntil(0)
			return nil, false
		}
		if d.ValueLit(a).IsUndef() {
			d.uncheckedEnqueue(a, nil)
		}
	}

	before := len(d.trail)
	ok = true
	if d.propagate() != nil {
		ok = false
	} else {
		for _, p := range d.trail[before:] {
			assign = append(assign, p)
		}
	}

	d.cancelUntil(0)
	return assign, ok
}

// Model returns assignments found in the last call to Solve.  If Solve could
// not find a model (perhaps under assupmtions) Model returns nil.
func (d *DPLL) Model() []LBool {
	return d.model
}

// Conflict returns the final, non-empty clause expressed in assumptions if
// Solve could not find a model.  If the last call to Solve found a model then
// Conflict returns nil.
func (d *DPLL) Conflict() []Lit {
	return d.conflict
}

func (d *DPLL) isRedundant(p Lit) bool {
	if d.seen[p.Var()] != SeenUndef && d.seen[p.Var()] != SeenSource {
		panic(fmt.Sprintf("variable seen: %d %d", p.Var(), d.seen[p.Var()]))
	}

	c := d.reason(p.Var())
	if c == nil {
		panic("missing reason")
	}

	d.clearStack()
	for i := uint32(1); ; i++ {
		// TODO verify at clause creation that len in less than math.MaxUint32
		if i < uint32(c.Len()) {
			// check the parent of p
			par := c.Lit[i]

			if d.level(par.Var()) == 0 || d.seen[par.Var()] == SeenSource || d.seen[par.Var()] == SeenRemovable {
				continue
			}

			// determine if the variable cannot be removed for local reasons
			if d.reason(par.Var()) == nil || d.seen[par.Var()] == SeenFailed {
				d.pushShrinkLit(0, p)
				for i := range d.analyzeStack { // shadow outer index var
					if d.seen[d.analyzeStack[i].p.Var()] == SeenUndef {
						d.seen[d.analyzeStack[i].p.Var()] = SeenFailed
						d.analyzeToClear = append(d.analyzeToClear, d.analyzeStack[i].p)
					}
				}
				return false
			}

			// could be redundant -- setup for a recursive check of the parent
			d.pushShrinkLit(i, p)
			i, p, c = 0, par, d.reason(par.Var())
		} else {
			if d.seen[p.Var()] == SeenUndef {
				d.seen[p.Var()] = SeenRemovable
				d.analyzeToClear = append(d.analyzeToClear, p)
			}

			if len(d.analyzeStack) == 0 {
				break
			}

			// continue with the top item on the stack
			i = d.analyzeStack[len(d.analyzeStack)-1].i
			p = d.analyzeStack[len(d.analyzeStack)-1].p
			c = d.reason(p.Var())

			d.analyzeStack = d.analyzeStack[:len(d.analyzeStack)-1]
		}
	}

	return true
}

// analyzeFinal expresses the final conflict in terms of assumptions.
// analyzeFinal returns the set of assupmtions that led to the assignment of p.
// analyzeFinal may return an empty set of assupmtions.
func (d *DPLL) analyzeFinal(p Lit) litSet {
	conflict := newLitSet()
	conflict.insert(p)

	if d.decisionLevel() == 0 {
		return conflict
	}

	d.seen[p.Var()] = SeenSource

	for i := len(d.trail) - 1; i >= d.trailLim[0]; i-- {
		v := d.trail[i].Var()
		if d.isSeen(v) {
			c := d.reason(v)
			if c == nil {
				if d.level(v) > 0 {
					panic(fmt.Sprintf("var level: %d", d.level(v)))
				}
				conflict.insert(d.trail[i].Inverse())
			} else {
				for j := 1; j < c.Len(); j++ {
					if d.level(c.Lit[j].Var()) > 0 {
						d.seen[c.Lit[j].Var()] = SeenSource
					}
				}
			}
			d.seen[v] = SeenUndef
		}
	}

	d.seen[p.Var()] = SeenUndef

	return conflict
}

// analyze a conflict to produce a learnt clause and backtrack level.  The
// current decision level must be greater than the root level.  The first
// literal in the resulting clause is the asserting literal at btlevel.  If
// outLearnt contains multiple variables outLearnt[1] has the maximum decision
// level of remaining variables.
func (d *DPLL) analyze(confl *Clause) (outLearnt []Lit, btlevel int) {
	pathc := 0
	p := LitUndef

	// generate the conflict clause
	outLearnt = append(outLearnt, LitUndef)
	index := len(d.trail) - 1
	for {
		if confl.Learnt {
			d.claBumpActivity(confl)
		}
		var j int
		if !p.IsUndef() {
			j = 1
		}
		for ; j < confl.Len(); j++ {
			q := confl.Lit[j]
			if !d.isSeen(q.Var()) && d.level(q.Var()) > 0 {
				d.varBumpActivity(q.Var())
				d.seen[q.Var()] = SeenSource
				if d.level(q.Var()) >= d.decisionLevel() {
					pathc++
				} else {
					outLearnt = append(outLearnt, q)
				}
			}
		}
		for !d.isSeen(d.trail[index].Var()) {
			index--
		}
		index--
		p = d.trail[index+1]
		confl = d.reason(p.Var())
		d.seen[p.Var()] = SeenUndef
		pathc--

		if pathc <= 0 {
			break
		}
	}
	outLearnt[0] = p.Inverse()

	// simplify the conflict clause
	d.analyzeToClear = append(d.analyzeToClear[:0], outLearnt...)
	var j int
	if d.CCMin == CCMinDeep {
		j = 1
		for i := 1; i < len(outLearnt); i++ {
			if d.reason(outLearnt[i].Var()) == nil || !d.isRedundant(outLearnt[i]) {
				outLearnt[j] = outLearnt[i]
				j++
			}
		}
	} else if d.CCMin == CCMinBasic {
		j = 1
		for i := 1; i < len(outLearnt); i++ {
			c := d.reason(outLearnt[i].Var())
			if c == nil {
				outLearnt[j] = outLearnt[i]
				j++
			} else {
				for k := 1; k < c.Len(); k++ {
					if !d.isSeen(c.Lit[k].Var()) && d.level(c.Lit[k].Var()) > 0 {
						outLearnt[j] = outLearnt[i]
						j++
						break
					}
				}
			}
		}
	} else {
		j = len(outLearnt)
	}

	d.nmaxLit += uint64(len(outLearnt))
	outLearnt = outLearnt[:j]
	d.ntotLit += uint64(j)

	// find the correct backtrack level
	if len(outLearnt) != 1 {
		imax := 1
		for i := 2; i < len(outLearnt); i++ {
			if d.level(outLearnt[i].Var()) > d.level(outLearnt[imax].Var()) {
				imax = i
			}
		}
		outLearnt[imax], outLearnt[1] = outLearnt[1], outLearnt[imax]
		btlevel = d.level(outLearnt[1].Var())
	}

	for j := range d.analyzeToClear {
		d.seen[d.analyzeToClear[j].Var()] = SeenUndef
	}
	// d.seen is now cleared

	return outLearnt, btlevel
}

func (d *DPLL) pickBranchLit() Lit {
	next := VarUndef

	if d.randf64() < d.RandVarFreq && d.orderHeap.Len() == 0 {
		next = d.orderHeap.vars[d.randn(d.orderHeap.Len())]
		if d.Value(next).IsUndef() && d.decision[next] {
			d.nrandDecisions++
		}
	}

	for next.IsUndef() || !d.Value(next).IsUndef() || !d.decision[next] {
		if d.orderHeap.Len() == 0 {
			next = VarUndef
			break
		}
		next, _ = d.orderHeap.RemoveMax()
	}

	if next.IsUndef() {
		return LitUndef
	}
	if !d.upolarity[next].IsUndef() {
		return Literal(next, d.upolarity[next].IsTrue())
	}
	if d.RandPol {
		return Literal(next, d.randf64() < 0.5)
	}
	return Literal(next, d.polarity[next])
}

func (d *DPLL) addClause(c []Lit) bool {
	if d.decisionLevel() != 0 {
		panic("decision level")
	}
	if !d.ok {
		return false
	}

	sort.Sort(litSlice(c))

	var j int
	for i, p := 0, LitUndef; i < len(c); i++ {
		if d.ValueLit(c[i]).IsTrue() || c[i] == p.Inverse() {
			return true
		} else if !d.ValueLit(c[i]).IsFalse() && c[i] != p {
			p = c[i]
			c[j] = p
			j++
		}
	}
	c = c[:j]

	switch len(c) {
	case 0:
		d.ok = false
		return false
	case 1:
		d.uncheckedEnqueue(c[0], nil)
		d.ok = d.propagate() == nil
		return d.ok
	default:
		c := d.newClause(c, false)
		d.clauses = append(d.clauses, c)
		d.attachClause(c)
	}

	return true
}

func (d *DPLL) attachClause(c *Clause) {
	if c.Len() < 2 {
		panic("small clause")
	}

	d.watches.Push(c.Lit[0].Inverse(), watcher{c, c.Lit[1]})
	d.watches.Push(c.Lit[1].Inverse(), watcher{c, c.Lit[0]})

	if c.Learnt {
		d.nlearnt++
		d.nlearntLit += uint64(c.Len())
	} else {
		d.nclauses++
		d.nclauseLit += uint64(c.Len())
	}
}

func (d *DPLL) cancelUntil(level int) {
	if d.decisionLevel() <= level {
		return
	}
	for c := len(d.trail) - 1; c >= d.trailLim[level]; c-- {
		v := d.trail[c].Var()
		if d.Verbosity >= 3 {
			log.Printf("UNASSIGN %v", v)
		}
		d.assigns[v] = LUndef
		if d.PhaseSaving > 1 || d.PhaseSaving == 1 && c > d.trailLim[len(d.trailLim)-1] {
			d.polarity[v] = d.trail[c].IsNeg()
		}
		d.insertVarOrder(v)
	}
	d.qhead = d.trailLim[level]
	d.trail = d.trail[:d.trailLim[level]]
	d.trailLim = d.trailLim[:level]
}

// CCMinMode controls conflict clause minimization.
type CCMinMode int

// Available conflict clause minimization modes.
const (
	CCMinInvalid CCMinMode = iota
	CCMinNone
	CCMinBasic
	CCMinDeep
)

// PhaseSavingLevel controls the amount of phase saving.
type PhaseSavingLevel int

// Avalaible levels of phase saving.
const (
	PhaseSavingInvalid PhaseSavingLevel = iota
	PhaseSavingNone
	PhaseSavingLimited
	PhaseSavingFull
)

type varData struct {
	Reason *Clause
	Level  int
}

type shrinkLit struct {
	i uint32
	p Lit
}

// Seen marks how a variable was seen
type Seen uint8

// Avaible methods of being seen.
const (
	SeenUndef Seen = iota
	SeenSource
	SeenRemovable
	SeenFailed
)

// IsSeen returns true if s is non-zero
func (s Seen) IsSeen() bool {
	return s != SeenUndef
}

// Luby uses the Luby algorithm to generate a restart sequence from factor y
// and current number of restarts x.
//
//		Finite subsequences of the Luby-sequence:
//		0: 1
//		1: 1 1 2
//		2: 1 1 2 1 1 2 4
//		3: 1 1 2 1 1 2 4 1 1 2 1 1 2 4 8
//		...
func Luby(y float64, x int) float64 {
	size, seq := 1, 0
	for size < x+1 {
		seq++
		size = (size << 1) + 1
	}
	for size-1 != x {
		size = (size - 1) >> 1
		seq--
		x = x % size
	}
	return math.Pow(y, float64(seq))
}
