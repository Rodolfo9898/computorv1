# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    sources.mk                                         :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: rperez-t <rperez-t@student.42belgium.be>   +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2026/11/16 17:00:00 by rperez-t          #+#    #+#              #
#    Updated: 2026/11/16 17:00:00 by rperez-t         ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

# One folder per program, each with its main() in main.go (like push_swap's
# push_swap/ and bonus/): computor/, and bonus/ for the bonus program. Then the
# folders both programs share.
MANDATORY_DIR = computor
BONUS_DIR = bonus

# Shared by both programs: reading the arguments, the reduced equation, the
# arithmetic, the reduction of the parsed terms, the printing and the solver.
SHARED_SRCS = input/input.go poly/poly.go numeric/numeric.go \
			equation/equation.go formatter/formatter.go \
			solver/solve.go solver/print.go

# ./computor_v1: its main() and the strict parser.
MANDATORY_SRCS = main.go parser.go

# ./computor_v1_bonus, which only make bonus builds: every one of its files starts
# with //go:build bonus_mode, so it is compiled with -tags bonus_mode only
# (BONUS_TAG), as a C bonus with -D BONUS_MODE=1. Its main(), free form entry
# (graded on its own), then the five features of the subject, one folder each:
# 1. entry mistakes, 2. fractions, 3. steps, 4. factored form, 5. check.
BONUS_SRCS = main.go freeform/parser.go freeform/natural.go \
			mistakes/mistakes.go mistakes/colors.go fractions/format.go \
			steps/output.go factored/factored.go check/check.go

SHARED_FILES = $(addprefix $(SRC_DIR)/,$(SHARED_SRCS))

MANDATORY_FILES = $(addprefix $(SRC_DIR)/$(MANDATORY_DIR)/,$(MANDATORY_SRCS)) \
			$(SHARED_FILES)
BONUS_FILES = $(addprefix $(SRC_DIR)/$(BONUS_DIR)/,$(BONUS_SRCS)) \
			$(SHARED_FILES)
