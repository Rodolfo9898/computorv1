package freeform

import (
	"fmt"
	"strconv"
	"strings"

	"computorv1/src/base"
	"computorv1/src/calculations"
)

// ParseEquation reads an equation in free form: "5 + 4 * X + X^2 = X^2", "2X = 4",
// "x^2 - 1 = 0". It also reads the strict form of the mandatory part, which is
// one of its cases. Every side is a sum of terms: [sign] [number] [*] [X[^n]].
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

// scanner walks one side of the equation, byte by byte.
type scanner struct {
	text string
	pos  int
}

func (s *scanner) skipSpaces() {
	for s.pos < len(s.text) && (s.text[s.pos] == ' ' || s.text[s.pos] == '\t') {
		s.pos++
	}
}

func (s *scanner) done() bool {
	s.skipSpaces()
	return s.pos >= len(s.text)
}

// peek is the next byte that is not a space, or 0 at the end.
func (s *scanner) peek() byte {
	if s.done() {
		return 0
	}
	return s.text[s.pos]
}

// near shows where a mistake is: the rest of the side from there.
func (s *scanner) near() string {
	if s.done() {
		return "the end of the side"
	}
	return fmt.Sprintf("%q", strings.TrimSpace(s.text[s.pos:]))
}

// digits reads [0-9]+ and returns them, or "" when there are none.
func (s *scanner) digits() string {
	start := s.pos
	for s.pos < len(s.text) && s.text[s.pos] >= '0' && s.text[s.pos] <= '9' {
		s.pos++
	}
	return s.text[start:s.pos]
}

func parseSide(expr string) (base.PolyCoeffs, error) {
	s := &scanner{text: expr}
	if s.done() {
		return nil, fmt.Errorf("a side of the equation is empty")
	}

	coeffs := make(base.PolyCoeffs)
	for first := true; !s.done(); first = false {
		sign := 1.0
		switch s.peek() {
		case '+':
			s.pos++
		case '-':
			sign = -1
			s.pos++
		default:
			if !first {
				return nil, fmt.Errorf("missing '+' or '-' before %s", s.near())
			}
		}

		coef, exp, err := parseTerm(s)
		if err != nil {
			return nil, err
		}
		coeffs[exp] = calculations.Sum(coeffs[exp], sign*coef)
	}
	return coeffs, nil
}

// parseTerm reads a term after its sign: "4", "4 * X", "4X", "X^2", "2.5 * x^3".
func parseTerm(s *scanner) (float64, int, error) {
	coef, hasCoef, err := parseNumber(s)
	if err != nil {
		return 0, 0, err
	}

	if s.peek() == '*' {
		if !hasCoef {
			return 0, 0, fmt.Errorf("'*' with no number before it, near %s", s.near())
		}
		s.pos++
		if c := s.peek(); c != 'X' && c != 'x' {
			return 0, 0, fmt.Errorf("expected X after '*', near %s", s.near())
		}
	}

	if c := s.peek(); c != 'X' && c != 'x' {
		if !hasCoef {
			return 0, 0, fmt.Errorf("expected a number or X, near %s", s.near())
		}
		return coef, 0, nil
	}
	s.pos++

	exp := 1
	if s.peek() == '^' {
		s.pos++
		s.skipSpaces()
		if s.pos < len(s.text) && s.text[s.pos] == '-' {
			return 0, 0, fmt.Errorf("negative exponent near %s: a polynomial has none", s.near())
		}
		digits := s.digits()
		if digits == "" {
			return 0, 0, fmt.Errorf("missing exponent after '^', near %s", s.near())
		}
		if s.pos < len(s.text) && s.text[s.pos] == '.' {
			return 0, 0, fmt.Errorf("the exponent %s... must be a whole number", digits)
		}
		exp, err = strconv.Atoi(digits)
		if err != nil {
			return 0, 0, fmt.Errorf("exponent %s is too large", digits)
		}
	}
	return coef, exp, nil
}

// parseNumber reads digits[.digits], if there is a number there; the coefficient
// of a term without one is 1.
func parseNumber(s *scanner) (float64, bool, error) {
	s.skipSpaces()
	whole := s.digits()
	if whole == "" {
		if s.pos < len(s.text) && s.text[s.pos] == '.' {
			return 0, false, fmt.Errorf("a number must start with a digit, near %s", s.near())
		}
		return 1, false, nil
	}

	text := whole
	if s.pos < len(s.text) && s.text[s.pos] == '.' {
		s.pos++
		decimals := s.digits()
		if decimals == "" {
			return 0, false, fmt.Errorf("missing digits after '%s.'", whole)
		}
		text += "." + decimals
	}

	value, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return 0, false, fmt.Errorf("number %s is too large", text)
	}
	return value, true, nil
}
