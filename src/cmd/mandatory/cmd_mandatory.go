package mandatory

import (
	"fmt"
	"os"

	"computorv1/src/base"
	"computorv1/src/bonus/colors"
	"computorv1/src/core/formatter"
	"computorv1/src/core/parser"
	"computorv1/src/execution"
	input "computorv1/src/io"
)

// Run solves the equation of the arguments, or of the standard input when there
// are none. A refused equation exits with status 1, its reason on stderr.
func Run(args []string) {
	argument := input.Read(args)
	coeffs, err := parser.ParseEquation(argument)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		PrintUsage()
		os.Exit(1)
	}

	fmt.Printf("Reduced form: %s = 0\n", formatter.Reduced(coeffs))

	degree := execution.Degree(coeffs)
	fmt.Printf("Polynomial degree: %d\n", degree)

	execution.SolveWithOptions(coeffs, degree, base.SolveOptions{})
}

func PrintUsage() {
	lines := []string{
		"\nUsage: ./computor_v1 \"<lhs> = <rhs>\"",
		" - Terms must be in the strict form: <coef> * X^<exp>",
		" - Coefficients can be negative or decimal; exponents are whole numbers",
		" - Degree > 2 will be reported as unsupported",
		"Examples:",
		"  ./computor_v1 \"5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0\"",
		"  ./computor_v1 \"6 * X^0 = 6 * X^0\"",
		"  ./computor_v1 \"10 * X^0 = 15 * X^0\"",
	}
	for _, line := range lines {
		fmt.Fprintln(os.Stderr, colors.Wrap(colors.Grey, line))
	}
}
