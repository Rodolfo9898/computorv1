package main

import (
	"fmt"
	"os"
	"strings"

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
		default:
			if strings.HasPrefix(os.Args[1], "-") {
				fmt.Printf("Unknown flag or command: %s\n\n", os.Args[1])
				printHelp()
				os.Exit(1)
			}
			mandatory.Run(os.Args[1:])
			return
		}
	}

	mandatory.Run(os.Args[1:])
}

func printHelp() {
	fmt.Println(colors.Wrap(colors.Grey, "Usage:"))
	fmt.Println(colors.Wrap(colors.Grey, "  ./computor_v1 \"<equation>\"          # mandatory strict mode"))
	fmt.Println(colors.Wrap(colors.Grey, "  ./computor_v1 --bonus \"<equation>\"   # bonus free-form mode"))
	fmt.Println(colors.Wrap(colors.Grey, "  ./computor_v1 --help                  # show this message"))
}
