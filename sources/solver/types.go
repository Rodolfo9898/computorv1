package solver

import "computorv1/sources/poly"

// Kind is the case an equation falls in.
type Kind int

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

// Hooks are what the bonus adds to the output of Print; the mandatory part
// leaves them all nil.
type Hooks struct {
	Number         func(float64) string // writes a solution; formatter.Number when nil
	BeforeCase     func(Solution)       // before the line that names the case
	AfterCase      func(Solution)       // after that line, before the solutions
	AfterSolutions func(Solution)       // after the solutions
}
