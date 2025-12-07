package steps

import (
	"fmt"

	"computorv1/src/core/formatter"
)

// Quadratic prints intermediate steps for quadratic solving.
func Quadratic(a, b, c, discriminant, sqrtDisc float64) {
	fmt.Printf("Discriminant Δ = b^2 - 4ac = %s^2 - 4*%s*%s = %s\n", formatter.Number(b), formatter.Number(a), formatter.Number(c), formatter.Number(discriminant))
	fmt.Printf("sqrt(|Δ|) ≈ %s\n", formatter.Number(sqrtDisc))
	fmt.Printf("Real part = -b / (2a) = -%s / (2*%s)\n", formatter.Number(b), formatter.Number(a))
	fmt.Printf("Imag part = sqrt(|Δ|) / (2a) = %s / (2*%s)\n", formatter.Number(sqrtDisc), formatter.Number(a))
}

// Linear prints intermediate steps for linear solving.
func Linear(a, b float64) {
	fmt.Printf("Linear solve: -b/a = -(%s)/(%s)\n", formatter.Number(b), formatter.Number(a))
}
