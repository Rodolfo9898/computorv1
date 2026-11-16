# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    Makefile                                           :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: rperez-t <rperez-t@student.42belgium.be>   +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2026/11/16 17:00:00 by rperez-t          #+#    #+#              #
#    Updated: 2026/11/16 17:00:00 by rperez-t         ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

include mk/colors.mk
include mk/config.mk
include mk/sources.mk

all: $(NAME)

# The bonus is in the same program, behind --bonus: nothing more to build.
bonus: $(NAME)
	@$(call explain,The bonus is part of $(NAME): run ./$(NAME) --bonus)

# go build compiles only the packages that changed (in its own cache, outside
# the project), so the program is rebuilt only when a source or go.mod changed.
$(NAME): $(GO_FILES) go.mod
	@if ! command -v $(GO) >/dev/null; then \
		printf '$(RED)%s$(RESET)\n' "$(GO) not found: install Go 1.25.4 or later, or use make docker-shell (see the README)"; \
		exit 1; \
	fi
	@$(call run,$(GO) build $(GO_FLAGS) -o $(NAME) $(CMD_DIR))

# Go keeps no object file in the project: they are in its build cache, shared
# by every Go project of this machine, which clean leaves alone.
clean:
	@$(call warn,nothing to clean: Go keeps its objects in its own build cache)

fclean: clean
	@$(call remove,$(NAME),Removing the program $(NAME))

re: fclean all

# Docker: for a machine without Go 1.25. Each rule rebuilds the image first, so
# it always tests the current sources (docker's cache makes this instant when
# nothing changed).
docker-build:
	docker build -t $(DOCKER_IMAGE) .

# An interactive shell exits with the status of its last command, which says
# nothing about the rule: only 125, docker itself failing, is an error.
docker-shell: docker-build
	@docker run --rm -it $(DOCKER_IMAGE) bash || [ $$? -ne 125 ]

docker-mandatory: docker-build
	docker run --rm $(DOCKER_IMAGE) bash tests/mandatory.sh

docker-bonus: docker-build
	docker run --rm $(DOCKER_IMAGE) bash tests/bonus.sh

.PHONY: all bonus clean fclean re docker-build docker-shell docker-mandatory docker-bonus
