#!/bin/bash
# Mandatory checks (./build/computor_v1 "<equation>"). Exit status: the number of failed checks.

source "$(dirname "$0")/setup.sh"

echo "### the examples of the subject (p.6)"
gives "5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0" \
	"Reduced form: 4 * X^0 + 4 * X^1 - 9.3 * X^2 = 0" "Polynomial degree: 2" \
	"Discriminant is strictly positive, the two solutions are:"
solves "5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0" "0.905239 -0.475131"
gives "5 * X^0 + 4 * X^1 = 4 * X^0" \
	"Reduced form: 1 * X^0 + 4 * X^1 = 0" "Polynomial degree: 1" "The solution is:"
solves "5 * X^0 + 4 * X^1 = 4 * X^0" "-0.25"
# The subject keeps the "0 * X^2" of this reduced form; leaving it out says the same.
matches "8 * X^0 - 6 * X^1 + 0 * X^2 - 5.6 * X^3 = 3 * X^0" \
	'^Reduced form: 5 \* X\^0 - 6 \* X\^1( \+ 0 \* X\^2)? - 5\.6 \* X\^3 = 0$' "its reduced form"
gives "8 * X^0 - 6 * X^1 + 0 * X^2 - 5.6 * X^3 = 3 * X^0" "Polynomial degree: 3" \
	"The polynomial degree is strictly greater than 2, I can't solve."
gives "6 * X^0 = 6 * X^0" "Reduced form: 0 * X^0 = 0" "Any real number is a solution."
gives "10 * X^0 = 15 * X^0" "Reduced form: -5 * X^0 = 0" "No solution."
gives "1 * X^0 + 2 * X^1 + 5 * X^2 = 0" \
	"Reduced form: 1 * X^0 + 2 * X^1 + 5 * X^2 = 0" "Polynomial degree: 2" \
	"Discriminant is strictly negative, the two complex solutions are:" \
	"x1 = -0.2 + 0.4i" "x2 = -0.2 - 0.4i"

echo "### zero, negative and non whole coefficients (p.7)"
gives "6 * X^0 + 11 * X^1 + 5 * X^2 = 1 * X^0 + 1 * X^1" \
	"Reduced form: 5 * X^0 + 10 * X^1 + 5 * X^2 = 0" "Discriminant is zero, the solution is:"
solves "6 * X^0 + 11 * X^1 + 5 * X^2 = 1 * X^0 + 1 * X^1" "-1"
gives "5 * X^0 + 3 * X^1 + 3 * X^2 = 1 * X^0 + 0 * X^1" \
	"Reduced form: 4 * X^0 + 3 * X^1 + 3 * X^2 = 0" \
	"Discriminant is strictly negative, the two complex solutions are:"
gives "1 * X^1 = 10 * X^0" "Reduced form: -10 * X^0 + 1 * X^1 = 0"
solves "1 * X^1 = 10 * X^0" "10"
gives "1 * X^2 - 30 * X^1 + 200 * X^0 = 0" "Reduced form: 200 * X^0 - 30 * X^1 + 1 * X^2 = 0"
solves "1 * X^2 - 30 * X^1 + 200 * X^0 = 0" "10 20"
gives "-0.5 * X^2 + 2.25 * X^1 - 1 * X^0 = 0" "Reduced form: -1 * X^0 + 2.25 * X^1 - 0.5 * X^2 = 0"
solves "-0.5 * X^2 + 2.25 * X^1 - 1 * X^0 = 0" "0.5 4"
solves "0.001 * X^2 - 1 * X^0 = 0" "31.6227766 -31.6227766"
gives "0 * X^2 + 3 * X^1 = 0 * X^0" "Polynomial degree: 1"
solves "0 * X^2 + 3 * X^1 = 0 * X^0" "0"
gives "-2.5 * X^0 = 0 * X^0" "Reduced form: -2.5 * X^0 = 0" "No solution."
gives "5 * X^0 = 15 * X^0" "Reduced form: -10 * X^0 = 0" "No solution."
gives "10000000000 * X^0 = 0" "Reduced form: 1e+10 * X^0 = 0"
# A small coefficient is a coefficient: only what cancels out is zero.
gives "0.0000000001 * X^2 + 1 * X^0 = 0" "Reduced form: 1 * X^0 + 1e-10 * X^2 = 0" \
	"Polynomial degree: 2" "Discriminant is strictly negative, the two complex solutions are:"
gives "0.1 * X^1 + 0.2 * X^1 = 0.3 * X^1 + 1 * X^0" "Reduced form: -1 * X^0 = 0" "No solution."

echo "### the degree is the one of the reduced form"
gives "1 * X^2 + 2 * X^1 = 1 * X^2 + 4 * X^0" "Polynomial degree: 1"
solves "1 * X^2 + 2 * X^1 = 1 * X^2 + 4 * X^0" "2"
gives "1 * X^3 + 1 * X^1 = 1 * X^3" "Polynomial degree: 1"
solves "1 * X^3 + 1 * X^1 = 1 * X^3" "0"

echo "### large coefficients and discriminants"
solves "1 * X^2 + 1000000 * X^1 = 0" "0 -1000000"
solves "1 * X^2 - 2000 * X^1 - 1 * X^0 = 0" "2000.0005 -0.0005"
# -b + sqrt(Δ) would subtract two numbers equal to 8 digits: -7.45e-09, not -1e-08.
solves "1 * X^2 + 100000000 * X^1 + 1 * X^0 = 0" "-100000000 -0.00000001"
e150="1$(printf '0%.0s' $(seq 150))"
e200="1$(printf '0%.0s' $(seq 200))"
solves "$e150 * X^2 + $e150 * X^1 = 0" "0 -1"
gives "$e200 * X^2 + $e200 * X^1 = 0" \
	"The coefficients are too large: the discriminant does not fit in a double."

echo "### the standard input, when there is no argument (p.7)"
solves_stdin "5 * X^0 + 4 * X^1 = 4 * X^0" "-0.25"
solves_stdin $'1 * X^2 - 1 * X^0 = 0\n' "1 -1"

echo "### malformed input: a message, no crash"
refuses "no '='" "5 * X^0 + 4 * X^1"
refuses "two '='" "1 * X^0 = 2 * X^0 = 3 * X^0"
refuses "words" "hello = world"
refuses "an empty argument" ""
refuses "an exponent too big for an int" "1 * X^99999999999999999999 = 0"
refuses "two signs in a row" "1 * X^0 ++ 2 * X^0 = 0"
refuses "two numbers in a row" "1 2 = 0"
refuses "an empty side" "= 1 * X^0"
refuses "a negative exponent" "1 * X^-1 = 0"
refuses "a sign with no term" "1 * X^0 - = 0"
refuses_stdin "an empty standard input" ""

echo "### a long equation: 8000 terms"
long="2 * X^1"
for i in $(seq 4000); do long="$long + 1 * X^1 - 1 * X^1"; done
solves "$long = 4 * X^0" "2"

echo "### mandatory: $FAILS failed check(s)"
exit "$FAILS"
