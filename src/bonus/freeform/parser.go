package freeform

import (
	"fmt"
	"strconv"
	"strings"

	"computorv1/src/base"
	"computorv1/src/core/parser"
)

// ParseEquation attempts the strict parser first, then falls back to free-form parsing.
func ParseEquation(input string) (base.PolyCoeffs, error) {
	coeffs, err := parser.ParseEquation(input)
	if err == nil {
		return coeffs, nil
	}
	return parseLoose(input)
}

func parseLoose(input string) (base.PolyCoeffs, error) {
	parts := strings.Split(input, "=")
	if len(parts) != 2 {
		return nil, fmt.Errorf("equation must contain '='")
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
		left[exp] -= coef
	}

	return left, nil
}

func parseSide(expr string) (base.PolyCoeffs, error) {
	coeffs := make(base.PolyCoeffs)
	for _, token := range splitTerms(expr) {
		if token == "" {
			continue
		}
		coef, exp, err := parseTerm(token)
		if err != nil {
			return nil, err
		}
		coeffs[exp] += coef
	}
	return coeffs, nil
}

func splitTerms(expr string) []string {
	trimmed := strings.ReplaceAll(expr, " ", "")
	var terms []string
	start := 0
	for i := 1; i < len(trimmed); i++ {
		if (trimmed[i] == '+' || trimmed[i] == '-') && trimmed[i-1] != '^' {
			terms = append(terms, trimmed[start:i])
			start = i
		}
	}
	terms = append(terms, trimmed[start:])
	return terms
}

func parseTerm(term string) (float64, int, error) {
	term = strings.TrimSpace(term)
	if term == "" {
		return 0, 0, fmt.Errorf("empty term")
	}

	xIndex := strings.IndexAny(term, "Xx")
	if xIndex == -1 {
		value, err := strconv.ParseFloat(term, 64)
		if err != nil {
			return 0, 0, fmt.Errorf("invalid constant term %q", term)
		}
		return value, 0, nil
	}

	coefStr := strings.TrimSuffix(term[:xIndex], "*")
	exp := 1
	if len(term) > xIndex+1 {
		if term[xIndex+1] == '^' {
			expPart := term[xIndex+2:]
			if expPart == "" {
				return 0, 0, fmt.Errorf("missing exponent in term %q", term)
			}
			expVal, err := strconv.Atoi(expPart)
			if err != nil {
				return 0, 0, fmt.Errorf("invalid exponent in term %q", term)
			}
			exp = expVal
		} else {
			return 0, 0, fmt.Errorf("invalid syntax near %q", term)
		}
	}

	coef := 1.0
	if coefStr == "-" {
		coef = -1
	} else if coefStr != "" && coefStr != "+" {
		val, err := strconv.ParseFloat(coefStr, 64)
		if err != nil {
			return 0, 0, fmt.Errorf("invalid coefficient in term %q", term)
		}
		coef = val
	}

	return coef, exp, nil
}
