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

# One Go package per folder of src/. The program is built from cmd/ (main.go),
# which picks the mandatory part (cmd/mandatory) or the bonus (cmd/bonus): one
# binary for both, the bonus only runs with --bonus.
CMD_SRCS = main.go mandatory/cmd_mandatory.go bonus/cmd_bonus.go

# The mandatory part: read, parse strictly, reduce, solve.
BASE_SRCS = functions.go structs.go types.go
CALCULATIONS_SRCS = numeric.go
CORE_SRCS = parser/parser.go formatter/formatter.go
EXECUTION_SRCS = solve.go
IO_SRCS = input.go

# The bonus: one package per feature.
BONUS_SRCS = check/check.go colors/colors.go factored/factored.go \
			fractions/format.go freeform/parser.go steps/output.go \
			validator/validator.go

GO_FILES = $(addprefix $(SRC_DIR)/cmd/,$(CMD_SRCS)) \
			$(addprefix $(SRC_DIR)/base/,$(BASE_SRCS)) \
			$(addprefix $(SRC_DIR)/calculations/,$(CALCULATIONS_SRCS)) \
			$(addprefix $(SRC_DIR)/core/,$(CORE_SRCS)) \
			$(addprefix $(SRC_DIR)/execution/,$(EXECUTION_SRCS)) \
			$(addprefix $(SRC_DIR)/io/,$(IO_SRCS)) \
			$(addprefix $(SRC_DIR)/bonus/,$(BONUS_SRCS))
