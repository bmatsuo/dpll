// Copyright 2016 Bryan Matsuo
//
// Use of this software is governed by the MIT license.  A copy of the license
// agreement can be found in the LICENSE file distributed with this software.

package dpll

import (
	"io"
	"os"

	"github.com/bmatsuo/dpll/dimacs"
)

// Solver is a simple interface for dpll solvers.  It is not comprehensive but
// it is adequate for simple tasks like aggregation/multi-solving and
// deserialization.
type Solver interface {
	NumVar() int
	NumClause() int
	NewVar(upol LBool, dvar bool) Var
	AddClause(p ...Lit) bool
	Interrupt()
	ClearInterrupt()
	Solve(assumptions ...Lit) bool
	SolveLimited(assumptions ...Lit) LBool
}

// DecodeFile decodes a CNF problem in DIMACS format from a file at the given
// path and adds the contained clauses into s.
//
//		solver := dpll.New(nil)
//		err := DecodeFile(solver, "problem.dimacs")
//		if err != nil {
//			// ...
//		}
func DecodeFile(s Solver, path string) (ok bool, err error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()
	return Decode(s, f)
}

// Decode is like DecodeFile. But, Decode reads a DIMACS formatted byte stream
// from r.
func Decode(s Solver, r io.Reader) (ok bool, err error) {
	dec := dimacs.NewDecoder(r)
	for dec.Decode() {
		dc := dec.Clause()
		ls := make([]Lit, len(dc))
		for i, dl := range dc {
			if dl < 0 {
				ls[i] = Literal(Var(-dl), true)
			} else {
				ls[i] = Literal(Var(dl), false)
			}
		}
		for _, p := range ls {
			for p.Var() > Var(s.NumVar()) {
				s.NewVar(LUndef, true)
				//log.Printf("Var(%d)", v)
			}
		}
		if !s.AddClause(ls...) {
			// a contradiction in the clauses was found
			return false, nil
		}
		//log.Printf("AddClause(%v) => %d", dc, s.NumClause())
	}
	if dec.Err() != nil {
		return false, dec.Err()
	}
	return true, nil
}
