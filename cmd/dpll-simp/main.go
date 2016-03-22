package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/bmatsuo/dpll"
)

func init() {
	runtime.GOMAXPROCS(1)
}

func main() {
	cpuProfile := flag.String("cpuprofile", "", "path to write a pprof cpu profile for execution")
	attemptSolve := flag.Bool("solve", true, "do not attempt to solve the problem")
	verbosity := flag.Int("v", 1, "verbosity level")
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatalf("%s expects exactly one argument", os.Args[0])
	}

	exitCode := 0
	const (
		SAT   = 10
		UNSAT = 20
		FAIL  = 1
	)
	defer func() {
		if e := recover(); e != nil {
			panic(e)
		}

		switch exitCode {
		case 10:
			fmt.Println("SATISFIABLE")
		case 20:
			fmt.Println("UNSATISFIABLE")
		default:
			fmt.Println("INDETERMINATE")
		}

		os.Exit(exitCode)
	}()

	if *cpuProfile != "" {
		fcpu, err := os.Create(*cpuProfile)
		if err != nil {
			log.Printf("failed creating pprof file: %v", err)
			exitCode = FAIL
			return
		}
		defer fcpu.Close()

		err = pprof.StartCPUProfile(fcpu)
		if err != nil {
			log.Printf("failed to start profiling: %v", err)
			exitCode = FAIL
			return
		}
		defer pprof.StopCPUProfile()
	}

	solver := dpll.NewSimp(&dpll.Opt{
		Verbosity: *verbosity,
	}, nil)

	parseStart := time.Now()
	_, err := dpll.DecodeFile(solver, flag.Arg(0))
	if err != nil {
		log.Print(err)
		return
	}
	parseEnd := time.Now()

	log.Printf("============================[ Problem Statistics ]=============================")
	if *verbosity >= 1 {
		log.Printf("|  Number of variables:  %12d                                         |", solver.NumVar())
		log.Printf("|  Number of clauses:    %12d                                         |", solver.NumClause())
		dur := parseEnd.Sub(parseStart)
		log.Printf("|  Parse time:           %12v                                         |", dur-dur%time.Microsecond)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func(c chan os.Signal) {
		for range c {
			signal.Stop(c)
			solver.Interrupt()
			break
		}
	}(sig)

	solver.Eliminate(true)
	simplifyEnd := time.Now()
	if *verbosity >= 1 {
		dur := simplifyEnd.Sub(parseEnd)
		log.Printf("|  Simplify time:        %12v                                         |", dur-dur%time.Microsecond)
		log.Printf("|                                                                             |")
	}

	if !solver.Okay() {
		if *verbosity >= 1 {
			log.Printf("===============================================================================")
			log.Printf("Solved by simplification")
			solver.PrintStats()
			log.Println()
		}
		fmt.Fprintln(os.Stderr)
		exitCode = UNSAT
		return
	}

	solution := dpll.LUndef
	if *attemptSolve {
		solution = solver.SolveLimited()
	} else {
		if *verbosity >= 1 {
			log.Printf("===============================================================================")
			log.Printf("Simplification did not yield a result.")
			log.Println()
			log.Printf("No solution attempted.")
			log.Println()
		}
	}

	if *verbosity >= 1 {
		solver.PrintStats()
		fmt.Fprintln(os.Stderr)
	}
	if solution.IsTrue() {
		exitCode = SAT
		return
	} else if solution.IsFalse() {
		exitCode = UNSAT
		return
	} else {
		exitCode = 0
		return
	}
}
