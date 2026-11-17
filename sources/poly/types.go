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
