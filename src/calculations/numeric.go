package calculations

// RelEpsilon is how small a sum must be, next to the numbers added, to be only
// rounding noise: 0.1 + 0.2 - 0.3 leaves 5.5e-17, which is not a term.
const RelEpsilon = 1e-12

// Abs returns |value|: the subject forbids the math library beyond + - * /.
func Abs(value float64) float64 {
	if value < 0 {
		return -value
	}
	return value
}

// Sum adds a and b, and gives exactly 0 when what is left is rounding noise. A
// small number on its own (1e-10 * X^2) is kept: only a cancellation is cleared.
// An overflow (+Inf, or NaN) is kept too, for the caller to see it.
func Sum(a, b float64) float64 {
	sum := a + b
	if sum-sum != 0 {
		return sum
	}
	if Abs(sum) <= RelEpsilon*(Abs(a)+Abs(b)) {
		return 0
	}
	return sum
}

// SqrtNewton is Newton's method, g = (g + v / g) / 2, started above the root: every
// step then goes down, until rounding stops it, which is the closest double.
func SqrtNewton(value float64) float64 {
	if value <= 0 {
		return 0
	}
	if value-value != 0 { // +Inf or NaN: no step would ever end
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
