package equation

import (
	"fmt"
	"strings"

	"computorv1/src/numeric"
	"computorv1/src/poly"
)

// Term is one term of a side, as it was read: "4 * X^1" is {Exp: 1, Coef: 4}.
type Term struct {
	Exp  int
	Coef float64
}

// Parse cuts input at its '=', reads both sides with side, then moves every
// term to the left: the equation reduced to P(X) = 0. Both programs use it, each
// with its own side parser: strict in the mandatory one, free form in the bonus.
func Parse(input string, side func(string) ([]Term, error)) (poly.Equation, error) {
	parts := strings.Split(input, "=")
	if len(parts) != 2 {
		return poly.Equation{}, fmt.Errorf("equation must contain exactly one '='")
	}

	left, err := side(parts[0])
	if err != nil {
		return poly.Equation{}, err
	}
	right, err := side(parts[1])
	if err != nil {
		return poly.Equation{}, err
	}
	return reduce(left, right)
}

// reduce adds up the terms of each exponent, those of the right side subtracted.
// A coefficient no larger than the rounding noise of its sum is a cancellation
// and becomes exactly 0; a small one written on its own (1e-10 * X^2) is far
// above its noise, and stays. Sizes that overflow a double are refused, rather
// than printed as +Inf.
func reduce(left, right []Term) (poly.Equation, error) {
	eq := poly.Equation{Coeffs: poly.PolyCoeffs{}, Noise: poly.PolyCoeffs{}}
	size := map[int]float64{}
	count := map[int]int{}
	add := func(terms []Term, sign float64) {
		for _, term := range terms {
			eq.Coeffs[term.Exp] += sign * term.Coef
			size[term.Exp] += numeric.Abs(term.Coef)
			count[term.Exp]++
		}
	}
	add(left, 1)
	add(right, -1)

	for exp, coef := range eq.Coeffs {
		if !numeric.IsFinite(size[exp]) {
			return poly.Equation{}, fmt.Errorf("the coefficients of X^%d are too large: their sum does not fit in a double", exp)
		}
		eq.Noise[exp] = numeric.SumNoise(count[exp], size[exp])
		if numeric.Abs(coef) <= eq.Noise[exp] {
			eq.Coeffs[exp] = 0
		}
	}
	return eq, nil
}
