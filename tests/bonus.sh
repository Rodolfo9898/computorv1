#!/bin/bash
# Bonus checks (./computor_v1 --bonus [mode] "<equation>"). Exit status: the number
# of failed checks.

source "$(dirname "$0")/setup.sh"

echo "### free form entry (p.8)"
MODE=(--bonus freeform)
gives "5 + 4 * X + X^2= X^2" "Polynomial degree: 1"
solves "5 + 4 * X + X^2= X^2" "-1.25"
solves "X^2 + X = 0" "0 -1"
solves "2X + 4 = 0" "-2"
solves "x^2 - 4 = 0" "2 -2"
solves "X^2 + 3 = 4 * X" "1 3"
solves "X = 2 * X" "0"
solves "X = 10" "10"
solves "5 * X^0 + 4 * X^1 = 4 * X^0" "-0.25"
gives "4 = 0" "No solution."
gives "7 = 7" "Any real number is a solution."
gives "-3 * X^3 + 2 * X^2 + 1 = 0" "The polynomial degree is strictly greater than 2, I can't solve."
gives "X^4 + X^2 = X^4" "Polynomial degree: 2"
solves "X^4 + X^2 = X^4" "0"

echo "### entry mistakes (vocabulary and syntax): refused in every mode"
refuses "two signs in a row" "1 ++ 2 = 0"
refuses "a '*' with nothing before it" "* X = 0"
refuses "a negative exponent (not a polynomial)" "X^-1 = 0"
refuses "two numbers in a row" "1 2 = 0"
refuses "an exponent left out" "X^ = 0"
refuses "a decimal exponent" "X^2.5 = 0"
refuses "a number after X" "X2 = 0"
refuses "a product of numbers" "2 * 3 * X = 0"
refuses "a sign with no term" "X - = 0"
MODE=(--bonus validation)
gives "@ + X = 0" "invalid character '@' in equation"
refuses "a character that is not in an equation" "@ + X = 0"
refuses "two '='" "X = 1 = 2"
refuses "an empty side" "= 0"
refuses "an empty equation" "  "

echo "### irreducible fractions"
MODE=(--bonus fractions)
gives "3 * X = 1" "x = 1/3"
gives "4 * X^2 - 1 = 0" "x1 = -1/2" "x2 = 1/2"
gives "1 * X^0 + 2 * X^1 + 5 * X^2 = 0" "x1 = -1/5 + 2/5i" "x2 = -1/5 - 2/5i"
gives "0.5 * X = 0.75" "x = 3/2"
# sqrt(2) has no fraction: an approximate one would pass for an exact solution.
lacks "X^2 - 2 = 0" '^x[12]? = .*/' "irrational solutions are not shown as fractions"
solves "X^2 - 2 = 0" "1.414213562 -1.414213562"

echo "### intermediate steps"
MODE=(--bonus steps)
gives "X^2 - 1 = 0" "Discriminant Δ = b^2 - 4ac = (0)^2 - 4 * (1) * (-1) = 4" "sqrt(Δ) = 2"
solves "X^2 - 1 = 0" "1 -1"
lacks "X^2 - 1 = 0" '^(Real|Imag) part' "no real or imaginary part for real solutions"
count_is "X^2 + 1 = 0" "Discriminant Δ" 1
gives "X^2 + 1 = 0" "Real part = -b / (2a) = -(0) / (2 * 1)" "Imag part = sqrt(-Δ) / (2a) = 2 / (2 * 1)"
gives "X^2 + 2 * X + 1 = 0" "x = -b / (2a) = -(2) / (2 * 1)"
gives "2 * X + 1 = 0" "Linear solve: x = -b / a = -(1) / 2"

echo "### natural reduced form"
MODE=(--bonus natural)
gives "5 + 4 * X + X^2= X^2" "Reduced form: 5 + 4 * X = 0"
gives "X^2 - X = 0" "Reduced form: -X + X^2 = 0"
gives "-1 = X" "Reduced form: -1 - X = 0"
gives "2.5 * X^2 = 10" "Reduced form: -10 + 2.5 * X^2 = 0"
gives "7 = 7" "Reduced form: 0 = 0"

echo "### factored form"
MODE=(--bonus factored)
gives "X^2 - 3 * X + 2 = 0" "Factored form: (X - 1) * (X - 2) = 0"
gives "-1 * X^2 + 4 = 0" "Factored form: -(X - 2) * (X + 2) = 0"
gives "X^2 + X = 0" "Factored form: (X + 1) * X = 0"
gives "X^2 + 2 * X + 1 = 0" "Factored form: (X + 1)^2 = 0"
gives "2 * X + 1 = 0" "Factored form: 2 * (X + 0.5) = 0"
gives "X - 3 = 0" "Factored form: X - 3 = 0"
gives "X^2 + 1 = 0" "Factored form: (X - (0 + 1i)) * (X - (0 - 1i)) = 0"
lacks "4 = 0" '^Factored form' "no factored form without a solution"

echo "### solution check"
MODE=(--bonus check)
gives "X = 10" "Check: P(10) = 0"
count_is "5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0" "Check: P(" 2
lacks "5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0" 'not zero' "both solutions give P(x) = 0"
gives "1 * X^0 + 2 * X^1 + 5 * X^2 = 0" "Check: P(-0.2 + 0.4i) = 0" "Check: P(-0.2 - 0.4i) = 0"
lacks "1 * X^2 + 1000000 * X^1 = 0" 'not zero' "large roots still give P(x) = 0"

echo "### everything at once (the default mode)"
MODE=(--bonus)
gives "5 + 4 * X + X^2= X^2" "Reduced form: 5 + 4 * X = 0" "x = -5/4" \
	"Factored form: 4 * (X + 5/4) = 0" "Check: P(-5/4) = 0"

echo "### interactive mode: one equation per line on the standard input"
MODE=(--bonus freeform)
run_stdin $'X = 2\n\nX^2 - 1 = 0\nquit\nX = 3\n'
got=$(printed_solutions)
! crashed && [ "$STATUS" -eq 0 ] && same_numbers "2 1 -1" "$got" \
	&& [ "$(printf '%s\n' "$OUTPUT" | grep -c '^Equation:')" -eq 2 ]
check $? "two equations solved, an empty line skipped, then quit (printed: $got)"
run_stdin $'X = 2\n1 ++ 2 = 0\nX = 4\n'
got=$(printed_solutions)
! crashed && [ "$STATUS" -eq 1 ] && same_numbers "2 4" "$got" \
	&& printf '%s\n' "$OUTPUT" | grep -q 'Parse error'
check $? "a refused line is reported, the next one still solved, status 1 (status $STATUS, printed: $got)"
run_stdin ""
! crashed && [ "$STATUS" -eq 0 ]
check $? "an empty input ends at once (status $STATUS)"

echo "### bonus: $FAILS failed check(s)"
exit "$FAILS"
