NAME="github.com/odpf/kay"

.PHONY: all build test clean dist vet proto install

all: test clean build

build: ## Build the kay binary
	@echo "> Building kay version ${APP_VERSION}..."
	go build -o kay .
	@echo "- Build complete"

test: ## Run the tests
	@echo "> Running tests..."
	go test ./... -coverprofile=coverage.out

coverage: ## Print code coverage
	go test -race -coverprofile coverage.txt -covermode=atomic ./... & go tool cover -html=coverage.out

vet: ## Run the go vet tool
	go vet ./...

clean: ## Clean the build artifacts
	@echo "> Cleaning artifacts..."
	rm -rf kay dist/

help: ## Display this help message
	@cat $(MAKEFILE_LIST) | grep -e "^[a-zA-Z_\-]*: *.*## *" | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'