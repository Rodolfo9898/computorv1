*This project has been created as part of the 42 curriculum by rperez-t.*

# computor_v1

## Description

A polynomial equation solver. `computor_v1` reads an equation of degree 2 or less,
moves every term to the left side, and prints the reduced form, the degree and the
solutions: one, two (real or complex, depending on the sign of the discriminant),
none, or every real number. An equation of degree 3 or more is reduced and its degree
given, but it is not solved.

The subject forbids math library functions other than `+ - * /`. The program
therefore computes its own square root (Newton's method) and absolute value, and
never imports Go's `math` package.

| Part | Command | What it does |
|---|---|---|
| Mandatory | `./computor_v1 "<equation>"` | Strict entry: every term is `a * X^p` |
| Mandatory | `./computor_v1` | The same, with the equation read on the standard input |
| Bonus | `./computor_v1 --bonus [mode] "<equation>"` | Free form entry, and the features of its mode |
| Bonus | `./computor_v1 --bonus [mode]` | Interactive: one equation per line on the standard input |

## Instructions

### Requirements

Go 1.25.4 or later (`go.mod`) and `make`. The project uses Go's standard library only:
nothing is downloaded at build time. Without Go, `make` says so and stops.

On a machine without Go, `make docker-mandatory` builds and tests it in Docker, and
`make docker-shell` gives a shell where it is built (see *Testing on Linux with Docker*).

### Building

```sh
make
```

| Rule | Effect |
|---|---|
| `all`, `computor_v1` | Build `./computor_v1`, again only when a source of `mk/sources.mk` or `go.mod` changed |
| `bonus` | The same: the bonus is in the same program, behind `--bonus` |
| `clean` | Nothing to remove: Go keeps its objects in its own build cache, outside the project |
| `fclean` | `clean`, and remove `./computor_v1` |
| `re` | `fclean`, then `all` |
| `docker-mandatory` | Run the mandatory checks in Docker (see *Testing on Linux with Docker*) |
| `docker-bonus` | Run the bonus checks in Docker |
| `docker-shell` | Open a bash shell in the Docker image, with the program built |
| `docker-build` | Build the image (the three rules above do it for you) |

### Running

```sh
./computor_v1 "5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0"
echo "5 * X^0 + 4 * X^1 = 4 * X^0" | ./computor_v1
./computor_v1 --help
```

The entry, as in the subject: terms `a * X^p`, joined by `+` or `-`, on both sides
of a single `=`. A constant can stand alone (the `0` of `... = 0`). Coefficients can
be negative or decimal, exponents are whole numbers, and `X` can be written `x`.

```
$ ./computor_v1 "5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0"
Reduced form: 4 * X^0 + 4 * X^1 - 9.3 * X^2 = 0
Polynomial degree: 2
Discriminant is strictly positive, the two solutions are:
x1 = 0.9052389908
x2 = -0.4751314639

$ ./computor_v1 "1 * X^0 + 2 * X^1 + 5 * X^2 = 0"
Reduced form: 1 * X^0 + 2 * X^1 + 5 * X^2 = 0
Polynomial degree: 2
Discriminant is strictly negative, the two complex solutions are:
x1 = -0.2 + 0.4i
x2 = -0.2 - 0.4i

$ ./computor_v1 "5 * X^0 + 4 * X^1 = 4 * X^0"
Reduced form: 1 * X^0 + 4 * X^1 = 0
Polynomial degree: 1
The solution is:
x = -0.25

$ ./computor_v1 "8 * X^0 - 6 * X^1 + 0 * X^2 - 5.6 * X^3 = 3 * X^0"
Reduced form: 5 * X^0 - 6 * X^1 - 5.6 * X^3 = 0
Polynomial degree: 3
The polynomial degree is strictly greater than 2, I can't solve.

$ ./computor_v1 "6 * X^0 = 6 * X^0"
Reduced form: 0 * X^0 = 0
Polynomial degree: 0
Any real number is a solution.

$ ./computor_v1 "10 * X^0 = 15 * X^0"
Reduced form: -5 * X^0 = 0
Polynomial degree: 0
No solution.
```

- Terms whose coefficient is zero are left out of the reduced form, which is sorted
  by exponent.
- The degree is the one of the reduced form: `1 * X^3 + 1 * X^1 = 1 * X^3` has degree 1.
- Solutions are printed with up to 10 significant digits: `10`, `0.25`, `1e+10`.
- An equation that starts with a minus sign is an equation, not a flag:
  `./computor_v1 "-2.5 * X^2 + 10 * X^1 = 0"` gives 4 and 0.
- An equation the parser does not understand is refused with exit status 1. The
  reason goes to the standard error, followed by the usage:

```
$ ./computor_v1 "1 * X^0 ++ 2 * X^0 = 0"
Error: invalid term "+" (expected a * X^p)
```

More equations worth trying:

```
6 * X^0 + 11 * X^1 + 5 * X^2 = 1 * X^0 + 1 * X^1    discriminant zero: x = -1
5 * X^0 + 3 * X^1 + 3 * X^2 = 1 * X^0 + 0 * X^1     discriminant negative
1 * X^2 - 30 * X^1 + 200 * X^0 = 0                   x = 10 and x = 20
1 * X^2 + 1000000 * X^1 = 0                          x = -1000000 and x = 0
0.0000000001 * X^2 + 1 * X^0 = 0                     degree 2: a small coefficient still counts
1 * X^2 + 2 * X^1 = 1 * X^2 + 4 * X^0                the X^2 cancel: degree 1, x = 2
```

### Bonus

```sh
./computor_v1 --bonus [mode] "<equation>"
./computor_v1 --bonus [mode]          # interactive
```

Free form entry is on in every mode. Each other mode turns on its own feature only;
`all`, the default, turns them all on.

| Mode | Feature |
|---|---|
| `all` (when no mode is given) | every feature below |
| `steps` | Intermediate steps |
| `fractions` | Irreducible fractions |
| `validation` | Vocabulary check, with coloured errors |
| `natural` | Natural reduced form |
| `factored` | Factored form |
| `check` | Solution check |
| `freeform` | none: free form entry only |
| `novalidation` | every feature but validation |

- **Free form entry.** A term is `[sign] [number] [*] [X[^n]]`: `X`, `4 * X`, `4X`,
  `X^3`, `-5`, `2.5 * x^2`, in any order. The strict form is one of its cases.
  Syntax mistakes are refused in every mode, with their place:
  - two signs (`1 ++ 2`) or two numbers (`1 2`) in a row
  - a `*` with nothing before it, or a number after `X` (`X2`)
  - a product of numbers (`2 * 3 * X`)
  - an exponent left out (`X^`), decimal (`X^2.5`) or negative (`X^-1`: not a
    polynomial)
- **Validation.** Before parsing, it refuses an empty equation, a character that
  cannot be in an equation (`@`), a number of `=` other than one, or an empty side.
  Each kind of error has its own colour.
- **Irreducible fractions.** A solution, or each part of a complex one, is printed as
  a fraction when it is one: `-5/4`, `-1/5 + 2/5i`. An irrational solution (`X^2 = 2`)
  stays a decimal, so that an approximation never passes for the exact value.
- **Intermediate steps.** The discriminant with its values, then only the formulas of
  the case at hand: the double root, the two real roots, or the real and imaginary
  parts.
- **Natural reduced form.** The reduced form as written by hand, as in the subject's
  example: `5 + 4 * X = 0` instead of `5 * X^0 + 4 * X^1 = 0`.
- **Factored form.** The polynomial as a product: `(X - 1) * (X - 2)`, `(X + 1)^2`,
  `4 * (X + 5/4)`, `(X - (-1 + 2i)) * (X - (-1 - 2i))`.
- **Solution check.** Each solution, real or complex, is put back into the reduced
  form, `P(x)`, which must give 0. A solution is a rounded double, so `P(x)` is 0 when
  it is within 1e-9 of the size of its terms; otherwise the value is printed, with
  `(not zero)`.
- **Interactive mode.** Without an equation, `--bonus` reads one per line until the
  end of the input or `quit`. A terminal gets a `>` prompt. Otherwise each equation
  is written before its result. The exit status is 1 when a line was refused; the
  lines after it are still solved.
- **Colours.** The usage in grey, and the errors in red, cyan or yellow, depending on
  their kind.

```
$ ./computor_v1 --bonus "5 + 4 * X + X^2= X^2"
Reduced form: 5 + 4 * X = 0
Polynomial degree: 1
Linear solve: x = -b / a = -(5) / 4
The solution is:
x = -5/4
Factored form: 4 * (X + 5/4) = 0
Check: P(-5/4) = 0

$ ./computor_v1 --bonus "x^2 + 2x + 5 = 0"
Reduced form: 5 + 2 * X + X^2 = 0
Polynomial degree: 2
Discriminant Δ = b^2 - 4ac = (2)^2 - 4 * (1) * (5) = -16
Discriminant is strictly negative, the two complex solutions are:
sqrt(-Δ) = 4
Real part = -b / (2a) = -(2) / (2 * 1)
Imag part = sqrt(-Δ) / (2a) = 4 / (2 * 1)
x1 = -1 + 2i
x2 = -1 - 2i
Factored form: (X - (-1 + 2i)) * (X - (-1 - 2i)) = 0
Check: P(-1 + 2i) = 0
Check: P(-1 - 2i) = 0

$ ./computor_v1 --bonus "X^ = 0"
Parse error: missing exponent after '^', near the end of the side

$ printf 'X = 2\nX^2 - 1 = 0\n' | ./computor_v1 --bonus freeform
Equation: X = 2
Reduced form: -2 * X^0 + 1 * X^1 = 0
Polynomial degree: 1
The solution is:
x = 2

Equation: X^2 - 1 = 0
Reduced form: -1 * X^0 + 1 * X^2 = 0
Polynomial degree: 2
Discriminant is strictly positive, the two solutions are:
x1 = -1
x2 = 1
```

More equations for the bonus:

```
X^2 + X = 0                   x = -1 and x = 0
4 = 0                         no solution
7 = 7                         any real number
X = 2 * X                     x = 0
-3 * X^3 + 2 * X^2 + 1 = 0    degree 3: not solved
X^4 + X^2 = X^4               the X^4 cancel: degree 2, x = 0
X^2 - 3 * X + 2 = 0           factored form (X - 1) * (X - 2)
```

### Testing

```sh
make
bash tests/mandatory.sh
bash tests/bonus.sh
```

The exit status of each script is its number of failed checks. Nothing is written to
disk: each output is kept in a variable. A failed check shows the start of the output.
The expected results are the mathematically correct ones, worked out by hand. They do
not come from the program's own output. Solutions are compared as numbers, within
1e-5, since the subject prints 6 decimals and the program 10 significant digits.

- `mandatory.sh` runs every example of the subject. It also checks:
  - zero, negative, decimal, tiny (1e-10) and large (up to 1e150) coefficients
  - coefficients so large that the discriminant overflows, which must be reported
  - equations whose higher terms cancel
  - two equations on the standard input, and one of 8000 terms
  - malformed equations (no `=`, two `=`, words, two signs or two numbers in a row,
    an empty side, a negative exponent, an empty argument or input, an exponent too
    big for an `int`), which must be refused with a message and a status other than
    0, without a crash
- `bonus.sh` checks:
  - free form equations, and the syntax mistakes it must refuse in every mode
  - each mode: validation, fractions (exact ones, and none for an irrational
    solution), steps (shown once, and only the ones that apply), the natural and
    factored forms, the check
  - the default mode, with everything at once
  - the interactive mode (several lines, an empty line, `quit`, a refused line, an
    empty input)

`COMPUTOR=path/to/binary bash tests/mandatory.sh` tests another build.

### Testing on Linux with Docker

The `Dockerfile` builds the program in the official Go image (Debian), and runs the
same checks there, so the result does not depend on your machine or on having Go.

```sh
make docker-mandatory                                   # mandatory part only
make docker-bonus                                       # bonus part only
make docker-shell                                       # bash inside the image
make docker-build && docker run --rm computorv1-test    # mandatory, then bonus
```

Each rule rebuilds the image first, so it always tests the current sources.
`make docker-mandatory` and `make docker-bonus` fail if any check fails. In
`make docker-shell` the program is `./computor_v1`, as on your machine, and `go`
is there too: `gofmt -l src` and `go vet ./...` print nothing.

## Project structure

```
Makefile               rules only (all, bonus, clean, fclean, re, docker-*)
mk/                    everything the Makefile includes
  colors.mk              the output's colours: commands green, errors red, warnings
                         yellow, explanations blue; clean and fclean say what they remove
  config.mk              name, Go command and flags, directories, Docker name
  sources.mk             source list per folder
Dockerfile             the test image (see Testing on Linux with Docker)
go.mod                 module computorv1, Go 1.25.4
src/
  cmd/
    main.go            chooses the mode: mandatory, --bonus or --help
    mandatory/         reads, parses (strict), prints the reduced form and degree, solves
    bonus/             the bonus modes, the interactive loop, the free form parser
  io/input.go          the equation: the arguments joined, or a line of the standard input
  core/
    parser/            the strict parser: a * X^p terms, a regular expression per term
    formatter/         reduced form (strict or natural), numbers, complex numbers
  execution/solve.go   degree, then the solution for degree 0, 1 or 2
  calculations/        square root (Newton), absolute value, sum that clears rounding noise
  base/                shared types: coefficients by exponent, options, errors
  bonus/
    freeform/          the free form parser
    validator/         the checks run before parsing
    fractions/         a float as an exact fraction, by continued fractions
    steps/             the intermediate steps
    factored/          the factored form
    check/             P(x) at each solution, in complex numbers
    colors/            the ANSI colours
tests/
  setup.sh             shared helpers: run, compare, PASS / FAIL
  mandatory.sh         the mandatory checks
  bonus.sh             the bonus checks
```

## Technical choices

**The equation becomes a map from exponent to coefficient** (`base.PolyCoeffs`). Each
side is parsed into one, the right one is subtracted from the left, and a term given
twice adds up. The reduced form is this map sorted by exponent, and the degree is the
highest exponent whose coefficient is not zero, so terms that cancel lower the degree.

**Zero is decided where it is made.** Every term is added with `calculations.Sum`.
When a sum is within 1e-12 of the size of the numbers added, it is rounding noise and
becomes exactly 0: `0.1 + 0.2 - 0.3` leaves 5.5e-17. After that, zero means `== 0`. A
small coefficient on its own, like `1e-10 * X^2`, is kept: a fixed threshold (the first
version cut everything under 1e-9) would turn that equation into one of degree 0. The
discriminant `b² - 4ac` goes through the same sum, so `5X² + 10X + 5` has exactly
Δ = 0.

**The square root is Newton's method, until it stops moving:** `g = (g + v / g) / 2`,
started at `v` (or 1 when `v < 1`), which is above the root. Every step then goes down,
and the first one that does not is the end: the result is the closest double. The
first version did 20 steps, too few above about 1e11. It gave 1279996.5 for √1e12.

**The real roots avoid subtracting close numbers.** When b² is much larger than 4ac,
`-b + √Δ` loses its digits. For `X² + 100000000X + 1` it gives -7.45e-09 for the small
root, -1e-08. The root of the same sign as `-b` is computed as `q / a`, with
`q = -(b ± √Δ) / 2`. The other one is `c / q`, since `x1 * x2 = c / a`.

**An overflow is reported.** A coefficient of 1e200 makes b² infinite: the program says
the coefficients are too large rather than printing `+Inf` or a wrong case.

**Numbers are printed with `%.10g`**, which already drops the zeros at the end of the
decimals and only those. The first version also trimmed every `0` at the end, so 10
became `1`.

**Strict parsing cuts each side before every sign**, except the sign of an exponent.
Each piece must then match `a * X^p` or be a constant (`src/core/parser/parser.go`).
A sign that follows another one is a piece of its own that matches nothing, which is
how `1 ++ 2` is refused.

**Free form parsing is a scanner** (`src/bonus/freeform/parser.go`): it reads a sign,
a number, a `*`, an `X`, a `^` and its digits, and says what it expected where it
stops. The first version cut on the signs and guessed, and accepted `1 2` as 12 and
`X^-1`.

**A fraction is shown only when it is exact.** The continued fraction of the value is
walked until a convergent equals it up to rounding (1e-13), with a denominator of
10⁶ at most. An irrational value has no such convergent, and stays a decimal.

**The bonus reuses the mandatory part.** The modes are options of the same solver
(`base.SolveOptions`), and each feature is a package of `src/bonus/` that the solver
calls when its option is on.

## Resources

**The subject's topics**
- Quadratic equation and discriminant - https://en.wikipedia.org/wiki/Quadratic_equation
- The numerically stable quadratic formula - https://en.wikipedia.org/wiki/Loss_of_significance
- Complex numbers - https://en.wikipedia.org/wiki/Complex_number
- Newton's method for square roots - https://en.wikipedia.org/wiki/Methods_of_computing_square_roots
- Continued fractions - https://en.wikipedia.org/wiki/Continued_fraction
- Floating point - https://docs.oracle.com/cd/E19957-01/806-3568/ncg_goldberg.html

**Tools**
- Go standard library: `strconv`, `regexp`, `fmt`, `bufio` - https://pkg.go.dev/std
- Docker - https://docs.docker.com/reference/dockerfile/

**Use of AI.** An AI assistant (Claude) reviewed the project against the subject. It
found the bugs described in *Technical choices* with its tests, and fixed them. It
also wrote the Dockerfile, the Makefile's `mk/` split, the shell tests and this README,
and added the natural form, factored form, solution check and interactive mode.
