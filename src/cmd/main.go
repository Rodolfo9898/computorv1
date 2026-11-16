package main

import (
	"fmt"
	"os"

	"computorv1/src/bonus/colors"
	bonuscmd "computorv1/src/cmd/bonus"
	mandatory "computorv1/src/cmd/mandatory"
)

func main() {
	if len(os.Args) > 1 {
		switch {
		case os.Args[1] == "--bonus" || os.Args[1] == "-bonus" || os.Args[1] == "-b":
			bonuscmd.Run(os.Args[2:])
			return
		case os.Args[1] == "-h" || os.Args[1] == "--help" || os.Args[1] == "help":
			printHelp()
			return
		case isFlag(os.Args[1]):
			fmt.Fprintf(os.Stderr, "Unknown flag or command: %s\n\n", os.Args[1])
			printHelp()
			os.Exit(1)
		}
	}

	mandatory.Run(os.Args[1:])
}

// isFlag tells a flag (-x, --name) from an equation that starts with a minus sign,
// like "-2.5 * X^0 = 0": a flag is letters and dashes only.
func isFlag(arg string) bool {
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

func printHelp() {
	fmt.Println(colors.Wrap(colors.Grey, "Usage:"))
	fmt.Println(colors.Wrap(colors.Grey, "  ./computor_v1 \"<equation>\"               # mandatory strict mode"))
	fmt.Println(colors.Wrap(colors.Grey, "  ./computor_v1                            # the same, the equation read on stdin"))
	fmt.Println(colors.Wrap(colors.Grey, "  ./computor_v1 --bonus [mode] \"<equation>\" # bonus free-form mode"))
	fmt.Println(colors.Wrap(colors.Grey, "  ./computor_v1 --bonus [mode]             # bonus interactive mode, one equation per line"))
	fmt.Println(colors.Wrap(colors.Grey, "  ./computor_v1 --help                     # show this message"))
}
