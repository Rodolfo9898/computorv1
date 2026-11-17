package formatter

import (
	"fmt"
	"sort"
	"strings"

	"computorv1/sources/poly"
)

// Reduced builds the reduced form of the subject: every term as a * X^p. A zero
// term written in the equation is kept up to the degree, as in the subject's
// "5 * X^0 - 6 * X^1 + 0 * X^2 - 5.6 * X^3"; above the degree, a term is zero
// because it cancelled out, and is left out.
func Reduced(coeffs poly.PolyCoeffs) string {
	degree := coeffs.Degree()
	keep := func(coef float64, exp int) bool { return exp <= degree }
	return Join(coeffs, keep, func(coef float64, exp int) string {
		return fmt.Sprintf("%s * X^%d", Number(coef), exp)
	}, "0 * X^0")
}

// Join writes the terms that keep accepts, by exponent, with " + " or " - "
// between them: term gets the absolute value of the coefficient.
func Join(coeffs poly.PolyCoeffs, keep func(coef float64, exp int) bool,
	term func(coef float64, exp int) string, empty string) string {
	var builder strings.Builder
	for _, exp := range sortedExponents(coeffs) {
		coef := coeffs[exp]
		if !keep(coef, exp) {
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

// Complex writes re + im i, each part written by number.
func Complex(re, im float64, number func(float64) string) string {
	sign := "+"
	if im < 0 {
		sign = "-"
		im = -im
	}
	return fmt.Sprintf("%s %s %si", number(re), sign, number(im))
}

func sortedExponents(coeffs poly.PolyCoeffs) []int {
	exps := make([]int, 0, len(coeffs))
	for exp := range coeffs {
		exps = append(exps, exp)
	}
	sort.Ints(exps)
	return exps
}
