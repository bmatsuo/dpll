#dpll

DPLL satisfiability solver based on [MiniSat](https://github.com/niklasso/minisat).

The goal of project is to provide a basic and somewhat efficient SAT solving
package that can be used for constraint solving in practical Go applications.
A command line program is provided though it should not be expected to
outperform MiniSat.

Like MiniSat the dpll package can be thought of as a core implementation
suitable for hacking on.  Updates to the dpll package itself should be limited
to bug fixes and performance tweaks.

##License (MIT)

Copyright 2016 Bryan Matsuo

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to conditions specified in the [LICENSE](LICENSE) file distributed
with this software.
