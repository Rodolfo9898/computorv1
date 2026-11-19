package solver

import (
	"fmt"

	"computorv1/sources/formatter"
	"computorv1/sources/numeric"
	"computorv1/sources/poly"
)

// Solve finds the solutions of a reduced equation of degree 2 or less.
func Solve(eq poly.Equation) Solution {
	coeffs := eq.Coeffs
	s := Solution{Coeffs: coeffs, A: coeffs[2], B: coeffs[1], C: coeffs[0], Scale: 1}
	switch coeffs.Degree() {
	case 0:
		s.Kind = NoSolution
		if s.C == 0 {
			s.Kind = AnyReal
		}
	case 1: // b * X + c = 0, b not zero
		s.Kind = OneRoot
		s.X1 = -s.C / s.B
	case 2:
		solveDegreeTwo(&s, eq.Noise)
	default:
		s.Kind = DegreeTooHigh
	}

	// A coefficient as small as 1e-310 next to one of 1 puts a solution past
	// the largest double: say so rather than print +Inf.
	for _, value := range []float64{s.X1, s.X2, s.Re, s.Im} {
		if !numeric.IsFinite(value) {
			s.Kind = RootTooLarge
		}
	}
	return s
}

// solveDegreeTwo solves a * X^2 + b * X + c = 0, a not zero. Coefficients so small
// that b² or 4ac would underflow are first divided by s.Scale, which keeps the roots.
func solveDegreeTwo(s *Solution, noise poly.PolyCoeffs) {
	s.Scale = numeric.ScaleFor(s.A, s.B, s.C)
	a, b, c := s.A/s.Scale, s.B/s.Scale, s.C/s.Scale
	delta, ok := numeric.Discriminant(a, b, c,
		noise[2]/s.Scale, noise[1]/s.Scale, noise[0]/s.Scale)
	if !ok {
		s.Kind = DeltaTooLarge
		return
	}
	s.Delta = delta

	switch {
	case delta == 0:
		s.Kind = DoubleRoot
		s.X1 = -b / (2 * a)
	case delta > 0:
		s.Kind = TwoRealRoots
		s.SqrtDelta = numeric.SqrtNewton(delta)
		s.X1, s.X2 = realRoots(a, b, c, s.SqrtDelta)
	default:
		s.Kind = TwoComplexRoots
		s.SqrtDelta = numeric.SqrtNewton(-delta)
		s.Re = -b / (2 * a)
		s.Im = s.SqrtDelta / (2 * a)
	}
}

// realRoots gives x1 = (-b - √Δ) / 2a and x2 = (-b + √Δ) / 2a without subtracting
// two close numbers: when b² is much larger than 4ac, -b + √Δ (b > 0) would lose
// its digits. That root comes from the other one instead, as c / q, since
// x1 * x2 = c / a.
func realRoots(a, b, c, sqrtDisc float64) (float64, float64) {
	if b >= 0 {
		q := -(b + sqrtDisc) / 2
		return q / a, c / q
	}
	q := (sqrtDisc - b) / 2
	return c / q, q / a
}

// Print writes the case of the solution and its solutions, as in the subject.
func Print(s Solution, hooks Hooks) {
	number := hooks.Number
	if number == nil {
		number = formatter.Number
	}

	call(hooks.BeforeCase, s)
	switch s.Kind {
	case AnyReal:
		fmt.Println("Any real number is a solution.")
	case NoSolution:
		fmt.Println("No solution.")
	case DegreeTooHigh:
		fmt.Println("The polynomial degree is strictly greater than 2, I can't solve.")
	case DeltaTooLarge:
		fmt.Println("The coefficients are too large: the discriminant does not fit in a double.")
	case RootTooLarge:
		fmt.Println("The coefficients are too far apart: a solution does not fit in a double.")
	case OneRoot:
		fmt.Println("The solution is:")
		call(hooks.AfterCase, s)
		fmt.Printf("x = %s\n", number(s.X1))
	case DoubleRoot:
		fmt.Println("Discriminant is zero, the solution is:")
		call(hooks.AfterCase, s)
		fmt.Printf("x = %s\n", number(s.X1))
	case TwoRealRoots:
		fmt.Println("Discriminant is strictly positive, the two solutions are:")
		call(hooks.AfterCase, s)
		fmt.Printf("x1 = %s\n", number(s.X1))
		fmt.Printf("x2 = %s\n", number(s.X2))
	case TwoComplexRoots:
		fmt.Println("Discriminant is strictly negative, the two complex solutions are:")
		call(hooks.AfterCase, s)
		fmt.Printf("x1 = %s\n", formatter.Complex(s.Re, s.Im, number))
		fmt.Printf("x2 = %s\n", formatter.Complex(s.Re, -s.Im, number))
	}
	call(hooks.AfterSolutions, s)
}

func call(hook func(Solution), s Solution) {
	if hook != nil {
		hook(s)
	}
}
