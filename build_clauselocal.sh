#!/bin/bash

MAXLOCAL=$1
if [ -z "$MAXLOCAL" ]; then
    MAXLOCAL=64
fi

echo "package dpll" > clauselocal.go

echo "var mkClauseLocal = []func([]Lit) *Clause {" >> clauselocal.go
for i in {1..64}; do
    echo "	$i: mkClauseLocal$i," >>clauselocal.go
done
echo "}" >> clauselocal.go

echo "var mkClauseExtraLocal = []func([]Lit) *Clause {" >> clauselocal.go
for i in {1..64}; do
    echo "	$i: mkClauseExtraLocal$i," >>clauselocal.go
done
echo "}" >> clauselocal.go

for i in {1..64}; do
    TYPES="
type clauseLocal$i struct {
	Clause
	Lit [$i]Lit
}
type clauseExtraLocal$i struct {
	Clause
	ClauseExtra
	Lit [$i]Lit
}"
    echo "$TYPES" >> clauselocal.go
done

for i in {1..64}; do
    FUNCS="
func mkClauseLocal$i(ps []Lit) *Clause {
	c := &clauseLocal$i{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal$i(ps []Lit) *Clause {
	c := &clauseExtraLocal$i{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
    c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}"
    echo "$FUNCS" >>clauselocal.go
done

goimports -w clauselocal.go
