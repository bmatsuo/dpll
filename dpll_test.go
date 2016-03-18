// Copyright 2016 Bryan Matsuo
//
// Use of this software is governed by the MIT license.  A copy of the license
// agreement can be found in the LICENSE file distributed with this software.

package dpll

import "testing"

func TestSolver_Solve_sat_factoring_2_3(t *testing.T) {
	d := New(nil)
	ok, err := DecodeFile(d, "testdata/factoring_2_3.cnf")
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Errorf("not ok")
	}
	sat := d.Solve()
	if !sat {
		t.Errorf("not satisfiable")
	}
	model := d.Model()
	if len(model) != d.NumVar()+1 {
		t.Errorf("model length: %d (!= %d)", len(model), d.NumVar()+1)
	}
}

func TestSolver_Solve_unsat(t *testing.T) {
	d := New(nil)
	_, err := DecodeFile(d, "testdata/unsat.cnf")
	if err != nil {
		t.Fatal(err)
	}
	sat := d.Solve()
	t.Logf("vars: %d clauses: %d", d.NumVar(), d.NumClause())
	if sat {
		model := d.Model()
		t.Errorf("satisfiable: %v", model)
		for v := 1; v < len(model); v++ {
			t.Errorf("Var(%d) == %v", v, model[v])
		}
	}
	conflict := d.Conflict()
	if len(conflict) == 0 {
		t.Errorf("empty conflict")
	}
}

func TestSolver_Solve_unsat_factoring_2_3(t *testing.T) {
	d := New(nil)
	_, err := DecodeFile(d, "testdata/factoring_2_3_UNSAT.cnf")
	if err != nil {
		t.Fatal(err)
	}
	sat := d.Solve()
	t.Logf("vars: %d clauses: %d", d.NumVar(), d.NumClause())
	if sat {
		model := d.Model()
		t.Errorf("satisfiable: %v", model)
		for v := 1; v < len(model); v++ {
			t.Errorf("Var(%d) == %v", v, model[v])
		}
	}
	conflict := d.Conflict()
	if len(conflict) == 0 {
		t.Errorf("empty conflict")
	}
}

func TestSolver_Solve_sat_factoring_3_5(t *testing.T) {
	d := New(nil)
	ok, err := DecodeFile(d, "testdata/factoring_3_5.cnf")
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Errorf("not ok")
	}
	sat := d.Solve()
	if !sat {
		t.Errorf("not satisfiable")
	}
	model := d.Model()
	if len(model) != d.NumVar()+1 {
		t.Errorf("model length: %d (!= %d)", len(model), d.NumVar()+1)
	}
}

func TestSolver_Solve_unsat_factoring_3_5(t *testing.T) {
	d := New(nil)
	_, err := DecodeFile(d, "testdata/factoring_3_5_UNSAT.cnf")
	if err != nil {
		t.Fatal(err)
	}
	sat := d.Solve()
	t.Logf("vars: %d clauses: %d", d.NumVar(), d.NumClause())
	if sat {
		model := d.Model()
		t.Errorf("satisfiable: %v", model)
		for v := 1; v < len(model); v++ {
			t.Errorf("Var(%d) == %v", v, model[v])
		}
	}
	conflict := d.Conflict()
	if len(conflict) == 0 {
		t.Errorf("empty conflict")
	}
}

func TestSolver_Solve_sat_factoring_5_7(t *testing.T) {
	d := New(nil)
	ok, err := DecodeFile(d, "testdata/factoring_5_7.cnf")
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Errorf("not ok")
	}
	sat := d.Solve()
	if !sat {
		t.Errorf("not satisfiable")
	}
	model := d.Model()
	if len(model) != d.NumVar()+1 {
		t.Errorf("model length: %d (!= %d)", len(model), d.NumVar()+1)
	}
}
