// Copyright 2016 Bryan Matsuo
//
// Use of this software is governed by the MIT license.  A copy of the license
// agreement can be found in the LICENSE file distributed with this software.

package dimacs

import (
	"reflect"
	"strings"
	"testing"
)

func TestDecoder(t *testing.T) {
	tests := []struct {
		in string
		p  *Problem
	}{
		{
			"p cnf 3 0\n",
			&Problem{3, [][]Lit{}},
		},
		{
			"p cnf 3 1\n0\n",
			&Problem{3, [][]Lit{
				{},
			}},
		},
		{
			"p cnf 3 1\n-1 0\n",
			&Problem{3, [][]Lit{
				{-1},
			}},
		},
		{
			"p cnf 3 1\nc a comment\n-1 0\n",
			&Problem{3, [][]Lit{
				{-1},
			}},
		},
	}

	for i, test := range tests {
		p, err := DecodeProblem(strings.NewReader(test.in))
		if err != nil {
			t.Errorf("test %d: %v", i, err)
			continue
		}
		if !reflect.DeepEqual(p, test.p) {
			t.Errorf("test %d: p %v (!= %v)", i, p, test.p)
		}
	}
}
