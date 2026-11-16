package steps

import (
	"fmt"

	"computorv1/src/core/formatter"
)

// Linear prints the solution of a * X + b = 0 with the values put in.
func Linear(a, b float64) {
	fmt.Printf("Linear solve: x = -b / a = -(%s) / %s\n", formatter.Number(b), formatter.Number(a))
}

// Discriminant prints how Δ of a * X^2 + b * X + c is worked out.
func Discriminant(a, b, c, discriminant float64) {
	fmt.Printf("Discriminant Δ = b^2 - 4ac = (%s)^2 - 4 * (%s) * (%s) = %s\n",
		formatter.Number(b), formatter.Number(a), formatter.Number(c), formatter.Number(discriminant))
}

// Double prints the single solution of a zero discriminant.
func Double(a, b float64) {
	fmt.Printf("x = -b / (2a) = -(%s) / (2 * %s)\n", formatter.Number(b), formatter.Number(a))
}

// TwoReal prints the two formulas of a positive discriminant.
func TwoReal(a, b, sqrtDisc float64) {
	fmt.Printf("sqrt(Δ) = %s\n", formatter.Number(sqrtDisc))
	fmt.Printf("x1 = (-b - sqrt(Δ)) / (2a) = (-(%s) - %s) / (2 * %s)\n", formatter.Number(b), formatter.Number(sqrtDisc), formatter.Number(a))
	fmt.Printf("x2 = (-b + sqrt(Δ)) / (2a) = (-(%s) + %s) / (2 * %s)\n", formatter.Number(b), formatter.Number(sqrtDisc), formatter.Number(a))
}

// Complex prints the real and imaginary parts of a negative discriminant.
func Complex(a, b, sqrtDisc float64) {
	fmt.Printf("sqrt(-Δ) = %s\n", formatter.Number(sqrtDisc))
	fmt.Printf("Real part = -b / (2a) = -(%s) / (2 * %s)\n", formatter.Number(b), formatter.Number(a))
	fmt.Printf("Imag part = sqrt(-Δ) / (2a) = %s / (2 * %s)\n", formatter.Number(sqrtDisc), formatter.Number(a))
}
