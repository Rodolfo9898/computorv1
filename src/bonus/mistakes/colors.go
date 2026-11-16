//go:build bonus_mode

package mistakes

import (
	"fmt"
	"io"
)

// The colours of the messages: one per kind of mistake, so that the kind reads at
// a glance.
const (
	red    = "\033[31m"
	yellow = "\033[33m"
	cyan   = "\033[36m"
	reset  = "\033[0m"
)

// Print writes a refused equation to w, in the colour of its kind: cyan for the
// vocabulary, yellow for the structure, red for a syntax mistake of the parser.
func Print(w io.Writer, err error) {
	color := red
	if mistake, ok := err.(Error); ok {
		switch mistake.Kind {
		case "vocabulary":
			color = cyan
		case "structure":
			color = yellow
		}
	}
	fmt.Fprintln(w, color+err.Error()+reset)
}
