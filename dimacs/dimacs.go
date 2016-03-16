// Copyright 2016 Bryan Matsuo
//
// Use of this software is governed by the MIT license.  A copy of the license
// agreement can be found in the LICENSE file distributed with this software.

// Package dimacs implements reading and writing of DIMACS format files for
// satisfiability problem statement.
package dimacs

import (
	"io"
	"os"
)

// Header precedes clause data in a DIMACS data stream
type Header struct {
	NumVar    int
	NumClause int
}

// Problem is the statement of a SAT problem in CNF.
type Problem struct {
	NumVar  int
	Clauses [][]Lit
}

// DecodeFile opens path and decodes its contents using DecodeProblem.
func DecodeFile(path string) (*Problem, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return DecodeProblem(f)
}

// DecodeProblem decodes the contents of r into a new Problem.
func DecodeProblem(r io.Reader) (*Problem, error) {
	d := NewDecoder(r)
	h := d.Header()
	if d.Err() != nil {
		return nil, d.Err()
	}
	p := &Problem{}
	p.NumVar = h.NumVar
	p.Clauses = make([][]Lit, 0, h.NumClause)
	for d.Decode() {
		p.newClause(d.Clause())
	}
	if d.Err() != nil {
		return nil, d.Err()
	}
	return p, nil
}

// EncodeFile encodes p in DIMACS format and writes the resulting bytes to a
// new file at path.
func EncodeFile(path string, p *Problem) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return EncodeProblem(f, p)
}

// EncodeProblem encodes p in DIMACS format and writes the resulting bytes to
// w.
func EncodeProblem(w io.Writer, p *Problem) error {
	enc := NewEncoder(w)
	err := enc.WriteHeader(&Header{
		NumVar:    p.NumVar,
		NumClause: len(p.Clauses),
	})
	if err != nil {
		return err
	}
	for _, clause := range p.Clauses {
		err = enc.Encode(clause)
		if err != nil {
			return err
		}
	}
	return enc.Close()
}

func (p *Problem) newClause(c []Lit) {
	_c := make([]Lit, len(c))
	copy(_c, c)
	p.Clauses = append(p.Clauses, _c)
}

// Lit is a simple representation of a clause literal.  Positive literals are
// equal to the variable they contain.  Negated literals are equal to the
// negative value (additive inverse) of the variable they contain.  The value 0
// is not a valid Lit.
type Lit int

// Var returns the variable of lit.  A valid literal has a positive variable.
func (lit Lit) Var() int {
	if lit < 0 {
		return int(-lit)
	}
	return int(lit)
}

// Neg returns true if the literal is negated (lit == -p for some variable p).
func (lit Lit) Neg() bool {
	return lit < 0
}
