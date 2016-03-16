// Copyright 2016 Bryan Matsuo
//
// Use of this software is governed by the MIT license.  A copy of the license
// agreement can be found in the LICENSE file distributed with this software.

package dimacs

import "testing"

func TestLit(t *testing.T) {
	tests := []struct {
		lit Lit
		neg bool
		v   int
	}{
		{0, false, 0}, // an invalid literal and variable
		{-1, true, 1},
		{1, false, 1},
	}

	for i, test := range tests {
		if test.lit.Var() != test.v {
			t.Errorf("test %d: variable %d (!= %d)", i, test.lit.Var(), test.v)
		}
		if test.lit.Neg() != test.neg {
			t.Errorf("test %d: neg %v (!= %v)", i, test.lit.Neg(), test.neg)
		}
	}
}
