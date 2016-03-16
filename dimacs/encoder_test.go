// Copyright 2016 Bryan Matsuo
//
// Use of this software is governed by the MIT license.  A copy of the license
// agreement can be found in the LICENSE file distributed with this software.

package dimacs

import (
	"bytes"
	"testing"
)

func TestEncoder(t *testing.T) {
	tests := []struct {
		out string
		p   *Problem
	}{
		{
			"p cnf 3 0\n",
			&Problem{3, nil},
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
	}

	for i, test := range tests {
		var buf bytes.Buffer
		err := EncodeProblem(&buf, test.p)
		if err != nil {
			t.Errorf("test %d: %v", i, err)
			continue
		}
		text := buf.String()
		if text != test.out {
			t.Errorf("test %d: output %q (!= %q)", i, text, test.out)
		}
	}
}
