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

bonus: $(BONUS_NAME)

# One program per folder, as in push_swap: src/computor for ./computor_v1 and
# src/bonus for ./computor_v1_bonus, both using the shared folders of src/. The
# files of src/bonus all start with //go:build bonus_mode: only -tags
# $(BONUS_TAG), the BONUS_MODE of a C Makefile, compiles them. go build compiles
# only the packages that changed (in its own cache, outside the project), so a
# program is rebuilt only when one of its sources or go.mod changed.
$(NAME): $(MANDATORY_FILES) go.mod
	@$(call need_go)
	@$(call run,$(GO) build $(GO_FLAGS) -o $(NAME) $(SRC_DIR)/$(MANDATORY_DIR))

$(BONUS_NAME): $(BONUS_FILES) go.mod
	@$(call need_go)
	@$(call run,$(GO) build $(GO_FLAGS) -tags $(BONUS_TAG) -o $(BONUS_NAME) $(SRC_DIR)/$(BONUS_DIR))

# Go keeps no object file in the project: they are in its build cache, shared
# by every Go project of this machine, which clean leaves alone.
clean:
	@$(call warn,nothing to clean: Go keeps its objects in its own build cache)

fclean: clean
	@$(call remove,$(NAME) $(BONUS_NAME),Removing the programs $(NAME) and $(BONUS_NAME))

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
