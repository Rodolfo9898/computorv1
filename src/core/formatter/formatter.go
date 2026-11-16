package formatter

import (
	"fmt"
	"sort"
	"strings"

	"computorv1/src/base"
	"computorv1/src/bonus/fractions"
)

// Reduced builds the reduced form of the subject: every term as a * X^p.
func Reduced(coeffs base.PolyCoeffs) string {
	return join(coeffs, func(coef float64, exp int) string {
		return fmt.Sprintf("%s * X^%d", Number(coef), exp)
	}, "0 * X^0")
}

// Natural builds the reduced form as it is written by hand, like the free form
// entry of the subject: "5 + 4 * X", "X^2 - X", "-X".
func Natural(coeffs base.PolyCoeffs) string {
	return join(coeffs, func(coef float64, exp int) string {
		if exp == 0 {
			return Number(coef)
		}
		x := "X"
		if exp > 1 {
			x = fmt.Sprintf("X^%d", exp)
		}
		if coef == 1 {
			return x
		}
		return Number(coef) + " * " + x
	}, "0")
}

// join writes the terms that are not zero, by exponent, with " + " or " - "
// between them: term gets the absolute value of the coefficient.
func join(coeffs base.PolyCoeffs, term func(coef float64, exp int) string, empty string) string {
	var builder strings.Builder
	for _, exp := range sortedExponents(coeffs) {
		coef := coeffs[exp]
		if coef == 0 {
			continue
		}

		display := coef
		if coef < 0 {
			display = -coef
		}
		switch {
		case builder.Len() == 0 && coef < 0:
			builder.WriteString("-")
		case builder.Len() > 0 && coef < 0:
			builder.WriteString(" - ")
		case builder.Len() > 0:
			builder.WriteString(" + ")
		}
		builder.WriteString(term(display, exp))
	}

	if builder.Len() == 0 {
		return empty
	}
	return builder.String()
}

// Number formats a float with up to 10 significant digits. %g already leaves out
// the zeros at the end of the decimals, and only those: 10 stays "10". A -0 left
// by a computation prints as 0.
func Number(value float64) string {
	if value == 0 {
		return "0"
	}
	return fmt.Sprintf("%.10g", value)
}

// SolutionWithOptions formats using fractions if requested, otherwise Number.
func SolutionWithOptions(value float64, useFractions bool) string {
	if useFractions {
		if frac, ok := fractions.ToFraction(value); ok {
			return frac
		}
	}
	return Number(value)
}

// ComplexSolutionWithOptions formats complex numbers using SolutionWithOptions.
func ComplexSolutionWithOptions(real, imag float64, useFractions bool) string {
	sign := "+"
	value := imag
	if value < 0 {
		sign = "-"
		value = -value
	}

	return fmt.Sprintf("%s %s %si", SolutionWithOptions(real, useFractions), sign, SolutionWithOptions(value, useFractions))
}

func sortedExponents(coeffs base.PolyCoeffs) []int {
	exps := make([]int, 0, len(coeffs))
	for exp := range coeffs {
		exps = append(exps, exp)
	}
	sort.Ints(exps)
	return exps
}
