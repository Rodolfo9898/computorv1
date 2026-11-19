//go:build bonus_mode

package freeform

import (
	"fmt"

	"computorv1/sources/formatter"
	"computorv1/sources/poly"
)

// Natural builds the reduced form as it is written by hand, as in the subject's
// free form example: "5 + 4 * X", "X^2 - X", "-X". Zero terms are left out, and
// each coefficient is written with the digits it needs (formatter.Coefficient).
func Natural(eq poly.Equation) string {
	keep := func(coef float64, exp int) bool { return coef != 0 }
	return formatter.Join(eq.Coeffs, keep, func(coef float64, exp int) string {
		number := formatter.Coefficient(coef, eq.Noise[exp])
		if exp == 0 {
			return number
		}
		x := "X"
		if exp > 1 {
			x = fmt.Sprintf("X^%d", exp)
		}
		if coef == 1 {
			return x
		}
		return number + " * " + x
	}, "0")
}
