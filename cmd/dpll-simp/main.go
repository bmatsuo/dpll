package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/bmatsuo/dpll"
)

func main() {
	runtime.GOMAXPROCS(1)
	verbosity := flag.Int("v", 1, "verbosity level")
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatalf("%s expects exactly one argument", os.Args[0])
	}
	solver := dpll.NewSimp(&dpll.Opt{
		Verbosity: *verbosity,
	}, nil)

	parseStart := time.Now()
	_, err := dpll.DecodeFile(solver, flag.Arg(0))
	if err != nil {
		log.Fatal(err)
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
		fmt.Println("UNSATISFIABLE")
		os.Exit(20)
	}

	// TODO: if simplification is all that is desired then exit as indeterminate result
	solution := solver.SolveLimited()

	if *verbosity >= 1 {
		solver.PrintStats()
		fmt.Fprintln(os.Stderr)
	}
	if solution.IsTrue() {
		fmt.Println("SATISFIABLE")
		os.Exit(10)
	} else if solution.IsFalse() {
		fmt.Println("UNSATISFIABLE")
		os.Exit(20)
	} else {
		fmt.Println("INDETERMINATE")
		os.Exit(0)
	}
}
