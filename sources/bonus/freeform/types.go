//go:build bonus_mode

package freeform

// scanner walks one side of the equation, byte by byte.
type scanner struct {
	text string
	pos  int
}
