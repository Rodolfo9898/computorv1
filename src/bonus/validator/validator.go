package validator

import (
	"fmt"
	"regexp"
	"strings"

	"computorv1/src/base"
)

var allowedChars = regexp.MustCompile(`^[0-9xX\+\-\*\^\.\s=]+$`)

// ValidateEquation performs bonus-level validation to surface clearer errors.
func ValidateEquation(input string) error {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return base.ValidationError{Type: "syntax", Message: "empty equation"}
	}

	if strings.Count(input, "=") != 1 {
		return base.ValidationError{Type: "syntax", Message: "equation must contain exactly one '='"}
	}

	if !allowedChars.MatchString(input) {
		for _, ch := range input {
			if !allowedChars.MatchString(string(ch)) {
				return base.ValidationError{Type: "chars", Message: fmt.Sprintf("invalid character '%c' in equation", ch)}
			}
		}
		return base.ValidationError{Type: "chars", Message: "equation contains invalid characters"}
	}

	parts := strings.SplitN(input, "=", 2)
	if strings.TrimSpace(parts[0]) == "" || strings.TrimSpace(parts[1]) == "" {
		return base.ValidationError{Type: "format", Message: "equation cannot have empty sides"}
	}

	return nil
}
