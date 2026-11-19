//go:build bonus_mode

package steps

// Bonus feature 3: the intermediate calculation steps.

import (
	"fmt"

	"computorv1/sources/formatter"
	"computorv1/sources/solver"
)

// BeforeCase shows how the case is found: the formula of degree 1, or Δ worked
// out with the values put in.
func BeforeCase(s solver.Solution) {
	switch s.Kind {
	case solver.OneRoot: // b * X + c = 0: a is b and b is c in -b / a
		fmt.Printf("Linear solve: x = -b / a = -(%s) / %s\n", formatter.Number(s.C), formatter.Number(s.B))
	case solver.DoubleRoot, solver.TwoRealRoots, solver.TwoComplexRoots:
		if s.Scale != 1 {
			fmt.Printf("b^2 and 4ac are too small for a double: a, b and c are divided by %s, which keeps the roots\n",
				formatter.Number(s.Scale))
		}
		a, b, c := scaled(s)
		fmt.Printf("Discriminant Δ = b^2 - 4ac = (%s)^2 - 4 * (%s) * (%s) = %s\n",
			b, a, c, formatter.Number(s.Delta))
	}
}

// AfterCase shows the formulas of the case at hand only, with the values put in.
func AfterCase(s solver.Solution) {
	a, b, _ := scaled(s)
	sqrtDelta := formatter.Number(s.SqrtDelta)
	switch s.Kind {
	case solver.DoubleRoot:
		fmt.Printf("x = -b / (2a) = -(%s) / (2 * %s)\n", b, a)
	case solver.TwoRealRoots:
		fmt.Printf("sqrt(Δ) = %s\n", sqrtDelta)
		fmt.Printf("x1 = (-b - sqrt(Δ)) / (2a) = (-(%s) - %s) / (2 * %s)\n", b, sqrtDelta, a)
		fmt.Printf("x2 = (-b + sqrt(Δ)) / (2a) = (-(%s) + %s) / (2 * %s)\n", b, sqrtDelta, a)
		fmt.Println("Computed as q / a and c / q, with q = -(b + sign(b) * sqrt(Δ)) / 2:" +
			" the same roots, without subtracting two close numbers")
	case solver.TwoComplexRoots:
		fmt.Printf("sqrt(-Δ) = %s\n", sqrtDelta)
		fmt.Printf("Real part = -b / (2a) = -(%s) / (2 * %s)\n", b, a)
		fmt.Printf("Imag part = sqrt(-Δ) / (2a) = %s / (2 * %s)\n", sqrtDelta, a)
	}
}

// scaled writes a, b and c as Δ was computed from them: divided by s.Scale.
func scaled(s solver.Solution) (string, string, string) {
	return formatter.Number(s.A / s.Scale), formatter.Number(s.B / s.Scale),
		formatter.Number(s.C / s.Scale)
}
