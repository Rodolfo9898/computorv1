//go:build bonus_mode

package freeform

import (
	"fmt"
	"strings"
)

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
