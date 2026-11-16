package numeric

// Unit is the unit roundoff of a double: a rounded result is within Unit times
// its size of the exact one.
const Unit = 0x1p-53

// Abs returns |value|: the subject forbids the math library beyond + - * /.
func Abs(value float64) float64 {
	if value < 0 {
		return -value
	}
	return value
}

// IsFinite is false for +Inf, -Inf and NaN, the results of an overflow.
func IsFinite(value float64) bool {
	return value-value == 0
}

// SumNoise bounds the rounding error of a sum of terms decimal numbers whose
// sizes add up to size: each number is rounded once when it is read, and each
// addition rounds once more. A sum no larger than this is a cancellation
// (0.1 + 0.2 - 0.3 leaves 5.5e-17), and 1 - 1.0000000000001 is not one.
func SumNoise(terms int, size float64) float64 {
	return 2 * float64(terms) * Unit * size
}

// Discriminant is b² - 4ac, made exactly 0 when it is no larger than what rounding
// may have put there: the noise of a, b and c carried through the formula, then
// the rounding of its products and of the subtraction. ok is false when b² or
// 4ac does not fit in a double.
func Discriminant(a, b, c, noiseA, noiseB, noiseC float64) (float64, bool) {
	bb, ac4 := b*b, 4*a*c
	delta := bb - ac4
	noise := 2*Abs(b)*noiseB + noiseB*noiseB +
		4*(Abs(a)*noiseC+Abs(c)*noiseA+noiseA*noiseC) +
		3*Unit*(bb+Abs(ac4))
	if !IsFinite(delta) || !IsFinite(noise) {
		return 0, false
	}
	if Abs(delta) <= noise {
		return 0, true
	}
	return delta, true
}

// SqrtNewton is Newton's method, g = (g + v / g) / 2, started above the root: every
// step then goes down, until rounding stops it, which is the closest double.
func SqrtNewton(value float64) float64 {
	if value <= 0 {
		return 0
	}
	if !IsFinite(value) { // +Inf or NaN: no step would ever end
		return value
	}

	guess := value
	if guess < 1 {
		guess = 1
	}
	for {
		next := 0.5 * (guess + value/guess)
		if next >= guess {
			return guess
		}
		guess = next
	}
}
