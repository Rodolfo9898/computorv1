package formatter

import (
	"fmt"
	"sort"
	"strings"

	"computorv1/src/base"
	"computorv1/src/bonus/fractions"
	"computorv1/src/calculations"
)

// Reduced builds a human-readable reduced polynomial form.
func Reduced(coeffs base.PolyCoeffs) string {
	var builder strings.Builder
	exponents := sortedExponents(coeffs)
	written := false

	for _, exp := range exponents {
		coef := calculations.CleanZero(coeffs[exp])
		if calculations.AlmostZero(coef) {
			continue
		}

		sign := " + "
		display := coef
		if coef < 0 {
			sign = " - "
			display = -coef
		}

		if !written {
			if coef < 0 {
				builder.WriteString("-")
			}
			builder.WriteString(fmt.Sprintf("%s * X^%d", Number(display), exp))
			written = true
			continue
		}

		builder.WriteString(sign)
		builder.WriteString(fmt.Sprintf("%s * X^%d", Number(display), exp))
	}

	if !written {
		return "0 * X^0"
	}
	return builder.String()
}

// Number formats a float with trimmed trailing zeros.
func Number(value float64) string {
	v := calculations.CleanZero(value)
	out := fmt.Sprintf("%.10g", v)
	out = strings.TrimRight(out, "0")
	out = strings.TrimRight(out, ".")
	if out == "" {
		return "0"
	}
	return out
}

// Solution tries to display an interesting irreducible fraction, otherwise a trimmed decimal.
func Solution(value float64) string {
	if frac, ok := fractions.ToFraction(value); ok {
		return frac
	}
	return Number(value)
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

// ComplexSolution formats complex numbers using Solution for parts.
func ComplexSolution(real, imag float64) string {
	sign := "+"
	value := imag
	if value < 0 {
		sign = "-"
		value = -value
	}

	return fmt.Sprintf("%s %s %si", Solution(real), sign, Solution(value))
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

// Complex formats a complex root using Number for consistency.
func Complex(real, imag float64) string {
	sign := "+"
	value := imag
	if value < 0 {
		sign = "-"
		value = -value
	}

	return fmt.Sprintf("%s %s %si", Number(real), sign, Number(value))
}

func sortedExponents(coeffs base.PolyCoeffs) []int {
	exps := make([]int, 0, len(coeffs))
	for exp := range coeffs {
		exps = append(exps, exp)
	}
	sort.Ints(exps)
	if len(exps) == 0 {
		return []int{0}
	}
	return exps
}
