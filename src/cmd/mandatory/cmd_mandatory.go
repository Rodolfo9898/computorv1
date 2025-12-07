package mandatory

import (
	"fmt"

	"computorv1/src/base"
	"computorv1/src/bonus/colors"
	"computorv1/src/core/formatter"
	"computorv1/src/core/parser"
	"computorv1/src/execution"
	input "computorv1/src/io"
)

func Run(args []string) {
	argument := input.Read(args)
	coeffs, err := parser.ParseEquation(argument)
	if err != nil {
		fmt.Println("Error:", err)
		PrintUsage()
		return
	}

	fmt.Printf("Reduced form: %s = 0\n", formatter.Reduced(coeffs))

	degree := execution.Degree(coeffs)
	fmt.Printf("Polynomial degree: %d\n", degree)

	execution.SolveWithOptions(coeffs, degree, base.SolveOptions{
		Verbose:      false,
		UseFractions: false,
	})
}

func PrintUsage() {
	fmt.Println(colors.Wrap(colors.Grey, "\nUsage: ./computor_v1 \"<lhs> = <rhs>\""))
	fmt.Println(colors.Wrap(colors.Grey, " - Terms must be in the strict form: <coef> * X^<exp>"))
	fmt.Println(colors.Wrap(colors.Grey, " - Exponents and coefficients can be negative or decimal"))
	fmt.Println(colors.Wrap(colors.Grey, " - Degree > 2 will be reported as unsupported"))
	fmt.Println(colors.Wrap(colors.Grey, "Examples:"))
	fmt.Println(colors.Wrap(colors.Grey, "  ./computor_v1 \"5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0\""))
	fmt.Println(colors.Wrap(colors.Grey, "  ./computor_v1 \"6 * X^0 = 6 * X^0\""))
	fmt.Println(colors.Wrap(colors.Grey, "  ./computor_v1 \"10 * X^0 = 15 * X^0\""))
}
