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
| Bonus | `./computor_v1_bonus [mode] "<equation>"` | Free form entry, and the features of its mode |
| Bonus | `./computor_v1_bonus [mode]` | The same, with the equation read on the standard input |

Two programs, laid out like push_swap: `make` builds `./computor_v1` from
`src/computor/`, the mandatory part, and `make bonus` builds `./computor_v1_bonus` from
`src/bonus/`. Each folder has its own `main.go`, and both use the shared folders of
`src/`. Every file of `src/bonus/` starts with `//go:build bonus_mode`, the Go form of
`#if BONUS_MODE`: only `make bonus` (`-tags bonus_mode`) compiles them (see
*Technical choices*).

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
| `all`, `computor_v1` | Build `./computor_v1`, the mandatory program, from `src/computor` and the shared folders; again only when one of them or `go.mod` changed |
| `bonus`, `computor_v1_bonus` | Build `./computor_v1_bonus` with `-tags bonus_mode`, from `src/bonus` and the shared folders |
| `clean` | Nothing to remove: Go keeps its objects in its own build cache, outside the project |
| `fclean` | `clean`, and remove `./computor_v1` and `./computor_v1_bonus` |
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
Reduced form: 5 * X^0 - 6 * X^1 + 0 * X^2 - 5.6 * X^3 = 0
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

- The reduced form is sorted by exponent. A zero term written in the equation is
  kept up to the degree (the `0 * X^2` above, as in the subject); above the degree a
  term is zero because it cancelled out, and is left out: `1 * X^3 + 1 * X^1 = 1 * X^3`
  gives `1 * X^1 = 0`. A lone `0`, the right side of `... = 0`, is not a term.
- The degree is the one of the reduced form: `1 * X^3 + 1 * X^1 = 1 * X^3` has degree 1.
- Solutions are printed with up to 10 significant digits: `10`, `0.25`, `1e+10`.
- An equation that starts with a minus sign is an equation, not a flag:
  `./computor_v1 "-2.5 * X^2 + 10 * X^1 = 0"` gives 4 and 0.
- Coefficients whose sum does not fit in a double (two of 1e308 added) are refused,
  and a solution that does not fit in one (a coefficient of 1e-310, written in
  full, next to one of 1) is reported: the program never prints `+Inf`.
- An equation the parser does not understand, an empty standard input, or one that
  cannot be read, is refused with exit status 1. The reason goes to the standard
  error, followed by the usage:

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
make bonus
./computor_v1_bonus [mode] "<equation>"
./computor_v1_bonus [mode]          # the equation read on the standard input
```

As the subject asks (p.8): free form entry, graded on its own, then five additional
features. Free form entry is on in every mode, and each feature is a folder of
`src/bonus/` and a mode that turns it on alone; `all`, the default, turns all five on.

| Mode | Feature | Folder |
|---|---|---|
| `freeform` | Free form entry only (graded on its own) | `src/bonus/freeform/` |
| `mistakes` | 1. Entry mistakes: vocabulary and syntax | `src/bonus/mistakes/` |
| `fractions` | 2. Solutions as irreducible fractions | `src/bonus/fractions/` |
| `steps` | 3. Intermediate calculation steps | `src/bonus/steps/` |
| `factored` | 4. Factored form | `src/bonus/factored/` |
| `check` | 5. Solution check, P(x) = 0 | `src/bonus/check/` |
| `all` (when no mode is given) | the five features at once | |

The first three are the subject's own examples; the factored form and the check are
the two features chosen to complete the five, since both are about the solutions.

- **Free form entry.** A term is `[sign] [number] [*] [X[^n]]`: `X`, `4 * X`, `4X`,
  `X^3`, `-5`, `2.5 * x^2`, in any order. The strict form is one of its cases. The
  reduced form is written as by hand, as in the subject's example: `5 + 4 * X = 0`
  instead of `5 * X^0 + 4 * X^1 = 0`. Syntax mistakes are refused in every mode,
  with their place:
  - two signs (`1 ++ 2`) or two numbers (`1 2`) in a row
  - a `*` with nothing before it, or a number after `X` (`X2`)
  - a product of numbers (`2 * 3 * X`)
  - an exponent left out (`X^`), decimal (`X^2.5`) or negative (`X^-1`: not a
    polynomial)
- **1. Entry mistakes (vocabulary and syntax).** Before parsing, the vocabulary is
  checked: an unknown word is named (`unknown word "hello"`), and so is a character
  that cannot be in an equation (`@`). So is the structure: one `=`, no empty side,
  no empty equation. Each kind has its colour: vocabulary in cyan, structure in
  yellow, the syntax mistakes of the parser in red. Without this feature, the
  mistakes are still refused, in plain text.
- **2. Irreducible fractions.** A solution, or each part of a complex one, is printed
  as a fraction when it is one: `-5/4`, `-1/5 + 2/5i`. An irrational solution
  (`X^2 = 2`) stays a decimal, so that an approximation never passes for the exact
  value.
- **3. Intermediate steps.** The discriminant with its values, then only the formulas
  of the case at hand: the double root, the two real roots, or the real and
  imaginary parts.
- **4. Factored form.** The polynomial as a product: `(X - 1) * (X - 2)`, `(X + 1)^2`,
  `4 * (X + 5/4)`, `(X - (-1 + 2i)) * (X - (-1 - 2i))`.
- **5. Solution check.** Each solution, real or complex, is put back into the reduced
  form, `P(x)`, which must give 0. A solution is a rounded double, so `P(x)` is 0 when
  it is within 1e-9 of the size of its terms; otherwise the value is printed, with
  `(not zero)`.

```
$ ./computor_v1_bonus "5 + 4 * X + X^2= X^2"
Reduced form: 5 + 4 * X = 0
Polynomial degree: 1
Linear solve: x = -b / a = -(5) / 4
The solution is:
x = -5/4
Factored form: 4 * (X + 5/4) = 0
Check: P(-5/4) = 0

$ ./computor_v1_bonus "x^2 + 2x + 5 = 0"
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

$ ./computor_v1_bonus "X^ = 0"
Parse error: missing exponent after '^', near the end of the side

$ ./computor_v1_bonus mistakes "hello + X = 0"
unknown word "hello": an equation has numbers, X, + - * ^ = and spaces only

$ echo "X^2 - 1 = 0" | ./computor_v1_bonus freeform
Reduced form: -1 + X^2 = 0
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
make && make bonus
bash tests/mandatory.sh     # tests ./computor_v1
bash tests/bonus.sh         # tests ./computor_v1_bonus
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
  - the cases that rounding must not blur: `1 = 1.0000000000001` has no solution,
    `X² - 2X + 0.9999999999999` has two roots, and a decimal double root still has
    Δ = 0 when its constant is split between both sides
  - sums of coefficients and solutions too large for a double, refused or reported
  - a standard input that cannot be read (a directory)
  - that `./computor_v1` has no bonus: free form entry is refused, and no bonus text
    is in its binary
- `bonus.sh` checks:
  - free form entry: the subject's example word for word, the reduced form written
    as by hand, the syntax mistakes it must refuse in every mode, and the equation
    read on the standard input
  - each of the five features in its own mode: 1. entry mistakes (unknown words and
    characters, structure, and the colour of each kind), 2. fractions (exact ones,
    and none for an irrational solution), 3. steps (shown once, and only the ones
    that apply), 4. the factored form, 5. the check
  - the default mode, with the five features at once

`COMPUTOR=path/to/binary bash tests/mandatory.sh` (or `bonus.sh`) tests another build.

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
`make docker-shell` the programs are `./computor_v1` and `./computor_v1_bonus`, as on
your machine, and `go` is there too: `gofmt -l src`, `go vet ./...` and
`go vet -tags bonus_mode ./...` print nothing.

## Project structure

```
Makefile               rules only (all, bonus, clean, fclean, re, docker-*)
mk/                    everything the Makefile includes
  colors.mk              the output's colours: commands green, errors red, warnings
                         yellow, explanations blue; clean and fclean say what they remove
  config.mk              names of both programs, Go command and flags, the bonus_mode
                         tag, directories, Docker name
  sources.mk             source lists: one folder per program, then the shared ones;
                         each program depends
                         on its own, so changing the bonus never rebuilds ./computor_v1
Dockerfile             the test image (see Testing on Linux with Docker)
go.mod                 module computorv1, Go 1.25.4
src/
  computor/            ./computor_v1 (make): never imports src/bonus
    main.go            main(): --help, read, parse, print the reduced form and degree, solve
    parser.go          the strict parser: a regular expression per term
  bonus/               ./computor_v1_bonus (make bonus): every file is //go:build bonus_mode
    main.go            main(): the modes, their hooks into the solver
    freeform/          free form entry (graded on its own): the parser, the natural form
    mistakes/          1. entry mistakes: vocabulary, structure, and their colours
    fractions/         2. a float as an exact fraction, by continued fractions
    steps/             3. the intermediate steps
    factored/          4. the factored form
    check/             5. P(x) at each solution, in complex numbers
  input/               shared: the equation from the arguments or the standard input; flags
  poly/                shared: the reduced equation, coefficients and their noise, degree
  numeric/             shared: square root (Newton), absolute value, rounding noise, discriminant
  equation/            shared: the reduction both parsers use: terms moved left, noise cleared
  formatter/           shared: the reduced form, numbers, complex numbers
  solver/              shared:
    solve.go           the case and the solutions for degree 0, 1 or 2, printing nothing
    print.go           writes them as in the subject, with the hooks the bonus fills in
tests/
  setup.sh             shared helpers: run, compare, PASS / FAIL
  mandatory.sh         the mandatory checks
  bonus.sh             the bonus checks
```

## Technical choices

**The equation becomes a map from exponent to coefficient** (`poly.PolyCoeffs`). Each
side is parsed into a list of terms, and `equation.Parse` (`src/equation`), shared by both parsers,
adds them up by exponent, those of the right side subtracted. The reduced form is this
map sorted by exponent, and the degree is the highest exponent whose coefficient is
not zero, so terms that cancel lower the degree.

**Zero is decided from the rounding noise.** Each decimal is rounded once when it is
read, and each addition rounds once more: a sum of n terms whose sizes add up to S is
within `2·n·u·S` of the exact one (u = 2⁻⁵³). A coefficient no larger than that is a
cancellation and becomes exactly 0: `0.1 + 0.2 - 0.3` leaves 5.5e-17. That bound, the
noise, is kept with each coefficient (`poly.Equation`). The discriminant carries the
noise of a, b and c through `b² - 4ac`, plus the rounding of its products, and is 0
when it is no larger: `0.01X² + 0.2X + 1.3 = 0.3` has a double root at -10.
- The first version cleared any sum within 1e-12 of its terms, a guess: `1 = 1.0000000000001`
  had every real as solution, and `X² - 2X + 0.9999999999999` (Δ = 4e-13) a double root.
- A tolerance of a few ulps of the last addition alone is not enough either: it lost 7%
  of 19701 decimal double roots once their constant was split between both sides,
  because the noise of the earlier additions was forgotten.
- A small coefficient on its own, like `1e-10 * X^2`, is far above its noise and is
  kept: a fixed threshold (an earlier version cut everything under 1e-9) would turn
  that equation into one of degree 0.

**The square root is Newton's method, until it stops moving:** `g = (g + v / g) / 2`,
started at `v` (or 1 when `v < 1`), which is above the root. Every step then goes down,
and the first one that does not is the end: the result is the closest double. The
first version did 20 steps, too few above about 1e11. It gave 1279996.5 for √1e12.

**The real roots avoid subtracting close numbers.** When b² is much larger than 4ac,
`-b + √Δ` loses its digits. For `X² + 100000000X + 1` it gives -7.45e-09 for the small
root, -1e-08. The root of the same sign as `-b` is computed as `q / a`, with
`q = -(b ± √Δ) / 2`. The other one is `c / q`, since `x1 * x2 = c / a`.

**An overflow is reported.** A coefficient of 1e200 makes b² infinite: the program says
the coefficients are too large rather than printing `+Inf` or a wrong case. Coefficients
whose sum passes the largest double are refused when the equation is reduced, and a
solution that does not fit (an X^1 coefficient of 1e-310 next to a constant of 1) is
reported instead of printed.

**Numbers are printed with `%.10g`**, which already drops the zeros at the end of the
decimals and only those. The first version also trimmed every `0` at the end, so 10
became `1`.

**Strict parsing cuts each side before every sign**, except the sign of an exponent.
Each piece must then match `a * X^p` or be a constant (`src/computor/parser.go`).
A sign that follows another one is a piece of its own that matches nothing, which is
how `1 ++ 2` is refused.

**Free form parsing is a scanner** (`src/bonus/freeform/parser.go`): it reads a sign,
a number, a `*`, an `X`, a `^` and its digits, and says what it expected where it
stops. The first version cut on the signs and guessed, and accepted `1 2` as 12 and
`X^-1`.

**A fraction is shown only when it is exact.** The continued fraction of the value is
walked until a convergent equals it up to rounding (1e-13), with a denominator of
10⁶ at most. An irrational value has no such convergent, and stays a decimal.

**The bonus reuses the mandatory part, which does not know about it.**
`solver.Solve` only computes; `solver.Print` writes the result and calls hooks
(`solver.Hooks`) at three places: before the line that names the case, after it,
and after the solutions, plus the function that writes a number. The mandatory part
passes no hook. The bonus fills them in from its mode: the steps, the factored form,
the check, and fractions for the numbers. No file outside `src/bonus` imports
one of its packages.

**Two programs, as in push_swap, and the BONUS_MODE trick with a Go build tag.**
push_swap has one folder per program, each with its `main.c` (`push_swap/`, and
`bonus/` for the checker), next to the folders both share; its bonus files are
wrapped in `#if BONUS_MODE`, and `make bonus` compiles them with `-D BONUS_MODE=1`.
Go has one package per folder, and a folder whose files say `package main` is a
program: `src/computor/main.go` and `src/bonus/main.go` are two of them, both named
`main.go` since they are in different folders. Go has no preprocessor, but a file
can start with a build constraint: every file of `src/bonus/` starts with
`//go:build bonus_mode`, so `go build ./src/bonus` refuses to build without the tag
("build constraints exclude all Go files"), and `make bonus` builds it with
`-tags bonus_mode`. The mandatory program holds no bonus code: `go list -deps
./src/computor` lists only `src/computor` and the shared folders, and
`tests/mandatory.sh` checks that no bonus text is in its binary.

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

**Use of AI.** An AI assistant (Claude) reviewed the finished project against the subject, pointed out failures and explained why they did not match the subject, and proposed fixes.
