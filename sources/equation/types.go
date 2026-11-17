package equation

// Term is one term of a side, as it was read: "4 * X^1" is {Exp: 1, Coef: 4}.
type Term struct {
	Exp  int
	Coef float64
}
