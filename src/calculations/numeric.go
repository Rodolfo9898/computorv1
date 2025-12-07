package calculations

import "math"

const Epsilon = 1e-9

func SqrtNewton(value float64) float64 {
	if value <= 0 {
		return 0
	}

	guess := value
	if guess < 1 {
		guess = 1
	}

	for i := 0; i < 20; i++ {
		guess = 0.5 * (guess + value/guess)
	}
	return guess
}

func AlmostZero(value float64) bool {
	return value > -Epsilon && value < Epsilon
}

func CleanZero(value float64) float64 {
	if AlmostZero(value) {
		return 0
	}
	return value
}

func Abs(value float64) float64 {
	return math.Abs(value)
}
