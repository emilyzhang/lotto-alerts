BIN_DIR  ?= ./bin
PKG_NAME ?= lotto-alerts

GOTOOLS := \
golang.org/x/tools/cmd/cover \

default: build

.PHONY: build
build:
	@echo "---> Building"
	go build -o ./bin/$(PKG_NAME) ./cmd

.PHONY: clean
clean:
	@echo "---> Cleaning"
	rm -rf $(BIN_DIR)

.PHONY: install
install:
	@echo "---> Installing dependencies"
	go mod download

.PHONY: lint
lint:
	@echo "---> Linting"
	$(BIN_DIR)/golangci-lint run

.PHONY: setup
setup:
	@echo "--> Installing development tools"
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(BIN_DIR) v1.16.0
	go get -u $(GOTOOLS)

.PHONY: lambda
lambda:
	@echo "--> Building for lambda"
	GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o ./lotto-alerts ./cmd
	zip function.zip lotto-alerts
	rm lotto-alerts
