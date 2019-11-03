BIN_DIR  ?= ./bin
PKG_NAME ?= lotto-alerts
# LDFLAGS  ?= "-s -w -X github.com/lob/captain/pkg/cmd.captainVersion=$(VERSION) -X github.com/lob/captain/pkg/docker.dockerVersion=$(DOCKER_VERSION)"
# VERSION  ?=
# DOCKER_VERSION ?= 1.32

# COVERAGE_PROFILE ?= coverage.out

GOTOOLS := \
golang.org/x/tools/cmd/cover \

default: build

.PHONY: build
build:
	@echo "---> Building"
	go build -o ./bin/$(PKG_NAME) ./cmd
	# go build -o ./bin/$(PKG_NAME) -ldflags $(LDFLAGS) ./cmd

.PHONY: clean
clean:
	@echo "---> Cleaning"
	rm -rf $(BIN_DIR)

# .PHONY: enforce
# enforce:
# 	@echo "---> Enforcing coverage"
# 	./scripts/coverage.sh $(COVERAGE_PROFILE)

# .PHONY: html
# html:
# 	@echo "---> Generating HTML coverage report"
# 	go tool cover -html $(COVERAGE_PROFILE)

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

.PHONY: test
test:
	@echo "---> Testing"
	ENVIRONMENT=test go test ./pkg/... -race -coverprofile $(COVERAGE_PROFILE)

# .PHONY: release
# release:
# 	@echo "---> Creating tagged release"
# ifndef VERSION
# 	$(error VERSION must be specified)
# endif
# 	git tag $(VERSION)
# 	# Check that the commit is tagged and starts with "v".
# 	[[ $$(git tag -l --points-at HEAD) == v* ]]
# 	git push origin
# 	git push origin --tags
