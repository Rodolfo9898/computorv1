//go:build bonus_mode

package fractions

// Bonus feature 2: the solutions as irreducible fractions, when relevant.

import (
	"fmt"

	"computorv1/sources/formatter"
	"computorv1/sources/numeric"
)

// Number writes value as an irreducible fraction when it is one, as a decimal
// otherwise: what the solutions are written with in the fractions mode.
func Number(value float64) string {
	if frac, ok := ToFraction(value); ok {
		return frac
	}
	return formatter.Number(value)
}

// ToFraction gives value as an irreducible fraction, and true, when it is one: -5/4
// for -1.25. An irrational value (sqrt(2)) has no such fraction and gives false,
// so that an approximation never passes for the exact solution.
func ToFraction(value float64) (string, bool) {
	sign := ""
	if value < 0 {
		sign = "-"
		value = -value
	}
	if value == 0 {
		return "0", true
	}
	if value > maxValue {
		return "", false
	}

	num, den, ok := continuedFraction(value)
	if !ok {
		return "", false
	}
	if den == 1 {
		return fmt.Sprintf("%s%d", sign, num), true
	}
	return fmt.Sprintf("%s%d/%d", sign, num, den), true
}

// continuedFraction walks the convergents h/k of x (each one already irreducible)
// and stops at the first equal to x up to rounding. It gives up when the
// denominator would pass maxDen.
func continuedFraction(x float64) (int, int, bool) {
	h1, h0 := 1, 0
	k1, k0 := 0, 1
	b := x
	for {
		if k1 > 0 && b > maxDen { // the next denominator would pass maxDen
			return 0, 0, false
		}
		a := int(b) // b >= 0: the conversion is the floor
		h2 := a*h1 + h0
		k2 := a*k1 + k0
		if k2 > maxDen {
			return 0, 0, false
		}
		if numeric.Abs(x-float64(h2)/float64(k2)) <= relTol*x {
			return h2, k2, true
		}

		h0, h1 = h1, h2
		k0, k1 = k1, k2
		frac := b - float64(a)
		if frac == 0 {
			return 0, 0, false
		}
		b = 1 / frac
	}
}
