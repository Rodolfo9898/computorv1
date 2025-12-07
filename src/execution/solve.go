package execution

import (
	"fmt"

	"computorv1/src/base"
	"computorv1/src/bonus/steps"
	"computorv1/src/calculations"
	"computorv1/src/core/formatter"
)

func Degree(coeffs base.PolyCoeffs) int {
	degree := -1
	for exp, coef := range coeffs {
		if !calculations.AlmostZero(coef) && exp > degree {
			degree = exp
		}
	}

	if degree < 0 {
		return 0
	}
	return degree
}

func Solve(coeffs base.PolyCoeffs, degree int) {
	SolveWithOptions(coeffs, degree, base.SolveOptions{})
}

func SolveWithOptions(coeffs base.PolyCoeffs, degree int, opts base.SolveOptions) {
	switch degree {
	case 0:
		solveDegreeZero(coeffs, opts)
	case 1:
		solveDegreeOne(coeffs, opts)
	case 2:
		solveDegreeTwo(coeffs, opts)
	default:
		fmt.Println("The polynomial degree is strictly greater than 2, I can't solve.")
	}
}

func solveDegreeZero(coeffs base.PolyCoeffs, opts base.SolveOptions) {
	constant := calculations.CleanZero(coeffs[0])
	if calculations.AlmostZero(constant) {
		fmt.Println("Any real number is a solution.")
		return
	}

	fmt.Println("No solution.")
}

func solveDegreeOne(coeffs base.PolyCoeffs, opts base.SolveOptions) {
	a := calculations.CleanZero(coeffs[1])
	b := calculations.CleanZero(coeffs[0])

	if calculations.AlmostZero(a) {
		solveDegreeZero(coeffs, opts)
		return
	}

	solution := -b / a
	if opts.Verbose {
		steps.Linear(a, b)
	}
	fmt.Println("The solution is:")
	fmt.Printf("x = %s\n", formatter.SolutionWithOptions(solution, opts.UseFractions))
}

func solveDegreeTwo(coeffs base.PolyCoeffs, opts base.SolveOptions) {
	a := calculations.CleanZero(coeffs[2])
	b := calculations.CleanZero(coeffs[1])
	c := calculations.CleanZero(coeffs[0])

	if calculations.AlmostZero(a) {
		solveDegreeOne(coeffs, opts)
		return
	}

	discriminant := b*b - 4*a*c
	if opts.Verbose {
		steps.Quadratic(a, b, c, discriminant, calculations.SqrtNewton(calculations.Abs(discriminant)))
	}

	if calculations.AlmostZero(discriminant) {
		solution := -b / (2 * a)
		fmt.Println("Discriminant is zero, the solution is:")
		if opts.Verbose {
			fmt.Printf("x = -b / (2a) = -%s / (2*%s)\n", formatter.Number(b), formatter.Number(a))
		}
		fmt.Printf("x = %s\n", formatter.SolutionWithOptions(solution, opts.UseFractions))
		return
	}

	if discriminant > 0 {
		sqrtDisc := calculations.SqrtNewton(discriminant)
		s1 := (-b - sqrtDisc) / (2 * a)
		s2 := (-b + sqrtDisc) / (2 * a)
		fmt.Println("Discriminant is strictly positive, the two solutions are:")
		if opts.Verbose {
			fmt.Printf("x1 = (-b - sqrt(Δ)) / (2a) = (-(%s) - %s) / (2*%s)\n", formatter.Number(b), formatter.Number(sqrtDisc), formatter.Number(a))
			fmt.Printf("x2 = (-b + sqrt(Δ)) / (2a) = (-(%s) + %s) / (2*%s)\n", formatter.Number(b), formatter.Number(sqrtDisc), formatter.Number(a))
		}
		fmt.Printf("x1 = %s\n", formatter.SolutionWithOptions(s1, opts.UseFractions))
		fmt.Printf("x2 = %s\n", formatter.SolutionWithOptions(s2, opts.UseFractions))
		return
	}

	sqrtDisc := calculations.SqrtNewton(-discriminant)
	realPart := -b / (2 * a)
	imagPart := sqrtDisc / (2 * a)
	if opts.Verbose {
		steps.Quadratic(a, b, c, discriminant, sqrtDisc)
	}
	fmt.Println("Discriminant is strictly negative, the two complex solutions are:")
	fmt.Printf("x1 = %s\n", formatter.ComplexSolutionWithOptions(realPart, imagPart, opts.UseFractions))
	fmt.Printf("x2 = %s\n", formatter.ComplexSolutionWithOptions(realPart, -imagPart, opts.UseFractions))
}
