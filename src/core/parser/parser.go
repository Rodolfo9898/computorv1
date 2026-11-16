package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"computorv1/src/base"
	"computorv1/src/calculations"
)

// termPattern matches polynomial terms in the form: coefficient * x ^ exponent (e.g., "3 * x ^ 2")
// Breakdown:
// - `^` and `$` anchor the whole term (it is trimmed first)
// - `([+-]?\s*\d+(?:\.\d+)?)` captures the coefficient: optional sign, digits, optional decimal part
// - `\*` matches the multiplication operator
// - `[Xx]` matches the variable (case-insensitive)
// - `\^` matches the exponent operator
// - `(\d+)` captures the exponent: digits only, a polynomial has no negative exponent
// The pattern is compiled once at startup for performance optimization.
var termPattern = regexp.MustCompile(`^([+-]?\s*\d+(?:\.\d+)?)\s*\*\s*[Xx]\s*\^\s*(\d+)$`)

// numberPattern matches a constant on its own, like the 0 of "... = 0".
var numberPattern = regexp.MustCompile(`^[+-]?\s*\d+(?:\.\d+)?$`)

// ParseEquation subtracts the right side from the left, and returns the coefficients by exponent.
func ParseEquation(input string) (base.PolyCoeffs, error) {
	parts := strings.Split(input, "=")
	if len(parts) != 2 {
		return nil, fmt.Errorf("equation must contain exactly one '='")
	}

	left, err := parseSide(parts[0])
	if err != nil {
		return nil, err
	}
	right, err := parseSide(parts[1])
	if err != nil {
		return nil, err
	}

	for exp, coef := range right {
		left[exp] = calculations.Sum(left[exp], -coef)
	}
	return left, nil
}

func parseSide(expr string) (base.PolyCoeffs, error) {
	if strings.TrimSpace(expr) == "" {
		return nil, fmt.Errorf("a side of the equation is empty")
	}

	coeffs := make(base.PolyCoeffs)
	for _, raw := range splitTerms(expr) {
		part := strings.TrimSpace(raw)
		if part == "" { // the spaces before a leading sign
			continue
		}

		if numberPattern.MatchString(part) {
			value, err := parseNumber(part)
			if err != nil {
				return nil, fmt.Errorf("invalid number %q", part)
			}
			coeffs[0] = calculations.Sum(coeffs[0], value)
			continue
		}

		match := termPattern.FindStringSubmatch(part)
		if match == nil {
			return nil, fmt.Errorf("invalid term %q (expected a * X^p)", part)
		}
		coeff, err := parseNumber(match[1])
		if err != nil {
			return nil, fmt.Errorf("invalid coefficient in term %q", part)
		}
		exp, err := strconv.Atoi(match[2])
		if err != nil {
			return nil, fmt.Errorf("invalid exponent in term %q", part)
		}
		coeffs[exp] = calculations.Sum(coeffs[exp], coeff)
	}
	return coeffs, nil
}

// splitTerms cuts a side before each sign, except the sign of an exponent:
// "4 * X^1 - 9.3 * X^2" gives "4 * X^1 " and "- 9.3 * X^2". A sign that follows
// another one starts a term of its own, which no pattern matches: "1 ++ 2" is refused.
func splitTerms(expr string) []string {
	var terms []string
	start := 0
	var last byte // the last character seen that is not a space
	for i := 0; i < len(expr); i++ {
		c := expr[i]
		if (c == '+' || c == '-') && last != 0 && last != '^' {
			terms = append(terms, expr[start:i])
			start = i
		}
		if c != ' ' && c != '\t' {
			last = c
		}
	}
	return append(terms, expr[start:])
}

// parseNumber reads a number the patterns accepted: a sign, maybe spaces, digits.
func parseNumber(text string) (float64, error) {
	return strconv.ParseFloat(strings.Join(strings.Fields(text), ""), 64)
}
