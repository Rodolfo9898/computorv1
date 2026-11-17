//go:build bonus_mode

package factored

// Bonus feature 4: the polynomial as a product of its factors.

import (
	"fmt"

	"computorv1/sources/formatter"
	"computorv1/sources/solver"
)

// Show prints the polynomial as a product of its factors, when it has solutions;
// number writes each value.
func Show(s solver.Solution, number func(float64) string) {
	switch s.Kind {
	case solver.OneRoot:
		linear(s.B, s.X1, number)
	case solver.DoubleRoot:
		double(s.A, s.X1, number)
	case solver.TwoRealRoots:
		show(lead(s.A, number) + parens(s.X1, number) + " * " + parens(s.X2, number))
	case solver.TwoComplexRoots:
		z1 := formatter.Complex(s.Re, s.Im, number)
		z2 := formatter.Complex(s.Re, -s.Im, number)
		show(lead(s.A, number) + "(X - (" + z1 + ")) * (X - (" + z2 + "))")
	}
}

// linear prints a * X + b as a * (X - x).
func linear(a, x float64, number func(float64) string) {
	l := lead(a, number)
	if l == "" {
		show(factor(x, number))
		return
	}
	show(l + parens(x, number))
}

// double prints a * X^2 + b * X + c, of zero discriminant, as a * (X - x)^2.
func double(a, x float64, number func(float64) string) {
	f := factor(x, number)
	if x != 0 {
		f = "(" + f + ")"
	}
	show(lead(a, number) + f + "^2")
}

func show(form string) {
	fmt.Printf("Factored form: %s = 0\n", form)
}
