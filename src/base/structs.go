package base

// ValidationError provides typed validation errors for input handling.
type ValidationError struct {
	Type    string
	Message string
}

// SolveOptions configures solver behavior.
type SolveOptions struct {
	Verbose      bool
	UseFractions bool
}

// BonusMode holds flags that control bonus behaviors.
type BonusMode struct {
	ShowSteps    bool
	UseFractions bool
	Validate     bool
}
