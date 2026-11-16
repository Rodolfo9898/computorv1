package factored

import (
	"fmt"

	"computorv1/src/core/formatter"
)

// Linear prints a * X + b as a * (X - x).
func Linear(a, x float64, useFractions bool) {
	l := lead(a, useFractions)
	if l == "" {
		show(factor(x, useFractions))
		return
	}
	show(l + parens(x, useFractions))
}

// Double prints a * X^2 + b * X + c, of zero discriminant, as a * (X - x)^2.
func Double(a, x float64, useFractions bool) {
	f := factor(x, useFractions)
	if x != 0 {
		f = "(" + f + ")"
	}
	show(lead(a, useFractions) + f + "^2")
}

// TwoReal prints a * (X - x1) * (X - x2).
func TwoReal(a, x1, x2 float64, useFractions bool) {
	show(lead(a, useFractions) + parens(x1, useFractions) + " * " + parens(x2, useFractions))
}

// Complex prints a * (X - z) * (X - conj(z)), with z = re + im * i.
func Complex(a, re, im float64, useFractions bool) {
	z1 := formatter.ComplexSolutionWithOptions(re, im, useFractions)
	z2 := formatter.ComplexSolutionWithOptions(re, -im, useFractions)
	show(lead(a, useFractions) + "(X - (" + z1 + ")) * (X - (" + z2 + "))")
}

func show(form string) {
	fmt.Printf("Factored form: %s = 0\n", form)
}

// lead is the leading coefficient before the factors: nothing for 1, "-" for -1.
func lead(a float64, useFractions bool) string {
	switch a {
	case 1:
		return ""
	case -1:
		return "-"
	}
	return formatter.SolutionWithOptions(a, useFractions) + " * "
}

// factor is X - x, written X + |x| for a negative x and X for zero.
func factor(x float64, useFractions bool) string {
	switch {
	case x == 0:
		return "X"
	case x < 0:
		return "X + " + formatter.SolutionWithOptions(-x, useFractions)
	}
	return "X - " + formatter.SolutionWithOptions(x, useFractions)
}

// parens puts a factor between parentheses, unless it is X alone.
func parens(x float64, useFractions bool) string {
	if x == 0 {
		return "X"
	}
	return "(" + factor(x, useFractions) + ")"
}
