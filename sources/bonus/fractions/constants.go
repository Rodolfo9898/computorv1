//go:build bonus_mode

package fractions

import "computorv1/sources/numeric"

const (
	// maxDen is the largest denominator tried: a solution of an equation with
	// a few decimals has a small one.
	maxDen = 1000000
	// maxValue keeps the numerators within an int.
	maxValue = 1e12
	// relTol: a fraction is the value when they differ by the rounding of the
	// computation only, about 7e-15 (64 unit roundoffs): 0.6 + 1.1 * X + 0.5 * X^2
	// gives -6/5 within 5 ulps, as Δ = 1.21 - 1.2 loses digits. A decimal written
	// with 14 digits or fewer, like 0.3333333333333 (1e-13 from 1/3), is farther.
	relTol = 64 * numeric.Unit
)
