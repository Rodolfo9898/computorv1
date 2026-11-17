package main

import "regexp"

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
