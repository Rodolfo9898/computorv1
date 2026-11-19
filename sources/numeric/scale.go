package numeric

// ScaleFor is what a * X^2 + b * X + c = 0 must be divided by so that b² and 4ac
// fit in a double: 1 when they already do. With a and c of 1e-200, 4ac (4e-400)
// would be 0, and X^2 + 1 = 0 would get Δ = 0 instead of -4. The scale is a power
// of two, so dividing by it is exact, and the roots stay the same.
func ScaleFor(a, b, c float64) float64 {
	if !underflows(a, b, c) {
		return 1
	}
	return powerOfTwoAbove(max(Abs(a), Abs(b), Abs(c)))
}

// underflows tells when b² or 4ac, of coefficients that are not zero, falls
// below the smallest normal double, where it loses its digits or becomes 0.
func underflows(a, b, c float64) bool {
	return (b != 0 && b*b < MinNormal) || (a != 0 && c != 0 && Abs(4*a*c) < MinNormal)
}

// powerOfTwoAbove is the smallest power of two not below value (finite, > 0).
func powerOfTwoAbove(value float64) float64 {
	scale := 1.0
	for scale < value {
		scale *= 2
	}
	for scale/2 >= value {
		scale /= 2
	}
	return scale
}
