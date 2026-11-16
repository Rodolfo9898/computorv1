package bonus

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"computorv1/src/base"
	"computorv1/src/bonus/colors"
	"computorv1/src/bonus/freeform"
	"computorv1/src/bonus/validator"
	"computorv1/src/core/formatter"
	"computorv1/src/execution"
)

// Run solves the equation of the arguments with the chosen mode. Without an
// equation it reads them from the standard input, one per line (interactive mode).
func Run(args []string) {
	mode, eqArgs := parseBonusArgs(args)
	if len(eqArgs) == 0 {
		interactive(mode)
		return
	}

	if !solve(mode, strings.Join(eqArgs, " ")) {
		PrintUsage()
		os.Exit(1)
	}
}

// solve validates, parses and solves one equation; false when it is refused.
func solve(mode base.BonusMode, equation string) bool {
	if mode.Validate {
		if err := validator.ValidateEquation(equation); err != nil {
			printValidationError(err)
			return false
		}
	}

	coeffs, err := freeform.ParseEquation(equation)
	if err != nil {
		fmt.Fprintln(os.Stderr, colors.Wrap(colors.Red, fmt.Sprintf("Parse error: %v", err)))
		return false
	}

	reduced := formatter.Reduced(coeffs)
	if mode.Natural {
		reduced = formatter.Natural(coeffs)
	}
	fmt.Printf("Reduced form: %s = 0\n", reduced)

	degree := execution.Degree(coeffs)
	fmt.Printf("Polynomial degree: %d\n", degree)

	execution.SolveWithOptions(coeffs, degree, base.SolveOptions{
		Verbose:      mode.ShowSteps,
		UseFractions: mode.UseFractions,
		Factored:     mode.Factored,
		Check:        mode.Check,
	})
	return true
}

// interactive reads one equation per line until the end of the input, or "quit",
// and solves each one. A terminal gets a prompt; otherwise each equation is
// written before its result, so that the output reads on its own. The exit status
// is 1 when an equation was refused.
func interactive(mode base.BonusMode) {
	prompt := isTerminal(os.Stdin)
	if prompt {
		fmt.Println(colors.Wrap(colors.Grey, "One equation per line; \"quit\" or Ctrl-D to stop."))
	}

	lines := bufio.NewScanner(os.Stdin)
	lines.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	refused := false
	count := 0
	for {
		if prompt {
			fmt.Print("> ")
		}
		if !lines.Scan() {
			break
		}
		line := strings.TrimSpace(lines.Text())
		if line == "" {
			continue
		}
		if line == "quit" || line == "exit" {
			break
		}
		if !prompt {
			if count > 0 {
				fmt.Println()
			}
			fmt.Printf("Equation: %s\n", line)
		}
		count++
		if !solve(mode, line) {
			refused = true
		}
	}
	if prompt {
		fmt.Println()
	}

	if err := lines.Err(); err != nil {
		fmt.Fprintln(os.Stderr, colors.Wrap(colors.Red, fmt.Sprintf("Error reading the standard input: %v", err)))
		os.Exit(1)
	}
	if refused {
		os.Exit(1)
	}
}

func isTerminal(file *os.File) bool {
	info, err := file.Stat()
	return err == nil && info.Mode()&os.ModeCharDevice != 0
}

func PrintUsage() {
	lines := []string{
		"\nUsage: ./computor_v1 --bonus [mode] \"<lhs> = <rhs>\"",
		"       ./computor_v1 --bonus [mode]          # interactive: one equation per line on stdin",
		"Modes (free form entry is always on):",
		" - all: every feature below [default]",
		" - steps: show the intermediate steps",
		" - fractions: show solutions as irreducible fractions, when they are exact",
		" - validation: check the characters and the '=' first, with coloured errors",
		" - natural: write the reduced form as by hand (5 + 4 * X)",
		" - factored: write the polynomial as a product of its factors",
		" - check: put each solution back into the reduced form",
		" - freeform: free form entry only",
		" - novalidation: every feature but validation",
		" - Accepts free-form terms: 'X', '4 * X', '4X', 'X^3', '-5', mixed order",
		"Examples:",
		"  ./computor_v1 --bonus \"5 + 4 * X + X^2 = X^2\"",
		"  ./computor_v1 --bonus steps \"X^2 + X = 0\"",
		"  ./computor_v1 --bonus fractions \"3 * X = 1\"",
		"  ./computor_v1 --bonus validation \"@ + X = 0\"",
		"  ./computor_v1 --bonus factored \"X^2 - 3 * X + 2 = 0\"",
	}
	for _, line := range lines {
		fmt.Fprintln(os.Stderr, colors.Wrap(colors.Grey, line))
	}
}

func parseBonusArgs(args []string) (base.BonusMode, []string) {
	all := base.BonusMode{ShowSteps: true, UseFractions: true, Validate: true,
		Natural: true, Factored: true, Check: true}
	if len(args) == 0 {
		return all, args
	}

	switch strings.ToLower(args[0]) {
	case "all":
		return all, args[1:]
	case "steps":
		return base.BonusMode{ShowSteps: true}, args[1:]
	case "fractions", "fraction":
		return base.BonusMode{UseFractions: true}, args[1:]
	case "validation", "validate":
		return base.BonusMode{Validate: true}, args[1:]
	case "natural":
		return base.BonusMode{Natural: true}, args[1:]
	case "factored", "factor":
		return base.BonusMode{Factored: true}, args[1:]
	case "check":
		return base.BonusMode{Check: true}, args[1:]
	case "freeform":
		return base.BonusMode{}, args[1:]
	case "novalidation", "no-validation":
		all.Validate = false
		return all, args[1:]
	default:
		// treat as equation, keep defaults (all bonuses)
		return all, args
	}
}

func printValidationError(err error) {
	color := colors.Red
	if ve, ok := err.(base.ValidationError); ok {
		switch ve.Type {
		case "chars":
			color = colors.Cyan
		case "format":
			color = colors.Yellow
		}
	}
	fmt.Fprintln(os.Stderr, colors.Wrap(color, err.Error()))
}
