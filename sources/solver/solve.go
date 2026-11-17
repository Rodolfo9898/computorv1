package solver

import (
	"computorv1/sources/numeric"
	"computorv1/sources/poly"
)

// Kind is the case an equation falls in.
type Kind int

const (
	AnyReal         Kind = iota // every term cancelled out: 0 = 0
	NoSolution                  // a constant that is not zero: c = 0
	DegreeTooHigh               // degree 3 or more: not solved
	DeltaTooLarge               // b² or 4ac does not fit in a double
	RootTooLarge                // a solution does not fit in a double
	OneRoot                     // degree 1
	DoubleRoot                  // degree 2, Δ = 0
	TwoRealRoots                // degree 2, Δ > 0
	TwoComplexRoots             // degree 2, Δ < 0
)

// Solution is what Solve found. Solve prints nothing: Print writes it, and the
// bonus adds its own lines through Hooks, so that the mandatory part does not
// depend on the bonus.
type Solution struct {
	Kind      Kind
	Coeffs    poly.PolyCoeffs
	A, B, C   float64 // the coefficients of X^2, X^1 and X^0
	Delta     float64 // b² - 4ac, exactly 0 when it is within rounding noise
	SqrtDelta float64 // √|Δ|
	X1, X2    float64 // the real solutions (X1 alone for OneRoot and DoubleRoot)
	Re, Im    float64 // the complex solutions are Re + Im i and Re - Im i
}

// Solve finds the solutions of a reduced equation of degree 2 or less.
func Solve(eq poly.Equation) Solution {
	coeffs := eq.Coeffs
	s := Solution{Coeffs: coeffs, A: coeffs[2], B: coeffs[1], C: coeffs[0]}
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

// solveDegreeTwo solves a * X^2 + b * X + c = 0, a not zero.
func solveDegreeTwo(s *Solution, noise poly.PolyCoeffs) {
	delta, ok := numeric.Discriminant(s.A, s.B, s.C, noise[2], noise[1], noise[0])
	if !ok {
		s.Kind = DeltaTooLarge
		return
	}
	s.Delta = delta

	switch {
	case delta == 0:
		s.Kind = DoubleRoot
		s.X1 = -s.B / (2 * s.A)
	case delta > 0:
		s.Kind = TwoRealRoots
		s.SqrtDelta = numeric.SqrtNewton(delta)
		s.X1, s.X2 = realRoots(s.A, s.B, s.C, s.SqrtDelta)
	default:
		s.Kind = TwoComplexRoots
		s.SqrtDelta = numeric.SqrtNewton(-delta)
		s.Re = -s.B / (2 * s.A)
		s.Im = s.SqrtDelta / (2 * s.A)
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
