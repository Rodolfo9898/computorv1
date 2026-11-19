package input

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// Read grabs the equation either from its only argument or STDIN (its first line).
// Several arguments are refused rather than joined: an equation left without
// quotes is cut by the shell, which also replaces a * by the names of files.
func Read(args []string) (string, error) {
	if len(args) > 1 {
		return "", fmt.Errorf("%d arguments: give the equation as a single one, between quotes", len(args))
	}
	if len(args) == 1 {
		return args[0], nil
	}

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil && err != io.EOF { // EOF: the last line has no newline
		return "", fmt.Errorf("cannot read the standard input: %v", err)
	}
	text = strings.TrimSpace(text)
	if text == "" {
		return "", errors.New("no equation on the standard input")
	}
	return text, nil
}

// IsHelp tells the arguments that ask for the usage.
func IsHelp(arg string) bool {
	return arg == "-h" || arg == "--help" || arg == "help"
}

// IsFlag tells a flag (-x, --name) from an equation that starts with a minus sign,
// like "-2.5 * X^0 = 0": a flag is letters and dashes only.
func IsFlag(arg string) bool {
	if len(arg) < 2 || arg[0] != '-' {
		return false
	}
	for _, c := range arg[1:] {
		if c != '-' && (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') {
			return false
		}
	}
	return true
}
