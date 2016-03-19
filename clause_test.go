// Copyright 2016 Bryan Matsuo
//
// Use of this software is governed by the MIT license.  A copy of the license
// agreement can be found in the LICENSE file distributed with this software.

package dpll

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestClause(t *testing.T) {
	l1 := []Lit{
		Literal(1, false),
		Literal(2, true),
	}
	c1 := NewClause(l1, false, false)
	if !reflect.DeepEqual(c1.Lit, l1) {
		t.Errorf("lit: %v (!= %v)", c1.Lit, l1)
	}
	if c1.ClauseExtra != nil {
		t.Errorf("extra: %v", c1.ClauseExtra)
	}
	c2 := NewClause(l1, true, false)
	if c2.ClauseExtra == nil {
		t.Fatalf("extra: %v", c2.ClauseExtra)
	}
	if c2.Abstraction == 0 {
		t.Fatalf("abstraction: %v", c2.Abstraction)
	}
	c3 := NewClause(l1, true, true)
	if c3.ClauseExtra == nil {
		t.Fatalf("extra: %v", c3.ClauseExtra)
	}
	if c3.Abstraction != 0 {
		t.Fatalf("abstraction: %v", c3.Abstraction)
	}

	c4 := NewClauseFrom(c2, true)
	if !reflect.DeepEqual(c4.Lit, c2.Lit) {
		t.Errorf("lit: %v (!= %v)", c4.Lit, c2.Lit)
	}
	if !reflect.DeepEqual(c4.ClauseExtra, c2.ClauseExtra) {
		t.Fatalf("extra: %v (!= %v)", c4.ClauseExtra, c2.ClauseExtra)
	}
	if c4.ClauseExtra == c2.ClauseExtra {
		// the ClauseExtra should not actually share memory
		t.Fatalf("extra: %x == %x)", unsafe.Pointer(c4.ClauseExtra), unsafe.Pointer(c2.ClauseExtra))
	}

	c5 := NewClauseFrom(c2, false)
	if c5.ClauseExtra != nil {
		t.Fatalf("extra: %v (!= nil)", c5.ClauseExtra)
	}
}

func TestClause_Subsumes(t *testing.T) {
	// TODO: Subsumes is not used in the core DPLL solver so it doesn't need a test yet
}

func TestClause_Strengthen(t *testing.T) {
	// TODO: Strengthen is not used in the core DPLL solver so it doesn't need a test yet
}
