package main

import (
	"flag"
	"log"
	"os"

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
	err := dpll.DecodeFile(d, flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	isSat := d.Solve()
	if isSat {
		log.Printf("SATISFIABLE")
	} else {
		log.Printf("UNSATISFIABLE")
	}
}
