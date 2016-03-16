// Copyright 2016 Bryan Matsuo
//
// Use of this software is governed by the MIT license.  A copy of the license
// agreement can be found in the LICENSE file distributed with this software.

package dpll

import "testing"

func TestLit(t *testing.T) {
	tests := []struct {
		v     Var
		pos   bool
		undef bool
		s     string
	}{
		{Var(1), true, false, "1"},
		{Var(1), false, false, "-1"},
		{Var(0), true, true, "0"},
		{Var(0), false, true, "0"},
	}

	for i, test := range tests {
		lit := NewLit(test.v, test.pos)
		if lit.Var() != test.v {
			t.Errorf("test %d: variable %d (!= %d)", i, lit.Var(), test.v)
		}
		if lit.IsPos() != test.pos {
			t.Errorf("test %d: positive %v (!= %v)", i, lit.IsPos(), test.pos)
		}
		if lit.IsUndef() != test.undef {
			t.Errorf("test %d: undef %v (!= %v)", i, lit.IsUndef(), test.undef)
		}
		if lit.String() != test.s {
			t.Errorf("test %d: %q (!= %q)", i, lit.String(), test.s)
		}
	}
}
