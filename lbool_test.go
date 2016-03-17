// Copyright 2016 Bryan Matsuo
//
// Use of this software is governed by the MIT license.  A copy of the license
// agreement can be found in the LICENSE file distributed with this software.

package dpll

import "testing"

func TestLBool_Equal(t *testing.T) {
	tests := []struct {
		a, b LBool
		eq   bool
	}{
		{LTrue, LTrue, true},
		{LFalse, LFalse, true},
		{LUndef, LUndef, true},
		{LTrue, LUndef, false},
		{LTrue, LFalse, false},
		{LFalse, LUndef, false},
	}

	for i, test := range tests {
		if test.a.Equal(test.b) && !test.eq {
			t.Errorf("test %d: %v == %v (%v != %v)", i, test.a, test.b, test.a, test.b)
		}
		if !test.a.Equal(test.b) && test.eq {
			t.Errorf("test %d: %v != %v (%v == %v)", i, test.a, test.b, test.a, test.b)
		}
		if test.b.Equal(test.a) && !test.eq {
			t.Errorf("test %d: %v == %v (%v != %v)", i, test.b, test.a, test.b, test.a)
		}
		if !test.b.Equal(test.a) && test.eq {
			t.Errorf("test %d: %v != %v (%v == %v)", i, test.b, test.a, test.b, test.a)
		}
	}
}

func TestLBool(t *testing.T) {
	tests := []struct {
		b       LBool
		isUndef bool
		isTrue  bool
		isFalse bool
	}{
		0:  {LUndef, true, false, false},
		1:  {LTrue, false, true, false},
		2:  {LFalse, false, false, true},
		3:  {LUndef.Or(LUndef), true, false, false},
		4:  {LUndef.Or(LTrue), false, true, false},
		5:  {LUndef.Or(LFalse), true, false, false},
		7:  {LTrue.Or(LUndef), false, true, false},
		6:  {LTrue.Or(LTrue), false, true, false},
		8:  {LTrue.Or(LFalse), false, true, false},
		9:  {LFalse.Or(LUndef), true, false, false},
		10: {LFalse.Or(LTrue), false, true, false},
		11: {LFalse.Or(LFalse), false, false, true},
		12: {LUndef.And(LUndef), true, false, false},
		13: {LUndef.And(LTrue), true, false, false},
		14: {LUndef.And(LFalse), false, false, true},
		15: {LTrue.And(LUndef), true, false, false},
		16: {LTrue.And(LTrue), false, true, false},
		17: {LTrue.And(LFalse), false, false, true},
		18: {LFalse.And(LUndef), false, false, true},
		19: {LFalse.And(LTrue), false, false, true},
		20: {LFalse.And(LFalse), false, false, true},
		21: {LUndef.Xor(true), true, false, false},
		22: {LUndef.Xor(false), true, false, false},
		23: {LTrue.Xor(true), false, false, true},
		24: {LTrue.Xor(false), false, true, false},
		25: {LFalse.Xor(true), false, true, false},
		26: {LFalse.Xor(false), false, false, true},
	}

	for i, test := range tests {
		if test.b.IsUndef() != test.isUndef {
			t.Errorf("test %d: is Undef %v (!= %v)", i, test.b.IsUndef(), test.isUndef)
		}
		if test.b.IsTrue() != test.isTrue {
			t.Errorf("test %d: is True %v (!= %v)", i, test.b.IsTrue(), test.isTrue)
		}
		if test.b.IsFalse() != test.isFalse {
			t.Errorf("test %d: is False %v (!= %v)", i, test.b.IsFalse(), test.isFalse)
		}
	}
}
