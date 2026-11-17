package main

import (
	"fmt"
	"io"
	"os"

	"computorv1/sources/formatter"
	"computorv1/sources/input"
	"computorv1/sources/solver"
)

// ./computor_v1, the mandatory program (make): solves the equation of the
// arguments, or of the standard input when there are none. It uses the shared
// folders of sources/ (input, poly, numeric, equation, formatter, solver) and never
// sources/bonus, which make does not even compile. A refused equation exits with
// status 1, its reason on stderr.
func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		switch {
		case input.IsHelp(args[0]):
			printUsage(os.Stdout)
			return
		case input.IsFlag(args[0]):
			refuse(fmt.Errorf("unknown flag or command: %s", args[0]))
		}
	}

	equation, err := input.Read(args)
	if err != nil {
		refuse(err)
	}
	eq, err := parseEquation(equation)
	if err != nil {
		refuse(err)
	}

	fmt.Printf("Reduced form: %s = 0\n", formatter.Reduced(eq.Coeffs))
	fmt.Printf("Polynomial degree: %d\n", eq.Coeffs.Degree())
	solver.Print(solver.Solve(eq), solver.Hooks{})
}

// refuse prints why the equation is refused, then the usage, and exits with 1.
func refuse(err error) {
	fmt.Fprintln(os.Stderr, "Error:", err)
	fmt.Fprintln(os.Stderr)
	printUsage(os.Stderr)
	os.Exit(1)
}

// printUsage writes how to call the mandatory program.
func printUsage(w io.Writer) {
	lines := []string{
		"Usage: ./computor_v1 \"<lhs> = <rhs>\"",
		"       ./computor_v1                 # the equation read on the standard input",
		" - Terms must be in the strict form: <coef> * X^<exp>",
		" - Coefficients can be negative or decimal; exponents are whole numbers",
		" - Degree > 2 will be reported as unsupported",
		"Examples:",
		"  ./computor_v1 \"5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0\"",
		"  ./computor_v1 \"6 * X^0 = 6 * X^0\"",
		"  ./computor_v1 \"10 * X^0 = 15 * X^0\"",
	}
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
}
