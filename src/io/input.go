package input

import (
	"bufio"
	"os"
	"strings"
)

// Read grabs the equation either from provided args or STDIN.
func Read(args []string) string {
	if len(args) > 0 {
		return strings.Join(args, " ")
	}

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}
