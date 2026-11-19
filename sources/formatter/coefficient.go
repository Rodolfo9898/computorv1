package formatter

import (
	"strconv"

	"computorv1/sources/numeric"
)

// Coefficient writes a coefficient of the reduced form with 10 significant digits,
// or more when 10 would write another number than the one solved: 1.00000000001
// stays 1.00000000001, not 1. noise is how far rounding may have moved the
// coefficient, so its digits are left out: 0.1 + 0.2 is 0.3, not 0.30000000000000004.
func Coefficient(value, noise float64) string {
	if value == 0 {
		return "0"
	}
	for digits := 10; digits < 17; digits++ {
		text := strconv.FormatFloat(value, 'g', digits, 64)
		read, err := strconv.ParseFloat(text, 64)
		if err == nil && numeric.Abs(read-value) <= noise {
			return text
		}
	}
	return strconv.FormatFloat(value, 'g', 17, 64)
}
