BUILD_DIR := build
BIN := $(BUILD_DIR)/computor_v1
CMD_DIR := src/cmd
SRCS := $(shell find src -name '*.go')
GO ?= go
RESET := \033[0m
RED := \033[31m
GREEN := \033[32m
YELLOW := \033[33m

all: $(BIN)
	@echo "$(GREEN)Target up to date$(RESET)"

$(BIN): $(SRCS) go.mod
	@mkdir -p $(BUILD_DIR)
	@$(GO) build -o $(BIN) ./$(CMD_DIR) && echo "$(GREEN)Built $(BIN)$(RESET)" || { echo "$(RED)Build failed$(RESET)"; exit 1; }

clean:
	@if [ -d "$(BUILD_DIR)" ]; then \
		rm -rf "$(BUILD_DIR)" && echo "$(GREEN)Removed $(BUILD_DIR)$(RESET)"; \
	else \
		echo "$(YELLOW)$(BUILD_DIR) does not exist$(RESET)"; \
	fi

fclean: clean

re: fclean all

.PHONY: all clean fclean re
