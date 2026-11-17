package poly

// PolyCoeffs maps each exponent written in the equation to its coefficient.
type PolyCoeffs map[int]float64

// Equation is an equation reduced to P(X) = 0. Noise gives, for each exponent,
// how far rounding may have moved its coefficient: the decimals of the entry are
// rounded when they are read, and again at each addition. The solver needs it to
// tell a zero discriminant from rounding noise.
type Equation struct {
	Coeffs PolyCoeffs
	Noise  PolyCoeffs
}

// Degree is the highest exponent whose coefficient is not zero, or 0 when none is.
// The parsers turn a coefficient that cancelled out into exactly 0, so a small
// one (1e-10 * X^2) still counts.
func (coeffs PolyCoeffs) Degree() int {
	degree := 0
	for exp, coef := range coeffs {
		if coef != 0 && exp > degree {
			degree = exp
		}
	}
	return degree
}
