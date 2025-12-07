package bonus

import (
	"fmt"
	"strings"

	"computorv1/src/base"
	"computorv1/src/bonus/colors"
	"computorv1/src/bonus/freeform"
	"computorv1/src/bonus/validator"
	"computorv1/src/core/formatter"
	"computorv1/src/execution"
	input "computorv1/src/io"
)

func Run(args []string) {
	mode, eqArgs := parseBonusArgs(args)
	if len(eqArgs) == 0 {
		PrintUsage()
		return
	}

	argument := input.Read(eqArgs)

	if mode.Validate {
		if err := validator.ValidateEquation(argument); err != nil {
			printValidationError(err)
			PrintUsage()
			return
		}
	}

	coeffs, err := freeform.ParseEquation(argument)
	if err != nil {
		fmt.Println(colors.Wrap(colors.Red, fmt.Sprintf("Parse error: %v", err)))
		PrintUsage()
		return
	}

	fmt.Printf("Reduced form: %s = 0\n", formatter.Reduced(coeffs))

	degree := execution.Degree(coeffs)
	fmt.Printf("Polynomial degree: %d\n", degree)

	execution.SolveWithOptions(coeffs, degree, base.SolveOptions{
		Verbose:      mode.ShowSteps,
		UseFractions: mode.UseFractions,
	})
}

func PrintUsage() {
	fmt.Println(colors.Wrap(colors.Grey, "\nUsage: ./computor_v1 --bonus [all|steps|fractions|validation|freeform|novalidation] \"<lhs> = <rhs>\""))
	fmt.Println(colors.Wrap(colors.Grey, " - all: enable all bonus displays (steps + fractions + validation) [default]"))
	fmt.Println(colors.Wrap(colors.Grey, " - steps: show intermediate solving steps (no fractions, no validation)"))
	fmt.Println(colors.Wrap(colors.Grey, " - fractions: show solutions as irreducible fractions (no steps, no validation)"))
	fmt.Println(colors.Wrap(colors.Grey, " - validation: run strict input validation only"))
	fmt.Println(colors.Wrap(colors.Grey, " - freeform: free-form parsing only (no steps, decimal output, no validation)"))
	fmt.Println(colors.Wrap(colors.Grey, " - novalidation: skip validation (keep steps/fractions defaults unless altered)"))
	fmt.Println(colors.Wrap(colors.Grey, " - Accepts free-form terms: 'X', '4 * X', 'X^3', '-5', mixed order"))
	fmt.Println(colors.Wrap(colors.Grey, "Examples:"))
	fmt.Println(colors.Wrap(colors.Grey, "  ./computor_v1 --bonus all \"5 + 4 * X + X^2 = X^2\""))
	fmt.Println(colors.Wrap(colors.Grey, "  ./computor_v1 --bonus steps \"X^2 + X = 0\""))
	fmt.Println(colors.Wrap(colors.Grey, "  ./computor_v1 --bonus fractions \"7 = 7\""))
	fmt.Println(colors.Wrap(colors.Grey, "  ./computor_v1 --bonus validation \"@ + X = 0\""))
	fmt.Println(colors.Wrap(colors.Grey, "  ./computor_v1 --bonus freeform \" 4 = 0\""))
}

func parseBonusArgs(args []string) (base.BonusMode, []string) {
	mode := base.BonusMode{ShowSteps: true, UseFractions: true, Validate: true}
	if len(args) == 0 {
		return mode, args
	}

	first := strings.ToLower(args[0])
	switch first {
	case "all":
		return mode, args[1:]
	case "steps":
		mode.ShowSteps = true
		mode.UseFractions = false
		mode.Validate = false
		return mode, args[1:]
	case "fractions", "fraction":
		mode.ShowSteps = false
		mode.UseFractions = true
		mode.Validate = false
		return mode, args[1:]
	case "validation", "validate":
		mode.ShowSteps = false
		mode.UseFractions = false
		mode.Validate = true
		return mode, args[1:]
	case "novalidation", "no-validation":
		mode.Validate = false
		return mode, args[1:]
	case "freeform":
		mode.ShowSteps = false
		mode.UseFractions = false
		mode.Validate = false
		return mode, args[1:]
	default:
		// treat as equation, keep defaults (all bonuses)
		return mode, args
	}
}

func printValidationError(err error) {
	if ve, ok := err.(base.ValidationError); ok {
		switch ve.Type {
		case "syntax":
			fmt.Println(colors.Wrap(colors.Red, ve.Message))
		case "chars":
			fmt.Println(colors.Wrap(colors.Cyan, ve.Message))
		case "format":
			fmt.Println(colors.Wrap(colors.Yellow, ve.Message))
		default:
			fmt.Println(colors.Wrap(colors.Red, ve.Message))
		}
		return
	}
	fmt.Println(colors.Wrap(colors.Red, err.Error()))
}
