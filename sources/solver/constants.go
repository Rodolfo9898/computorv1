package solver

const (
	AnyReal         Kind = iota // every term cancelled out: 0 = 0
	NoSolution                  // a constant that is not zero: c = 0
	DegreeTooHigh               // degree 3 or more: not solved
	DeltaTooLarge               // b² or 4ac does not fit in a double
	RootTooLarge                // a solution does not fit in a double
	OneRoot                     // degree 1
	DoubleRoot                  // degree 2, Δ = 0
	TwoRealRoots                // degree 2, Δ > 0
	TwoComplexRoots             // degree 2, Δ < 0
)
