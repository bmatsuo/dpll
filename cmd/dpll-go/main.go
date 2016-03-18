package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bmatsuo/dpll"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatalf("%s expects exactly one argument", os.Args[0])
	}
	d := dpll.New(&dpll.Opt{
		Verbosity: 1,
	})
	parseStart := time.Now()
	_, err := dpll.DecodeFile(d, flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	parseEnd := time.Now()
	log.Printf("============================[ Problem Statistics ]=============================")
	if d.Verbosity >= 1 {
		log.Printf("|  Number of variables:  %12d                                         |", d.NumVar())
		log.Printf("|  Number of clauses:    %12d                                         |", d.NumClause())
		dur := parseEnd.Sub(parseStart)
		log.Printf("|  Parse time:           %12v                                         |", dur-dur%time.Microsecond)
	}
	if !d.Simplify() {
		// TODO: handle output for non-tty outputs
		if d.Verbosity >= 1 {
			log.Printf("===============================================================================")
			log.Printf("Solved by unit propagation")
			d.PrintStats()
			log.Println()
		}
		fmt.Fprintln(os.Stderr)
		fmt.Println("UNSATISFIABLE")
		os.Exit(20)
	}

	solution := d.SolveLimited()
	if d.Verbosity >= 1 {
		d.PrintStats()
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
