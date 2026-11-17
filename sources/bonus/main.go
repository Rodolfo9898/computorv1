//go:build bonus_mode

package main

import (
	"fmt"
	"io"
	"os"

	"computorv1/sources/bonus/freeform"
	"computorv1/sources/bonus/mistakes"
	"computorv1/sources/input"
	"computorv1/sources/solver"
)

// ./computor_v1_bonus, the bonus program (make bonus): every file of sources/bonus
// starts with //go:build bonus_mode, the BONUS_MODE of a C Makefile, so only
// go build -tags bonus_mode compiles them. As in the subject (p.8): free form
// entry, always on and graded on its own (freeform/), then five additional
// features, one folder each, which the mode turns on:
//
//  1. entry mistakes, vocabulary and syntax   mistakes/
//  2. solutions as irreducible fractions      fractions/
//  3. intermediate calculation steps          steps/
//  4. the factored form                       factored/
//  5. the check of each solution, P(x) = 0    check/
func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		switch {
		case input.IsHelp(args[0]):
			printUsage(os.Stdout)
			return
		case input.IsFlag(args[0]):
			refuse(fmt.Errorf("unknown flag or command: %s", args[0]), plain)
		}
	}

	m, eqArgs := parseMode(args)
	report := plain
	if m.Mistakes {
		report = mistakes.Print
	}
	equation, err := input.Read(eqArgs) // no equation: one line of the standard input
	if err != nil {
		refuse(err, report)
	}
	if m.Mistakes {
		if err := mistakes.Check(equation); err != nil {
			refuse(err, report)
		}
	}
	eq, err := freeform.ParseEquation(equation)
	if err != nil {
		refuse(fmt.Errorf("Parse error: %v", err), report)
	}

	fmt.Printf("Reduced form: %s = 0\n", freeform.Natural(eq.Coeffs))
	fmt.Printf("Polynomial degree: %d\n", eq.Coeffs.Degree())
	solver.Print(solver.Solve(eq), m.hooks())
}

// plain writes a refused equation without colour, when the mistakes feature is off.
func plain(w io.Writer, err error) {
	fmt.Fprintln(w, err)
}

// refuse reports why the equation is refused, then the usage, and exits with 1.
func refuse(err error, report func(io.Writer, error)) {
	report(os.Stderr, err)
	fmt.Fprintln(os.Stderr)
	printUsage(os.Stderr)
	os.Exit(1)
}

// printUsage writes how to call the bonus program.
func printUsage(w io.Writer) {
	lines := []string{
		"Usage: ./computor_v1_bonus [mode] \"<lhs> = <rhs>\"",
		"       ./computor_v1_bonus [mode]          # the equation read on the standard input",
		"Free form entry is always on: 'X', '4 * X', '4X', 'X^3', '-5', in any order,",
		"and the reduced form is written as by hand (5 + 4 * X = 0).",
		"Modes:",
		" - all: the five features below [default]",
		" - mistakes: 1. vocabulary and syntax mistakes, named and coloured by kind",
		" - fractions: 2. solutions as irreducible fractions, when they are exact",
		" - steps: 3. the intermediate calculation steps",
		" - factored: 4. the polynomial as a product of its factors",
		" - check: 5. each solution put back into the reduced form, P(x) = 0",
		" - freeform: free form entry only",
		"Examples:",
		"  ./computor_v1_bonus freeform \"5 + 4 * X + X^2= X^2\"",
		"  ./computor_v1_bonus mistakes \"hello + X = 0\"",
		"  ./computor_v1_bonus fractions \"3 * X = 1\"",
		"  ./computor_v1_bonus steps \"X^2 + X = 0\"",
		"  ./computor_v1_bonus factored \"X^2 - 3 * X + 2 = 0\"",
		"  ./computor_v1_bonus check \"X^2 + 2X + 5 = 0\"",
	}
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
}
