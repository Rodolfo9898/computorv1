//go:build bonus_mode

package mistakes

// Bonus feature 1, entry mistakes (vocabulary and syntax). Check reads the
// equation before the parser: the vocabulary (only numbers, X, + - * ^ = and
// spaces), one '=' and two sides that are not empty. The syntax mistakes the
// free form parser finds (1 ++ 2, X^, 2 * 3 * X...) are reported by Print too.

import (
	"fmt"
	"strings"
	"unicode"
)

func (e Error) Error() string {
	return e.Message
}

// allowed is the vocabulary of an equation, one character at a time.
func allowed(c rune) bool {
	return (c >= '0' && c <= '9') || c == 'x' || c == 'X' || c == '.' ||
		strings.ContainsRune("+-*^= \t", c)
}

// Check gives the first vocabulary or structure mistake of input, or nil.
func Check(input string) error {
	if strings.TrimSpace(input) == "" {
		return Error{Kind: "structure", Message: "empty equation"}
	}
	if strings.Count(input, "=") != 1 {
		return Error{Kind: "structure", Message: "equation must contain exactly one '='"}
	}
	if err := checkVocabulary(input); err != nil {
		return err
	}
	parts := strings.SplitN(input, "=", 2)
	if strings.TrimSpace(parts[0]) == "" || strings.TrimSpace(parts[1]) == "" {
		return Error{Kind: "structure", Message: "equation cannot have empty sides"}
	}
	return nil
}

// checkVocabulary names the first thing that is not in the vocabulary: a whole
// word ("hello", "Xavier") when the character is in one, the character alone
// otherwise ('@').
func checkVocabulary(input string) error {
	runes := []rune(input)
	for i, c := range runes {
		if allowed(c) {
			continue
		}
		start, end := i, i+1
		for start > 0 && unicode.IsLetter(runes[start-1]) {
			start--
		}
		for end < len(runes) && unicode.IsLetter(runes[end]) {
			end++
		}
		if end-start > 1 && unicode.IsLetter(c) {
			return Error{Kind: "vocabulary", Message: fmt.Sprintf(
				"unknown word %q: an equation has numbers, X, + - * ^ = and spaces only",
				string(runes[start:end]))}
		}
		return Error{Kind: "vocabulary", Message: fmt.Sprintf("invalid character '%c' in equation", c)}
	}
	return nil
}
