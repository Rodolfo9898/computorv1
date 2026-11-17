//go:build bonus_mode

package check

// relTol: P(x) is zero when it is this small next to the size of its terms. A
// solution is rounded to a double, so P(x) is rarely exactly 0.
const relTol = 1e-9
