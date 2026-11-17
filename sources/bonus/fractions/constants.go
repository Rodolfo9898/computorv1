//go:build bonus_mode

package fractions

const (
	// maxDen is the largest denominator tried: a solution of an equation with
	// a few decimals has a small one.
	maxDen = 1000000
	// maxValue keeps the numerators within an int.
	maxValue = 1e12
	// relTol: a fraction is the value when they differ by rounding only.
	relTol = 1e-13
)
