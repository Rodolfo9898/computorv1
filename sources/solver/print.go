package solver

import (
	"fmt"

	"computorv1/sources/formatter"
)

// Hooks are what the bonus adds to the output of Print; the mandatory part
// leaves them all nil.
type Hooks struct {
	Number         func(float64) string // writes a solution; formatter.Number when nil
	BeforeCase     func(Solution)       // before the line that names the case
	AfterCase      func(Solution)       // after that line, before the solutions
	AfterSolutions func(Solution)       // after the solutions
}

// Print writes the case of the solution and its solutions, as in the subject.
func Print(s Solution, hooks Hooks) {
	number := hooks.Number
	if number == nil {
		number = formatter.Number
	}

	call(hooks.BeforeCase, s)
	switch s.Kind {
	case AnyReal:
		fmt.Println("Any real number is a solution.")
	case NoSolution:
		fmt.Println("No solution.")
	case DegreeTooHigh:
		fmt.Println("The polynomial degree is strictly greater than 2, I can't solve.")
	case DeltaTooLarge:
		fmt.Println("The coefficients are too large: the discriminant does not fit in a double.")
	case RootTooLarge:
		fmt.Println("The coefficients are too far apart: a solution does not fit in a double.")
	case OneRoot:
		fmt.Println("The solution is:")
		call(hooks.AfterCase, s)
		fmt.Printf("x = %s\n", number(s.X1))
	case DoubleRoot:
		fmt.Println("Discriminant is zero, the solution is:")
		call(hooks.AfterCase, s)
		fmt.Printf("x = %s\n", number(s.X1))
	case TwoRealRoots:
		fmt.Println("Discriminant is strictly positive, the two solutions are:")
		call(hooks.AfterCase, s)
		fmt.Printf("x1 = %s\n", number(s.X1))
		fmt.Printf("x2 = %s\n", number(s.X2))
	case TwoComplexRoots:
		fmt.Println("Discriminant is strictly negative, the two complex solutions are:")
		call(hooks.AfterCase, s)
		fmt.Printf("x1 = %s\n", formatter.Complex(s.Re, s.Im, number))
		fmt.Printf("x2 = %s\n", formatter.Complex(s.Re, -s.Im, number))
	}
	call(hooks.AfterSolutions, s)
}

func call(hook func(Solution), s Solution) {
	if hook != nil {
		hook(s)
	}
}
