package execution

import (
	"fmt"

	"computorv1/src/base"
	"computorv1/src/bonus/check"
	"computorv1/src/bonus/factored"
	"computorv1/src/bonus/steps"
	"computorv1/src/calculations"
	"computorv1/src/core/formatter"
)

// Degree is the highest exponent whose coefficient is not zero, or 0 when none is.
// The parsers add terms with calculations.Sum, so a term that cancels out is
// exactly 0 here, and a small one (1e-10 * X^2) is not.
func Degree(coeffs base.PolyCoeffs) int {
	degree := 0
	for exp, coef := range coeffs {
		if coef != 0 && exp > degree {
			degree = exp
		}
	}
	return degree
}

func SolveWithOptions(coeffs base.PolyCoeffs, degree int, opts base.SolveOptions) {
	switch degree {
	case 0:
		solveDegreeZero(coeffs)
	case 1:
		solveDegreeOne(coeffs, opts)
	case 2:
		solveDegreeTwo(coeffs, opts)
	default:
		fmt.Println("The polynomial degree is strictly greater than 2, I can't solve.")
	}
}

func solveDegreeZero(coeffs base.PolyCoeffs) {
	if coeffs[0] == 0 {
		fmt.Println("Any real number is a solution.")
		return
	}
	fmt.Println("No solution.")
}

// solveDegreeOne solves a * X + b = 0, a not zero.
func solveDegreeOne(coeffs base.PolyCoeffs, opts base.SolveOptions) {
	a, b := coeffs[1], coeffs[0]
	solution := -b / a

	if opts.Verbose {
		steps.Linear(a, b)
	}
	fmt.Println("The solution is:")
	fmt.Printf("x = %s\n", formatter.SolutionWithOptions(solution, opts.UseFractions))
	if opts.Factored {
		factored.Linear(a, solution, opts.UseFractions)
	}
	if opts.Check {
		check.Real(coeffs, solution, opts.UseFractions)
	}
}

// solveDegreeTwo solves a * X^2 + b * X + c = 0, a not zero.
func solveDegreeTwo(coeffs base.PolyCoeffs, opts base.SolveOptions) {
	a, b, c := coeffs[2], coeffs[1], coeffs[0]
	discriminant := calculations.Sum(b*b, -4*a*c)
	if discriminant-discriminant != 0 { // +Inf or NaN: b * b or 4 * a * c overflowed
		fmt.Println("The coefficients are too large: the discriminant does not fit in a double.")
		return
	}
	if opts.Verbose {
		steps.Discriminant(a, b, c, discriminant)
	}

	switch {
	case discriminant == 0:
		solution := -b / (2 * a)
		fmt.Println("Discriminant is zero, the solution is:")
		if opts.Verbose {
			steps.Double(a, b)
		}
		fmt.Printf("x = %s\n", formatter.SolutionWithOptions(solution, opts.UseFractions))
		if opts.Factored {
			factored.Double(a, solution, opts.UseFractions)
		}
		if opts.Check {
			check.Real(coeffs, solution, opts.UseFractions)
		}

	case discriminant > 0:
		sqrtDisc := calculations.SqrtNewton(discriminant)
		x1, x2 := realRoots(a, b, c, sqrtDisc)
		fmt.Println("Discriminant is strictly positive, the two solutions are:")
		if opts.Verbose {
			steps.TwoReal(a, b, sqrtDisc)
		}
		fmt.Printf("x1 = %s\n", formatter.SolutionWithOptions(x1, opts.UseFractions))
		fmt.Printf("x2 = %s\n", formatter.SolutionWithOptions(x2, opts.UseFractions))
		if opts.Factored {
			factored.TwoReal(a, x1, x2, opts.UseFractions)
		}
		if opts.Check {
			check.Real(coeffs, x1, opts.UseFractions)
			check.Real(coeffs, x2, opts.UseFractions)
		}

	default:
		sqrtDisc := calculations.SqrtNewton(-discriminant)
		realPart := -b / (2 * a)
		imagPart := sqrtDisc / (2 * a)
		fmt.Println("Discriminant is strictly negative, the two complex solutions are:")
		if opts.Verbose {
			steps.Complex(a, b, sqrtDisc)
		}
		fmt.Printf("x1 = %s\n", formatter.ComplexSolutionWithOptions(realPart, imagPart, opts.UseFractions))
		fmt.Printf("x2 = %s\n", formatter.ComplexSolutionWithOptions(realPart, -imagPart, opts.UseFractions))
		if opts.Factored {
			factored.Complex(a, realPart, imagPart, opts.UseFractions)
		}
		if opts.Check {
			check.Complex(coeffs, realPart, imagPart, opts.UseFractions)
			check.Complex(coeffs, realPart, -imagPart, opts.UseFractions)
		}
	}
}

// realRoots gives x1 = (-b - √Δ) / 2a and x2 = (-b + √Δ) / 2a without subtracting
// two close numbers: when b² is much larger than 4ac, -b + √Δ (b > 0) would lose
// its digits. That root comes from the other one instead, as c / q, since
// x1 * x2 = c / a.
func realRoots(a, b, c, sqrtDisc float64) (float64, float64) {
	if b >= 0 {
		q := -(b + sqrtDisc) / 2
		return q / a, c / q
	}
	q := (sqrtDisc - b) / 2
	return c / q, q / a
}
