//go:build bonus_mode

package main

// mode is the set of the five features turned on; free form entry always is.
type mode struct {
	Mistakes  bool
	Fractions bool
	Steps     bool
	Factored  bool
	Check     bool
}
