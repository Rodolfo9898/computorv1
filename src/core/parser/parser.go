package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"computorv1/src/base"
)

// termPattern matches polynomial terms in the form: coefficient * x ^ exponent (e.g., "3 * x ^ 2")
// Breakdown:
// - `^\s*` and `\s*$` allow optional whitespace at the start and end
// - `([+-]?\s*\d+(?:\.\d+)?)` captures the coefficient: optional sign, digits, optional decimal part
// - `\*` matches the multiplication operator
// - `[Xx]` matches the variable (case-insensitive)
// - `\^` matches the exponent operator
// - `([+-]?\d+)` captures the exponent: optional sign followed by digits
// The pattern is compiled once at startup for performance optimization.
var termPattern = regexp.MustCompile(`^\s*([+-]?\s*\d+(?:\.\d+)?)\s*\*\s*[Xx]\s*\^\s*([+-]?\d+)\s*$`)

// ParseEquation normalizes and subtracts right-hand terms, returning coefficients and max exponent seen.
func ParseEquation(input string) (base.PolyCoeffs, error) {
	parts := strings.Split(input, "=")
	if len(parts) != 2 {
		return nil, fmt.Errorf("equation must contain '='")
	}

	left, leftMax, err := parseSide(parts[0])
	if err != nil {
		return nil, err
	}
	right, rightMax, err := parseSide(parts[1])
	if err != nil {
		return nil, err
	}

	maxExp := leftMax
	if rightMax > maxExp {
		maxExp = rightMax
	}

	for exp, coef := range right {
		left[exp] -= coef
		if exp > maxExp {
			maxExp = exp
		}
	}

	return left, nil
}

func parseSide(expr string) (base.PolyCoeffs, int, error) {
	coeffs := make(base.PolyCoeffs)
	maxExp := 0

	normalized := strings.ReplaceAll(expr, "-", "+-")
	parts := strings.Split(normalized, "+")
	for _, raw := range parts {
		part := strings.TrimSpace(raw)
		if part == "" {
			continue
		}

		match := termPattern.FindStringSubmatch(part)
		if match == nil {
			numberVal, numberErr := strconv.ParseFloat(strings.ReplaceAll(part, " ", ""), 64)
			if numberErr != nil {
				return nil, 0, fmt.Errorf("invalid term %q", part)
			}
			coeffs[0] += numberVal
			continue
		}

		coeffStr := strings.ReplaceAll(match[1], " ", "")
		coeff, err := strconv.ParseFloat(coeffStr, 64)
		if err != nil {
			return nil, 0, fmt.Errorf("invalid coefficient in term %q", part)
		}

		exp, err := strconv.Atoi(strings.TrimSpace(match[2]))
		if err != nil {
			return nil, 0, fmt.Errorf("invalid exponent in term %q", part)
		}

		coeffs[exp] += coeff
		if exp > maxExp {
			maxExp = exp
		}
	}

	return coeffs, maxExp, nil
}
