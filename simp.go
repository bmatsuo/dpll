// Copyright 2016 Bryan Matsuo
//
// Use of this software is governed by the MIT license.  A copy of the license
// agreement can be found in the LICENSE file distributed with this software.

package dpll

import (
	"fmt"
	"log"
	"os"
)

// Marks used by the Simp solver
const (
	MarkTouch Mark = 2
)

// SimpOpt contain options for the Simp solver.
type SimpOpt struct {
	Asymm            bool    // Shrink clauses by asymmetric branching
	RCheck           bool    // Check if a clause is already implied (costly)
	NoElim           bool    // Do not perform variable elimination
	NoExtend         bool    // When true the caller does not need to know the full model
	Grow             int     // Allow a variable elimination step to grow by the given number of clauses
	ClauseLimit      int     // Variables are not eliminated if it produces a resolvent with a length above this limit
	SubsumptionLimit int     // Do not check if subsumption against a clause larger than this
	SimpGarbageFrac  float64 // The fraction of wasted memory allowed before a garbage collection is triggered during simplification
}

var simpOptDefault = &SimpOpt{
	ClauseLimit:      20,
	SubsumptionLimit: 1000,
	SimpGarbageFrac:  0.5,
}

func mergeSimpOpt(o1, o2 *SimpOpt) *SimpOpt {
	o := &SimpOpt{}
	*o = *o1
	if o2 == nil {
		return o
	}
	if o2.Asymm {
		o.Asymm = true
	}
	if o2.RCheck {
		o.RCheck = true
	}
	if o2.NoElim {
		o.NoElim = true
	}
	if o2.NoExtend {
		o.NoExtend = true
	}
	if o2.Grow != 0 {
		o.Grow = o2.Grow
	}
	if o2.ClauseLimit != 0 {
		o.ClauseLimit = o2.ClauseLimit
	}
	if o2.SubsumptionLimit != 0 {
		o.SubsumptionLimit = o2.SubsumptionLimit
	}
	if o2.SimpGarbageFrac != 0 {
		o.SubsumptionLimit = o2.SubsumptionLimit
	}
	return o
}

// Simp is a solver that simplifies clauses before solving.  Do not call
// runtime.SetFinalizer on a Simp it will never be garbage collected.
type Simp struct {
	SimpOpt

	nmerge    int
	nasymmlit int
	nelimvars int

	// solver state
	elimOrder     int
	useSimp       bool
	maxSimpVar    Var
	elimClauses   []uint32
	touched       []bool
	occurs        *clauseOccLists
	numOcc        []int
	elimHeap      *elimQueue
	subQueue      *clauseQueue
	frozen        []bool
	frozenVars    []Var
	eliminated    []bool
	bwdsubAssigns int
	numTouched    int

	bwdsubTempUnit *Clause

	d *DPLL
}

var _ Solver = (*Simp)(nil)

// NewSimp returns a new simplifying solver initialized with the given search
// and simplification options.
func NewSimp(opt *Opt, simpOpt *SimpOpt) *Simp {
	d := New(opt)
	d.useExtra = true
	d.removeSat = false

	s := &Simp{
		SimpOpt:        *mergeSimpOpt(simpOptDefault, simpOpt),
		elimOrder:      1,
		useSimp:        true,
		occurs:         newClauseOccLists(),
		subQueue:       newClauseQueue(1),
		numOcc:         make([]int, 2),
		frozen:         make([]bool, 1),
		eliminated:     make([]bool, 1),
		touched:        make([]bool, 1),
		bwdsubTempUnit: d.newClause([]Lit{0}, false),
		d:              d,
	}
	s.elimHeap = newElimQueue(&s.numOcc)

	d.addClauseFn = s.AddClause
	d.removeClauseFn = s.removeClause
	d.garbageCollectFn = s.garbageCollect

	return s
}

// NumClause implements Solver
func (s *Simp) NumClause() int {
	return s.d.NumClause()
}

// NumVar implements Solver
func (s *Simp) NumVar() int {
	return s.d.NumVar()
}

// NewVar behaves like DPLL.NewVar.
func (s *Simp) NewVar(upol LBool, dvar bool) Var {
	v := s.d.NewVar(upol, dvar)

	s.frozen = append(s.frozen, false)
	s.eliminated = append(s.frozen, false)
	if s.useSimp {
		// because numOcc maps literals to counts the new variable will take up
		// the next two positions for its positive and negative literals
		// respectively.
		s.numOcc = append(s.numOcc, 0)
		s.numOcc = append(s.numOcc, 0)

		s.occurs.Init(v)
		s.touched = append(s.touched, false)
		s.elimHeap.Push(v)
	}

	return v
}

// ReleaseVar behaves like DPLL.ReleaseVar.  An eliminated variable cannot be
// released.
func (s *Simp) ReleaseVar(p Lit) {
	if s.IsEliminated(p.Var()) {
		panic("cannot release eliminated variable")
	}
	if !s.useSimp && p.Var() > s.maxSimpVar {
		s.d.ReleaseVar(p)
	} else {
		// don't allow the variable to be reused
		s.d.AddClause(p)
	}
}

// FreezeVar is an alternative to SetFrozen
func (s *Simp) FreezeVar(v Var) {
	if !s.frozen[v] {
		s.frozen[v] = true
		s.frozenVars = append(s.frozenVars, v)
	}
}

// Thaw unfreezes all variables
func (s *Simp) Thaw() {
	if !s.useSimp {
		for i := range s.frozenVars {
			s.frozen[s.frozenVars[i]] = false
		}
	} else {
		for i := range s.frozenVars {
			s.frozen[s.frozenVars[i]] = false
			s.updateElimHeap(s.frozenVars[i])
		}
	}
	s.frozenVars = s.frozenVars[:0]
}

// SetFrozen lets variables be frozen, preventing them from being eliminated.
func (s *Simp) SetFrozen(v Var, frozen bool) {
	s.frozen[v] = frozen
	if s.useSimp && !frozen {
		s.updateElimHeap(v)
	}
}

// IsEliminated returns true if the v was eliminated.
func (s *Simp) IsEliminated(v Var) bool {
	return s.eliminated[v]
}

func (s *Simp) updateElimHeap(v Var) {
	if !s.useSimp {
		panic("simplification disabled")
	}
	if s.elimHeap.Contains(v) || (!s.frozen[v] && !s.IsEliminated(v) && s.d.Value(v).IsUndef()) {
		s.elimHeap.Update(v)
	}
}

// AddClause behaves like DPLL.AddClause.  The literals given to AddClause
// cannot contain eliminated variables.
func (s *Simp) AddClause(ps ...Lit) bool {
	for i := range ps {
		if s.IsEliminated(ps[i].Var()) {
			panic("clause contains eliminated variable")
		}
	}

	numclause := len(s.d.clauses)

	if s.RCheck && s.implied(ps) {
		return true
	}

	if !s.d.AddClause(ps...) {
		return false
	}

	if s.useSimp && len(s.d.clauses) == numclause+1 {
		c := s.d.clauses[numclause]
		s.subQueue.Insert(c)
		for _, p := range c.Lit {
			s.occurs.Push(p.Var(), c)
			s.numOcc[p]++
			s.touched[p.Var()] = true
			s.numTouched++
			if s.elimHeap.Contains(p.Var()) {
				s.elimHeap.Fix(p.Var())
			}
		}
	}

	return true
}

// Okay returns true if s hasn't yet found a contradiction
func (s *Simp) Okay() bool {
	return s.d.ok
}

// PrintStats prints solver stats after Solve has returned
func (s *Simp) PrintStats() {
	s.d.PrintStats()
}

// SolveSimp is like Solve but allows something...
func (s *Simp) SolveSimp(assump []Lit, presimp bool, turnOffSimp bool) bool {
	s.d.budgetOff()
	s.d.assumptions = assump
	return s.solve(presimp, turnOffSimp).IsTrue()
}

// SolveLimitedSimp is like SolveLimited but allows something...
func (s *Simp) SolveLimitedSimp(assump []Lit, presimp bool, turnOffSimp bool) LBool {
	s.d.assumptions = assump
	return s.solve(presimp, turnOffSimp)
}

// Solve implements Solver
func (s *Simp) Solve(assump ...Lit) bool {
	return s.SolveSimp(assump, false, false)
}

// SolveLimited implements Solver
func (s *Simp) SolveLimited(assump ...Lit) LBool {
	return s.SolveLimitedSimp(assump, false, false)
}

func (s *Simp) solve(doSimp, turnOffSimp bool) LBool {
	if doSimp && !s.useSimp {
		doSimp = false
	}

	var extraFrozen []Var
	result := LTrue

	if doSimp {
		for i := range s.d.assumptions {
			v := s.d.assumptions[i].Var()
			if s.IsEliminated(v) {
				panic("assumption includes eliminated variable")
			}

			if !s.frozen[v] {
				s.SetFrozen(v, true)
				extraFrozen = append(extraFrozen, v)
			}
		}

		result = LiftBool(s.Eliminate(turnOffSimp))
	}

	if result.IsTrue() {
		result = s.d.solve()
	} else if s.d.Verbosity >= 1 {
		log.Printf("===============================================================================")
	}

	if result.IsTrue() && !s.NoExtend {
		s.extendModel()
	}

	if doSimp {
		for i := range extraFrozen {
			s.SetFrozen(extraFrozen[i], false)
		}
	}

	return result
}

func (s *Simp) removeClause(c *Clause) {
	if s.useSimp {
		for _, p := range c.Lit {
			s.numOcc[p]--
			s.updateElimHeap(p.Var())
			s.occurs.Smudge(p.Var())
		}
	}
	s.d.removeClause(c)
}

// Eliminate performs simplification and variable elimination turnOffElim
// should be true the last time that Eliminate is called to free memory used
// during variable elimination.
func (s *Simp) Eliminate(turnOffElim bool) bool {
	if !s.d.Simplify() {
		return false
	}
	if !s.useSimp {
		return true
	}

	for s.numTouched > 0 || s.bwdsubAssigns < len(s.d.trail) || s.elimHeap.Len() > 0 {
		s.gatherTouchedClauses()

		if (s.subQueue.Len() > 0 || s.bwdsubAssigns < len(s.d.trail)) && !s.backwardSubsumptionCheck(true) {
			s.d.ok = false
			goto cleanup
		}

		for count := 0; ; count++ {
			elim, ok := s.elimHeap.RemoveMin()
			if !ok {
				break
			}

			if s.d.asyncInterrupt {
				break
			}

			if s.IsEliminated(elim) || !s.d.Value(elim).IsUndef() {
				continue
			}

			if s.d.Verbosity >= 2 && count%1000 == 0 {
				fmt.Fprintf(os.Stderr, "                    elimination remaining: %10d\r", s.elimHeap.Len())
			}

			if s.Asymm {
				// temporary freeze elim to prevent it from ending up on the queue again
				prev := s.frozen[elim]
				s.frozen[elim] = true
				if !s.asymmVar(elim) {
					s.d.ok = false
					goto cleanup
				}
				s.frozen[elim] = prev
			}

			// check if the variable was set by asymmetry branching
			if !s.NoElim && s.d.Value(elim).IsUndef() && !s.frozen[elim] {
				if !s.eliminateVar(elim) {
					s.d.ok = false
					goto cleanup
				}
			}

			s.d.checkGarbageFrac(s.SimpGarbageFrac, false)
		}

		if s.subQueue.Len() != 0 {
			panic("non-empty subsumption queue")
		}
	}

cleanup:

	if turnOffElim {
		s.touched = nil
		s.occurs = nil
		s.numOcc = nil
		s.elimHeap = nil
		s.subQueue = nil
		s.useSimp = false
		s.d.removeSat = true
		s.maxSimpVar = Var(s.d.NumVar())

		s.d.rebuildOrderHeap()
		s.garbageCollect()
	} else {
		s.d.checkGarbage()
	}

	if s.d.Verbosity >= 1 {
		log.Printf("|  Eliminated clauses:   %12d (%10.2f MB)                         |",
			len(s.elimClauses), float64(len(s.elimClauses))*float64(4)/float64(1024*1024))
	}

	return s.d.ok
}

func (s *Simp) extendModel() {
	var j uint32
	var p Lit

nextClause:
	for i := len(s.elimClauses) - 1; i > 0; i -= int(j) {
		for j = s.elimClauses[i]; j > 1; {
			if !s.d.ValueLitModel(Lit(s.elimClauses[i])).IsFalse() {
				continue nextClause
			}
			j--
			i--
		}

		p = Lit(s.elimClauses[i])
		s.d.model[p.Var()] = LiftBool(!p.IsNeg())
	}
}

func (s *Simp) substitute(v Var, p Lit) bool {
	if s.frozen[v] {
		panic("frozen variable cannot be eliminated")
	}
	if s.IsEliminated(v) {
		panic("variable already eliminated")
	}
	if !s.d.Value(v).IsUndef() {
		panic("variable has a value")
	}

	s.eliminated[v] = true
	s.d.SetDecision(v, false)

	cs := s.occurs.Lookup(v)
	for _, c := range cs {
		var sc []Lit

		for j := range c.Lit {
			q := c.Lit[j]
			if q.Var() == v {
				sc = append(sc, p.Xor(q.IsNeg()))
			} else {
				sc = append(sc, p)
			}
		}

		s.removeClause(c)
		if !s.AddClause(sc...) {
			s.d.ok = false
			return false
		}
	}

	return true
}

func (s *Simp) eliminateVar(v Var) bool {
	if s.frozen[v] {
		panic("frozen variable cannot be eliminated")
	}
	if s.IsEliminated(v) {
		panic("variable already eliminated")
	}
	if !s.d.Value(v).IsUndef() {
		panic("variable has a value")
	}

	cs := s.occurs.Lookup(v)
	var pos, neg []*Clause
	for i := range cs {
		for j := range cs[i].Lit {
			if cs[i].Lit[j].Var() == v {
				if cs[i].Lit[j].IsNeg() {
					neg = append(neg, cs[i])
				} else {
					pos = append(pos, cs[i])
				}
			}
		}
	}

	var ok bool
	count := 0
	clauseSize := 0

	for i := range pos {
		for j := range neg {
			ok, clauseSize = s.mergeSize(pos[i], neg[j], v)
			if ok {
				count++
				if count > len(cs)+s.Grow || (s.ClauseLimit != -1 && clauseSize > s.ClauseLimit) {
					return true
				}
			}
		}
	}

	s.eliminated[v] = true
	s.d.SetDecision(v, false)
	s.nelimvars++

	if len(pos) > len(neg) {
		for i := range neg {
			s.mkElimClauseFrom(v, neg[i])
		}
		s.mkElimClause(Literal(v, false))
	} else {
		for i := range pos {
			s.mkElimClauseFrom(v, pos[i])
		}
		s.mkElimClause(Literal(v, true))
	}

	for i := range cs {
		s.removeClause(cs[i])
	}

	for i := range pos {
		for j := range neg {
			ok, psResolvent := s.merge(pos[i], neg[j], v)
			if ok && !s.AddClause(psResolvent...) {
				return false
			}
		}
	}

	s.occurs.RemoveAll(v, true)
	if len(s.d.watches.Occurrences(Literal(v, false))) == 0 {
		s.d.watches.RemoveAll(Literal(v, false), true)
	}
	if len(s.d.watches.Occurrences(Literal(v, true))) == 0 {
		s.d.watches.RemoveAll(Literal(v, true), true)
	}

	return s.backwardSubsumptionCheck(false)
}

func (s *Simp) mkElimClause(p Lit) {
	s.elimClauses = append(s.elimClauses, uint32(p))
	s.elimClauses = append(s.elimClauses, 1) // undefined literal
}

func (s *Simp) mkElimClauseFrom(v Var, c *Clause) {
	first := len(s.elimClauses)
	vpos := -1

	for i := range c.Lit {
		s.elimClauses = append(s.elimClauses, uint32(c.Lit[i]))
		if c.Lit[i].Var() == v {
			vpos = i + first
		}
	}
	if vpos == -1 {
		panic("variable not found")
	}

	// swap the first literal with the v literal
	s.elimClauses[first], s.elimClauses[vpos] = s.elimClauses[vpos], s.elimClauses[first]

	s.elimClauses = append(s.elimClauses, uint32(c.Len()))
}

func (s *Simp) asymmVar(v Var) bool {
	s.assertSimp()
	cs := s.occurs.Lookup(v)
	if !s.d.Value(v).IsUndef() || len(cs) == 0 {
		return true
	}
	for i := range cs {
		if !s.asymm(v, cs[i]) {
			return false
		}
	}

	return s.backwardSubsumptionCheck(false)
}

func (s *Simp) asymm(v Var, c *Clause) bool {
	if s.d.decisionLevel() != 0 {
		panic("non-root decision level")
	}

	if c.Mark != 0 || s.d.satisfied(c) {
		return true
	}

	s.d.newDecisionLevel()
	var p Lit
	for i := range c.Lit {
		if c.Lit[i].Var() != v && !s.d.ValueLit(c.Lit[i]).IsFalse() {
			s.d.uncheckedEnqueue(c.Lit[i].Inverse(), nil)
		} else {
			p = c.Lit[i]
		}
	}

	if s.d.propagate() != nil {
		s.d.cancelUntil(0)
		s.nasymmlit++
		if !s.strengthenClause(c, p) {
			return false
		}
	} else {
		s.d.cancelUntil(0)
	}

	return true
}

func (s *Simp) backwardSubsumptionCheck(verbose bool) bool {
	if s.d.decisionLevel() != 0 {
		panic("non-root decision level")
	}

	count := 0
	numSubsumed := 0
	numDeletedLiterals := 0
	for s.subQueue.Len() > 0 || s.bwdsubAssigns < len(s.d.trail) {
		// user interrupt -- empty subsumption queue and return immediately
		if s.d.asyncInterrupt {
			s.subQueue.clear()
			s.bwdsubAssigns = len(s.d.trail)
			break
		}

		if s.subQueue.Len() == 0 && s.bwdsubAssigns < len(s.d.trail) {
			p := s.d.trail[s.bwdsubAssigns]
			s.bwdsubAssigns++
			s.subQueue.Insert(s.d.newClause([]Lit{p}, false))
		}

		c := s.subQueue.Pop()
		if c.Mark != 0 {
			continue
		}

		if verbose && s.d.Verbosity >= 2 {
			if count%1000 == 0 {
				fmt.Fprintf(os.Stderr, "                    subsumption left: %10d (%10d subsumed, %10d deleted literals)\r", s.subQueue.Len(), numSubsumed, numDeletedLiterals)
			}
			count++
		}

		if c.Len() == 1 && !s.d.ValueLit(c.Lit[0]).IsTrue() {
			// unit clauses should have been propagated at this point
			panic("unsatisfied unit clause")
		}

		// find best variable to scan
		best := c.Lit[0].Var()
		bestSize := len(s.occurs.Occurrences(c.Lit[0].Var()))
		for i := 1; i < c.Len(); i++ {
			if len(s.occurs.Occurrences(c.Lit[i].Var())) < bestSize {
				best = c.Lit[i].Var()
				bestSize = len(s.occurs.Occurrences(c.Lit[i].Var()))
			}
		}

		cs := s.occurs.Lookup(best)
		for j := 0; j < len(cs); j++ {
			if c.Mark != 0 {
				break
			}
			if cs[j].Mark == 0 && cs[j] != c && (s.SubsumptionLimit == -1 || cs[j].Len() < s.SubsumptionLimit) {
				ok, p := c.Subsumes(cs[j])
				if !ok {
					continue
				}

				if p.IsUndef() {
					numSubsumed++
					s.removeClause(cs[j])
				} else {
					numDeletedLiterals++

					if !s.strengthenClause(cs[j], p.Inverse()) {
						return false
					}

					if p.Var() == best {
						j--
					}
				}
			}
		}
	}
	return true
}

func (s *Simp) implied(ps []Lit) bool {
	if s.d.decisionLevel() != 0 {
		panic("non-root decision level")
	}
	s.d.newDecisionLevel()
	for i := range ps {
		if val := s.d.ValueLit(ps[i]); val.IsTrue() {
			s.d.cancelUntil(0)
			return true
		} else if !val.IsFalse() {
			if val != LUndef {
				panic("unexpected value")
			}
			s.d.uncheckedEnqueue(ps[i].Inverse(), nil)
		}
	}

	ok := s.d.propagate() != nil
	s.d.cancelUntil(0)
	return ok
}

// FIXME:
// this whole function feels inefficient -- double marking stuff already
// present in the queue... it is strange.
func (s *Simp) gatherTouchedClauses() {
	if s.numTouched == 0 {
		return
	}

	qlen := s.subQueue.Len()
	for i := 0; i < qlen; i++ {
		c := s.subQueue.Clause(i)
		if c.Mark == 0 {
			c.Mark = MarkTouch
		}
	}

	for i := 1; i < s.d.NumVar(); i++ {
		if s.touched[i] {
			cs := s.occurs.Lookup(Var(i))
			for j := range cs {
				if cs[j].Mark == 0 {
					s.subQueue.Insert(cs[j])
					cs[j].Mark = MarkTouch
				}
			}
			s.touched[i] = false
		}
	}

	qlen = s.subQueue.Len()
	for i := 0; i < qlen; i++ {
		c := s.subQueue.Clause(i)
		if c.Mark == MarkTouch {
			c.Mark = 0
		}
	}

	s.numTouched = 0
}

// merge returns clause literals from merging c1 and c2 ond v.  If merge
// returns false then "clause" (?) is always satisfied and ps should not be
// used.
func (s *Simp) merge(c1 *Clause, c2 *Clause, v Var) (ok bool, ps []Lit) {
	s.nmerge++
	if c1.Len() < c2.Len() {
		c1, c2 = c2, c1
	}

	// NOTE:
	// use i as an index to c1.Lit and j as an index into c2.Lit or things will
	// be confusing.

shortList:
	for j := range c2.Lit {
		if c2.Lit[j].Var() != v {
			vq := c2.Lit[j].Var()
			for i := range c1.Lit {
				if c1.Lit[i].Var() == vq {
					if c1.Lit[i] == c2.Lit[j].Inverse() {
						return false, nil
					}
					continue shortList
				}
			}
			ps = append(ps, c2.Lit[j])
		}
	}

	for i := range c1.Lit {
		if c1.Lit[i].Var() != v {
			ps = append(ps, c1.Lit[i])
		}
	}

	return true, ps
}

// merge returns the number of clause literals that result from merging c1 and
// c2 ond v.  If mergeSize returns false then "clause" (?) is always satisfied
// and size should not be used.
func (s *Simp) mergeSize(c1 *Clause, c2 *Clause, v Var) (ok bool, size int) {
	s.nmerge++
	if c1.Len() < c2.Len() {
		c1, c2 = c2, c1
	}

	// NOTE:
	// use i as an index to c1.Lit and j as an index into c2.Lit or things will
	// be confusing.

	size = c1.Len() - 1

shortList:
	for j := range c2.Lit {
		if c2.Lit[j].Var() != v {
			vq := c2.Lit[j].Var()
			for i := range c1.Lit {
				if c1.Lit[i].Var() == vq {
					if c1.Lit[i] == c2.Lit[j].Inverse() {
						return false, 0
					}
					continue shortList
				}
			}
			size++
		}
	}

	return true, size
}

func (s *Simp) strengthenClause(c *Clause, p Lit) bool {
	if s.d.decisionLevel() != 0 {
		panic("non-root decision level")
	}
	s.assertSimp()

	// TODO: if !s.subQueue.Contains(c) then insert
	s.subQueue.Insert(c)

	if c.Len() == 2 {
		s.removeClause(c)
		if !c.Strengthen(p) {
			panic("could not strengthen")
		}
	} else {
		s.d.detachClause(c, true)
		if !c.Strengthen(p) {
			panic("could not strengthen")
		}
		s.d.attachClause(c)
		// FIXME: this is going to be balls slow potentially lots of clauses
		s.occurs.Remove(p.Var(), c)
		s.numOcc[p]--
		s.updateElimHeap(p.Var())
	}

	if c.Len() == 1 {
		if !s.d.enqueue(c.Lit[0], nil) {
			return false
		}
		return s.d.propagate() == nil
	}
	return true
}

func (s *Simp) assertSimp() {
	if !s.useSimp {
		panic("simplification disabled")
	}
}

func (s *Simp) garbageCollect() {
	s.relocAll()
	s.d.garbageCollect()
}

// for now not much to do in here... if relocation really turns into something
// then there will be something to do.
func (s *Simp) relocAll() {
	if s.occurs != nil {
		s.occurs.CleanAll()
	}
}
