package check

import (
	"fmt"

	"computorv1/src/base"
	"computorv1/src/calculations"
	"computorv1/src/core/formatter"
)

// relTol: P(x) is zero when it is this small next to the size of its terms. A
// solution is rounded to a double, so P(x) is rarely exactly 0.
const relTol = 1e-9

// Real prints P(x), the left side of the reduced form at the solution x: 0 when
// the solution is right.
func Real(coeffs base.PolyCoeffs, x float64, useFractions bool) {
	show(coeffs, x, 0, formatter.SolutionWithOptions(x, useFractions))
}

// Complex prints P(z) for z = re + im * i, worked out in complex numbers.
func Complex(coeffs base.PolyCoeffs, re, im float64, useFractions bool) {
	show(coeffs, re, im, formatter.ComplexSolutionWithOptions(re, im, useFractions))
}

func show(coeffs base.PolyCoeffs, re, im float64, label string) {
	pRe, pIm, scale := evaluate(coeffs, re, im)
	value := "0"
	if calculations.Abs(pRe) > relTol*scale || calculations.Abs(pIm) > relTol*scale {
		value = formatter.Number(pRe)
		if im != 0 {
			value = formatter.ComplexSolutionWithOptions(pRe, pIm, false)
		}
		value += " (not zero)"
	}
	fmt.Printf("Check: P(%s) = %s\n", label, value)
}

// evaluate gives P(z), the sum of coef * z^exp, and the sum of the sizes of its
// terms (|coef| * (|re| + |im|)^exp), to which the result is compared.
func evaluate(coeffs base.PolyCoeffs, re, im float64) (float64, float64, float64) {
	var sumRe, sumIm, scale float64
	size := calculations.Abs(re) + calculations.Abs(im)
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
		scale += calculations.Abs(coef) * powSize
	}
	return sumRe, sumIm, scale
}
