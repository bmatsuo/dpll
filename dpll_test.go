// Copyright 2016 Bryan Matsuo
//
// Use of this software is governed by the MIT license.  A copy of the license
// agreement can be found in the LICENSE file distributed with this software.

package dpll

import "testing"

func TestSolver_Solve_sat_factoring_3_5(t *testing.T) {
	d := New(nil)
	err := DecodeFile(d, "testdata/factoring_3_5.cnf")
	if err != nil {
		t.Fatal(err)
	}
	sat := d.Solve()
	if !sat {
		t.Errorf("not satisfiable")
	}
}

func TestSolver_Solve_unsat_factoring_3_5(t *testing.T) {
	d := New(nil)
	err := DecodeFile(d, "testdata/factoring_3_5_UNSAT.cnf")
	if err != nil {
		t.Fatal(err)
	}
	sat := d.Solve()
	if !sat {
		t.Errorf("not satisfiable")
	}
}

func TestSolver_Solve_sat_factoring_5_7(t *testing.T) {
	d := New(nil)
	err := DecodeFile(d, "testdata/factoring_5_7.cnf")
	if err != nil {
		t.Fatal(err)
	}
	sat := d.Solve()
	if !sat {
		t.Errorf("not satisfiable")
	}
}
