//go:build bonus_mode

package mistakes

import (
	"fmt"
	"io"
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
