//go:build bonus_mode

package main

import (
	"strings"

	"computorv1/sources/bonus/check"
	"computorv1/sources/bonus/factored"
	"computorv1/sources/bonus/fractions"
	"computorv1/sources/bonus/steps"
	"computorv1/sources/formatter"
	"computorv1/sources/solver"
)

// parseMode reads the mode in the first argument, if there is one: all the
// features when there is none.
func parseMode(args []string) (mode, []string) {
	all := mode{Mistakes: true, Fractions: true, Steps: true, Factored: true, Check: true}
	if len(args) == 0 {
		return all, args
	}
	switch strings.ToLower(args[0]) {
	case "all":
		return all, args[1:]
	case "freeform":
		return mode{}, args[1:]
	case "mistakes", "validation":
		return mode{Mistakes: true}, args[1:]
	case "fractions":
		return mode{Fractions: true}, args[1:]
	case "steps":
		return mode{Steps: true}, args[1:]
	case "factored":
		return mode{Factored: true}, args[1:]
	case "check":
		return mode{Check: true}, args[1:]
	}
	return all, args // no mode: the first argument is the equation
}

// hooks plugs the features of the mode into the output of the shared solver.
func (m mode) hooks() solver.Hooks {
	number := formatter.Number
	if m.Fractions {
		number = fractions.Number
	}
	hooks := solver.Hooks{Number: number}
	if m.Steps {
		hooks.BeforeCase = steps.BeforeCase
		hooks.AfterCase = steps.AfterCase
	}
	hooks.AfterSolutions = func(s solver.Solution) {
		if m.Factored {
			factored.Show(s, number)
		}
		if m.Check {
			check.Show(s, number)
		}
	}
	return hooks
}
