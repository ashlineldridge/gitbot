NAME      := gitbot
BUILD_DIR := build
BIN_DIR   := $(BUILD_DIR)/bin
BINARIES  := $(BIN_DIR)/linux/amd64/$(NAME) $(BIN_DIR)/darwin/amd64/$(NAME)
SRC       := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

NO_COLOR    := \033[0m
OK_COLOR    := \033[32;01m
ERROR_COLOR := \033[31;01m
WARN_COLOR  := \033[33;01m

export GO111MODULE=on

all: test $(BINARIES)

.PHONY: test
test:
	@echo "\n$(OK_COLOR)====> Running tests$(NO_COLOR)"
	go test ./...

.PHONY: clean
clean:
	@echo "\n$(OK_COLOR)====> Cleaning$(NO_COLOR)"
	go clean ./... && rm -rf ./$(BUILD_DIR)

$(BINARIES): splitted=$(subst /, ,$@)
$(BINARIES): os=$(word 3, $(splitted))
$(BINARIES): arch=$(word 4, $(splitted))
$(BINARIES): cmd=$(basename $(word 5, $(splitted)))
$(BINARIES): $(SRC)
	@echo "\n$(OK_COLOR)====> Building $@$(NO_COLOR)"
	GOOS=$(os) GOARCH=$(arch) CGO_ENABLED=0 go build -mod=vendor -a -o $@ cmd/$(cmd)/$(cmd).go
