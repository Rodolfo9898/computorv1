//go:build bonus_mode

package freeform

// Free form entry, graded on its own (subject p.8): this parser, and the reduced
// form written as by hand (natural.go).

import (
	"fmt"
	"strconv"

	"computorv1/sources/equation"
	"computorv1/sources/poly"
)

// ParseEquation reads an equation in free form: "5 + 4 * X + X^2 = X^2", "2X = 4",
// "x^2 - 1 = 0". It also reads the strict form of the mandatory part, which is
// one of its cases. Every side is a sum of terms: [sign] [number] [*] [X[^n]].
func ParseEquation(input string) (poly.Equation, error) {
	return equation.Parse(input, parseSide)
}

func parseSide(expr string) ([]equation.Term, error) {
	s := &scanner{text: expr}
	if s.done() {
		return nil, fmt.Errorf("a side of the equation is empty")
	}

	var terms []equation.Term
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

		coef, exp, hasX, err := parseTerm(s)
		if err != nil {
			return nil, err
		}
		if coef == 0 && !hasX { // a lone 0 is the right side of "... = 0", not a term
			continue
		}
		terms = append(terms, equation.Term{Exp: exp, Coef: sign * coef})
	}
	return terms, nil
}

// parseTerm reads a term after its sign: "4", "4 * X", "4X", "X^2", "2.5 * x^3".
// hasX is false for a constant written without X.
func parseTerm(s *scanner) (coef float64, exp int, hasX bool, err error) {
	coef, hasCoef, err := parseNumber(s)
	if err != nil {
		return 0, 0, false, err
	}

	if s.peek() == '*' {
		if !hasCoef {
			return 0, 0, false, fmt.Errorf("'*' with no number before it, near %s", s.near())
		}
		s.pos++
		if c := s.peek(); c != 'X' && c != 'x' {
			return 0, 0, false, fmt.Errorf("expected X after '*', near %s", s.near())
		}
	}

	if c := s.peek(); c != 'X' && c != 'x' {
		if !hasCoef {
			return 0, 0, false, fmt.Errorf("expected a number or X, near %s", s.near())
		}
		return coef, 0, false, nil
	}
	s.pos++

	exp = 1
	if s.peek() == '^' {
		s.pos++
		s.skipSpaces()
		if s.pos < len(s.text) && s.text[s.pos] == '-' {
			return 0, 0, false, fmt.Errorf("negative exponent near %s: a polynomial has none", s.near())
		}
		digits := s.digits()
		if digits == "" {
			return 0, 0, false, fmt.Errorf("missing exponent after '^', near %s", s.near())
		}
		if s.pos < len(s.text) && s.text[s.pos] == '.' {
			return 0, 0, false, fmt.Errorf("the exponent %s... must be a whole number", digits)
		}
		exp, err = strconv.Atoi(digits)
		if err != nil {
			return 0, 0, false, fmt.Errorf("exponent %s is too large", digits)
		}
	}
	return coef, exp, true, nil
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
