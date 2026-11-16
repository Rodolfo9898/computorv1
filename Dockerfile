# Linux test image for computor_v1: builds ./computor_v1 with the Go toolchain the
# project needs (go.mod asks for Go 1.25.4 or later), then runs the checks:
# tests/mandatory.sh, then tests/bonus.sh, both using tests/setup.sh.
#   - every example of the subject must give its reduced form, its degree and its
#     solutions, and so must equations with zero, negative, decimal or large
#     coefficients, and an equation read on the standard input;
#   - malformed equations must be refused with an error, without a crash;
#   - the bonus modes (free form, validation, fractions, steps, natural form,
#     factored form, check, interactive) must do what
#     ./computor_v1 --help and the README say.
# The exit status is the total number of failed checks.
#
#   make docker-mandatory                                   # mandatory part only
#   make docker-bonus                                       # bonus part only
#   make docker-shell                                       # bash inside the image
#   make docker-build && docker run --rm computorv1-test    # mandatory, then bonus
FROM golang:1.25-trixie

# The Go image has no make; bash runs the test scripts.
RUN apt-get update \
	&& DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
		make bash \
	&& rm -rf /var/lib/apt/lists/*

WORKDIR /work
COPY . .
RUN make

CMD ["bash", "-c", "bash tests/mandatory.sh; m=$?; bash tests/bonus.sh; exit $((m + $?))"]
