//go:build bonus_mode

package check

// Bonus feature 5: each solution put back into the reduced form, P(x) = 0.

import (
	"fmt"

	"computorv1/sources/formatter"
	"computorv1/sources/numeric"
	"computorv1/sources/poly"
	"computorv1/sources/solver"
)

// Show prints P(x), the left side of the reduced form, at each solution: 0 when
// the solution is right. number writes the solutions.
func Show(s solver.Solution, number func(float64) string) {
	switch s.Kind {
	case solver.OneRoot, solver.DoubleRoot:
		show(s.Coeffs, s.X1, 0, number(s.X1))
	case solver.TwoRealRoots:
		show(s.Coeffs, s.X1, 0, number(s.X1))
		show(s.Coeffs, s.X2, 0, number(s.X2))
	case solver.TwoComplexRoots:
		show(s.Coeffs, s.Re, s.Im, formatter.Complex(s.Re, s.Im, number))
		show(s.Coeffs, s.Re, -s.Im, formatter.Complex(s.Re, -s.Im, number))
	}
}

func show(coeffs poly.PolyCoeffs, re, im float64, label string) {
	pRe, pIm, scale := evaluate(coeffs, re, im)
	value := "0"
	if numeric.Abs(pRe) > relTol*scale || numeric.Abs(pIm) > relTol*scale {
		value = formatter.Number(pRe)
		if im != 0 {
			value = formatter.Complex(pRe, pIm, formatter.Number)
		}
		value += " (not zero)"
	}
	fmt.Printf("Check: P(%s) = %s\n", label, value)
}

// evaluate gives P(z), the sum of coef * z^exp, and the sum of the sizes of its
// terms (|coef| * (|re| + |im|)^exp), to which the result is compared.
func evaluate(coeffs poly.PolyCoeffs, re, im float64) (float64, float64, float64) {
	var sumRe, sumIm, scale float64
	size := numeric.Abs(re) + numeric.Abs(im)
	for exp, coef := range coeffs {
		if coef == 0 {
			continue
		}
		powRe, powIm, powSize := 1.0, 0.0, 1.0
		for i := 0; i < exp; i++ {
			powRe, powIm = powRe*re-powIm*im, powRe*im+powIm*re
			powSize *= size
		}
		sumRe += coef * powRe
		sumIm += coef * powIm
		scale += numeric.Abs(coef) * powSize
	}
	return sumRe, sumIm, scale
}
