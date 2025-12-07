package fractions

import (
	"fmt"
	"math"

	"computorv1/src/calculations"
)

// ToFraction converts a float to an exact or tight rational using continued fractions.
// Returns an irreducible fraction string and true on success; false if no reasonable convergent found.
func ToFraction(value float64) (string, bool) {
	const maxDen = 1000000
	const tol = 1e-12

	sign := ""
	if value < 0 {
		sign = "-"
		value = -value
	}
	if calculations.AlmostZero(value) {
		return "0", true
	}

	num, den := continuedFraction(value, maxDen, tol)
	if den == 0 {
		return "", false
	}

	g := gcd(num, den)
	num /= g
	den /= g

	if den == 1 {
		return fmt.Sprintf("%s%d", sign, num), true
	}
	return fmt.Sprintf("%s%d/%d", sign, num, den), true
}

// continuedFraction finds a rational approximation using continued fractions within tol or maxDen.
func continuedFraction(x float64, maxDen int, tol float64) (int, int) {
	h1, h0 := 1, 0
	k1, k0 := 0, 1
	b := x
	for i := 0; i < 1000; i++ {
		a := int(math.Floor(b))
		h2 := a*h1 + h0
		k2 := a*k1 + k0

		if k2 > maxDen {
			break
		}

		if math.Abs(x-float64(h2)/float64(k2)) < tol {
			return h2, k2
		}

		h0, h1 = h1, h2
		k0, k1 = k1, k2

		frac := b - float64(a)
		if calculations.AlmostZero(frac) {
			break
		}
		b = 1.0 / frac
	}
	// Return best found so far
	if k1 == 0 {
		return 0, 0
	}
	return h1, k1
}

func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	if a == 0 {
		return 1
	}
	return a
}
