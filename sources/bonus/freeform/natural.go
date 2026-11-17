//go:build bonus_mode

package freeform

import (
	"fmt"

	"computorv1/sources/formatter"
	"computorv1/sources/poly"
)

// Natural builds the reduced form as it is written by hand, as in the subject's
// free form example: "5 + 4 * X", "X^2 - X", "-X". Zero terms are left out.
func Natural(coeffs poly.PolyCoeffs) string {
	keep := func(coef float64, exp int) bool { return coef != 0 }
	return formatter.Join(coeffs, keep, func(coef float64, exp int) string {
		if exp == 0 {
			return formatter.Number(coef)
		}
		x := "X"
		if exp > 1 {
			x = fmt.Sprintf("X^%d", exp)
		}
		if coef == 1 {
			return x
		}
		return formatter.Number(coef) + " * " + x
	}, "0")
}
