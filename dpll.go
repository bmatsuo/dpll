// Copyright 2016 Bryan Matsuo
//
// Use of this software is governed by the MIT license.  A copy of the license
// agreement can be found in the LICENSE file distributed with this software.

/*
Package dpll provides a solver for problems propositional satisfiability
(SAT) using the DPLL search algorithm.  The specific search implementation
used is based on the MiniSat program.

The DPLL solver included runs single-threaded because the DPLL algorithm does
not extend simply into multiple threads.  Furthermore the memory requirements
for multithreaded solving can become extremely high.

Forked packages may be modified to allow for multi-threaded solving using the
DPLL type under the hood.  Alternatively, a simple way to achieve
multi-threaded solving is to run multiple single-threaded solvers (potenially
identical solvers with different random seeds) and wait for the first to
finish.
*/
package dpll
