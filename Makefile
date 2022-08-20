NAME=github.com/odpf/kay
VERSION=$(shell git describe --tags --always --first-parent 2>/dev/null)
COMMIT=$(shell git rev-parse --short HEAD)
BUILD_TIME=$(shell date)
COVERAGE_DIR=coverage
BUILD_DIR=dist
EXE=kay
PROTON_COMMIT := "f1b701ccaac3cb0a27b99cae6cd7a9e32188673d"

.PHONY: all build clean tidy format test test-coverage proto

all: clean test build format lint

tidy: ## Tidy up the code
	@echo "Tidy up go.mod..."
	@go mod tidy -v

install: ## Install the binary
	@echo "Installing Kay to ${GOBIN}..."
	@go install
	
format: ## Format the code
	@echo "Running gofumpt..."
	@gofumpt -l -w .

lint: ## Lint the code
	@echo "Running lint checks using golangci-lint..."
	@golangci-lint run

clean: tidy ## Clean up the build artifacts
	@echo "Cleaning up build directories..."
	@rm -rf ${COVERAGE_DIR} ${BUILD_DIR}
	@echo "Running go-generate..."
	@go generate ./...

test: tidy ## Run the tests
	@mkdir -p ${COVERAGE_DIR}
	@echo "Running unit tests..."
	@go test ./... -coverprofile=${COVERAGE_DIR}/coverage.out

test-coverage: test ## Run the tests with coverage
	@echo "Generating coverage report..."
	@go tool cover -html=${COVERAGE_DIR}/coverage.out

build: clean ## Build the binary
	@mkdir -p ${BUILD_DIR}
	@echo "Running build for '${VERSION}' in '${BUILD_DIR}/'..."
	@CGO_ENABLED=0 go build -ldflags '-X "${NAME}/pkg/version.Version=${VERSION}" -X "${NAME}/pkg/version.Commit=${COMMIT}" -X "${NAME}/pkg/version.BuildTime=${BUILD_TIME}"' -o ${BUILD_DIR}/${EXE}

download: ## Download go dependencies
	@go mod download

proto: ## Generate the protobuf files
	@echo " > generating protobuf from odpf/proton"
	@echo " > [info] make sure correct version of dependencies are installed using 'make setup'"
	@buf generate https://github.com/odpf/proton/archive/${PROTON_COMMIT}.zip#strip_components=1 --template buf.gen.yaml --path odpf/kay
	@echo " > protobuf compilation finished"

setup: ## Install dependencies
	@echo "> installing dependencies"
	@go mod tidy
	@go install github.com/vektra/mockery/v2@v2.12.2
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.0
	@go get google.golang.org/protobuf/proto@v1.28.0
	@go get google.golang.org/grpc@v1.46.0
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.9.0
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.9.0
	@go install github.com/bufbuild/buf/cmd/buf@v1.4.0
	@go install github.com/envoyproxy/protoc-gen-validate@v0.6.7

help: ## Display this help message
	@cat $(MAKEFILE_LIST) | grep -e "^[a-zA-Z_\-]*: *.*## *" | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'



