//go:build bonus_mode

package factored

// lead is the leading coefficient before the factors: nothing for 1, "-" for -1.
func lead(a float64, number func(float64) string) string {
	switch a {
	case 1:
		return ""
	case -1:
		return "-"
	}
	return number(a) + " * "
}

// factor is X - x, written X + |x| for a negative x and X for zero.
func factor(x float64, number func(float64) string) string {
	switch {
	case x == 0:
		return "X"
	case x < 0:
		return "X + " + number(-x)
	}
	return "X - " + number(x)
}

// parens puts a factor between parentheses, unless it is X alone.
func parens(x float64, number func(float64) string) string {
	if x == 0 {
		return "X"
	}
	return "(" + factor(x, number) + ")"
}
