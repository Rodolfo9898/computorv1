# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    colors.mk                                          :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: rperez-t <rperez-t@student.42belgium.be>   +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2026/11/16 17:00:00 by rperez-t          #+#    #+#              #
#    Updated: 2026/11/16 17:00:00 by rperez-t         ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

# The recipes run in bash: 'set -o pipefail' keeps the exit status of a command
# whose output goes through COLORIZE (make 3.81, the one of macOS, has no
# .SHELLFLAGS, so run sets it itself).
SHELL := /bin/bash

RED		= \033[0;31m
GREEN	= \033[0;32m
YELLOW	= \033[0;33m
BLUE	= \033[0;34m
RESET	= \033[0m

# What the Go toolchain prints: errors (file.go:line:column: ..., and what go
# says it cannot do) in red, the package each group of errors is in (# name)
# in yellow, go commands in green, the rest as it is.
COLORIZE = awk '\
	/\.go:[0-9]+:[0-9]+: |^go: |cannot find|\*\*\* .*Error [0-9]/ \
		{ printf "$(RED)%s$(RESET)\n", $$0; next } \
	/^\# / { printf "$(YELLOW)%s$(RESET)\n", $$0; next } \
	/^go / { printf "$(GREEN)%s$(RESET)\n", $$0; next } \
	{ print }'

# $(call explain,text): what is going to happen, in blue (no comma in the text).
explain = printf '$(BLUE)%s$(RESET)\n' '$(1)'

# $(call warn,text): a warning, in yellow (no comma in the text).
warn = printf '$(YELLOW)warning: %s$(RESET)\n' '$(1)'

# $(call run,command): shows the command in green, then runs it. What it prints
# goes through COLORIZE, and the rule fails when the command fails.
run = set -o pipefail; printf '$(GREEN)%s$(RESET)\n' '$(1)'; \
	$(1) 2>&1 | $(COLORIZE)

# $(call need_go): stops the rule, with a message, when Go is not installed.
need_go = if ! command -v $(GO) >/dev/null; then \
		printf '$(RED)%s$(RESET)\n' "$(GO) not found: install Go 1.25.4 or later, or use make docker-shell (see the README)"; \
		exit 1; \
	fi

# $(call exists,paths): true when at least one of the paths (or patterns) exists.
exists = [ -n "$$(ls -d $(1) 2>/dev/null)" ]

# $(call remove,paths,why[,name]): when one of the paths exists, says why in
# blue, then shows the rm in green and runs it. Otherwise, a warning: nothing
# to remove (named after name when given: handy for a long list of objects).
remove = if $(call exists,$(1)); then \
		$(call explain,$(2)); $(call run,rm -rf $(1)); \
	else \
		$(call warn,$(if $(3),no $(3) to remove,nothing to remove: $(1) does not exist)); \
	fi
