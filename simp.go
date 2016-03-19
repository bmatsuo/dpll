// Copyright 2016 Bryan Matsuo
//
// Use of this software is governed by the MIT license.  A copy of the license
// agreement can be found in the LICENSE file distributed with this software.

package dpll

// Simp is a solver that simplifies clauses before solving.
type Simp struct {
	grow            int  // Allow a variable elimination step to grow by a number of clauses (default to zero).
	clauseLimit     int  // Variables are not eliminated if it produces a resolvent with a length above this limit.
	subsumpLimit    int  // Do not check if subsumption against a clause larger than this.
	simpGarbageFrac int  // A different limit for when to issue a GC during simplification (see garbageFrac).
	useAsymm        bool // Shrink clauses by asymmetric branching
	useRCheck       bool // Check if a clause is implied. Costly, and subsumes subsumtions
	useElim         bool // Perform variable elimination
	extendModel     bool // Flag to indicate whether the user needs to look at the full model

	nmerge    int
	nasymmlit int
	nelimvars int

	// solver state
	elimOrder     int
	useSimp       bool
	maxSimpVar    Var
	elimClauses   uint32
	touched       uint8
	occurs        *clauseOccLists
	numOcc        map[Lit]int
	elimHeap      elimQueue
	subQueue      *clauseQueue
	frozen        []uint8
	frozenVars    []Var
	eliminated    []uint8
	bwdsubAssigns int
	numTouched    int

	DPLL
}
