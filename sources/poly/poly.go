package poly

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
