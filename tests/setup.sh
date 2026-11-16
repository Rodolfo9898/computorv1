#!/bin/bash
# Shared helpers, sourced by tests/mandatory.sh and tests/bonus.sh.
# Run from the root of the repository, after make. Nothing is written to disk.

# Bytes, not characters: the same results whatever the locale.
export LC_ALL=C

COMPUTOR=${COMPUTOR:-./computor_v1}
FAILS=0
# The solutions are compared as numbers, within 1e-5 of their size (an expected 0
# within 1e-12): the subject prints 6 decimals, the program 10 significant digits.
TOL=1e-5
# The arguments put before the equation: bonus.sh sets them to choose a mode.
MODE=()

if [ ! -x "$COMPUTOR" ]; then
	echo "$COMPUTOR not found: run make first (or make docker-mandatory, see the README)"
	exit 1
fi

# A stuck program is killed after 10 seconds (status 124), where timeout exists.
LIMIT=()
command -v timeout >/dev/null && LIMIT=(timeout 10)

check() { # $1 = condition result (0 = ok), $2 = description; a failure shows the output
	if [ "$1" -eq 0 ]; then
		echo "  PASS  $2"
	else
		echo "  FAIL  $2"
		printf '%s\n' "$OUTPUT" | head -12 | sed 's/^/        | /'
		FAILS=$((FAILS + 1))
	fi
}

no_colours() { sed 's/\x1b\[[0-9;]*m//g'; } # the ANSI colour codes of the usage and errors

# Runs computor_v1 with MODE and the equation as its argument: OUTPUT holds what it
# wrote (both outputs, without colours), STATUS its exit status.
run() { # $1 = equation
	OUTPUT=$("${LIMIT[@]}" "$COMPUTOR" "${MODE[@]}" "$1" 2>&1 </dev/null)
	STATUS=$?
	OUTPUT=$(printf '%s\n' "$OUTPUT" | no_colours)
}

# The same, with no argument: the equation is written on its standard input.
run_stdin() { # $1 = the standard input
	OUTPUT=$(printf '%s' "$1" | "${LIMIT[@]}" "$COMPUTOR" "${MODE[@]}" 2>&1)
	STATUS=$?
	OUTPUT=$(printf '%s\n' "$OUTPUT" | no_colours)
}

has_line() { printf '%s\n' "$OUTPUT" | grep -qxF -- "$1"; } # $1 = a whole line, word for word

crashed() { # killed, stopped by a signal, or a Go panic
	[ "$STATUS" -ge 124 ] || printf '%s\n' "$OUTPUT" | grep -qE '^(panic:|goroutine )'
}

label() { # how a check names its equation: the mode, then the equation, cut at 70 bytes
	local eq=$1
	[ "${#eq}" -gt 70 ] && eq="${eq:0:67}..."
	echo "${MODE[*]:+${MODE[*]} }\"$eq\""
}

# Each expected line is in the output, word for word, in any order.
gives() { # $1 = equation, the rest = the lines
	local eq=$1 line missing=""
	shift
	run "$eq"
	for line in "$@"; do
		has_line "$line" || missing="$missing [$line]"
	done
	! crashed && [ -z "$missing" ]
	check $? "$(label "$eq") gives $# expected line(s)${missing:+, missing:$missing}"
}

# One line of the output matches an extended regular expression.
matches() { # $1 = equation, $2 = regular expression, $3 = what it means
	run "$1"
	! crashed && printf '%s\n' "$OUTPUT" | grep -qE -- "$2"
	check $? "$(label "$1"): $3"
}

# No line of the output matches an extended regular expression.
lacks() { # $1 = equation, $2 = regular expression, $3 = what it means
	run "$1"
	! crashed && ! printf '%s\n' "$OUTPUT" | grep -qE -- "$2"
	check $? "$(label "$1"): $3"
}

# A fixed string is found on exactly n lines.
count_is() { # $1 = equation, $2 = the string, $3 = n
	local got
	run "$1"
	got=$(printf '%s\n' "$OUTPUT" | grep -cF -- "$2")
	! crashed && [ "$got" -eq "$3" ]
	check $? "$(label "$1"): \"$2\" printed $3 time(s), got $got"
}

# The real solutions printed ("x = v", "x1 = v", "x2 = v") are the expected ones, in
# any order, within TOL.
same_numbers() { # $1 = expected, $2 = printed: numbers separated by spaces
	awk -v want="$1" -v got="$2" -v tol="$TOL" '
	function sort(a, n,  i, j, t) {
		for (i = 2; i <= n; i++)
			for (j = i; j > 1 && a[j - 1] > a[j]; j--) { t = a[j]; a[j] = a[j - 1]; a[j - 1] = t }
	}
	BEGIN {
		n = split(want, w, " "); m = split(got, g, " ")
		if (n != m) exit 1
		for (i = 1; i <= n; i++) { w[i] += 0; g[i] += 0 }
		sort(w, n); sort(g, m)
		for (i = 1; i <= n; i++) {
			d = w[i] - g[i]; if (d < 0) d = -d
			s = w[i] < 0 ? -w[i] : w[i]; if (s == 0) s = 1e-7
			if (d > tol * s) exit 1
		}
	}'
}

printed_solutions() {
	printf '%s\n' "$OUTPUT" | sed -n 's/^x[12]\{0,1\} = \([-+0-9.e]*\)$/\1/p' | tr '\n' ' '
}

solves() { # $1 = equation, $2 = the expected solutions, separated by spaces
	local got
	run "$1"
	got=$(printed_solutions)
	! crashed && same_numbers "$2" "$got"
	check $? "$(label "$1") solves to $2 (printed: ${got:-nothing})"
}

solves_stdin() { # $1 = the standard input, $2 = the expected solutions
	local got
	run_stdin "$1"
	got=$(printed_solutions)
	! crashed && same_numbers "$2" "$got"
	check $? "standard input \"${1//$'\n'/\\n}\" solves to $2 (printed: ${got:-nothing})"
}

# Refused: no crash, a message, no reduced form (nothing was solved), and a status
# other than 0.
refuses() { # $1 = description, $2 = equation
	run "$2"
	! crashed && [ "$STATUS" -ne 0 ] && [ -n "$OUTPUT" ] && ! solved
	check $? "$1: $(label "$2") is refused (status $STATUS)"
}

refuses_stdin() { # $1 = description, $2 = the standard input
	run_stdin "$2"
	! crashed && [ "$STATUS" -ne 0 ] && [ -n "$OUTPUT" ] && ! solved
	check $? "$1 is refused (status $STATUS)"
}

solved() { printf '%s\n' "$OUTPUT" | grep -q '^Reduced form:'; } # a reduced form was printed
