# ---------- Import make config
conf ?= Makefile.conf
include ${conf}
export ${shell sed 's/=.*//' ${conf}}

# ---------- Help Menu
.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' ${MAKEFILE_LIST} | sort

.DEFAULT_GOAL := help

# ***************** Generator targets *********************

pb-codegen: ## Generate code out of proto files
	protoc -I ${INCLUDE_PROTO_DIR} ${INPUT_PROTO_FILES} --go_out=plugins=grpc:${PROTO_OUTPUT_DIR}

# ***************** build targets *********************

all: ## build and test
	test build

test: ## build and test
	${GOTEST} -v ./...

clean: ## delete binary
	${GOCLEAN}
	rm -f ${BINARY}

deps: ## cleanup and reinstall
	${GO} mod tidy
	${GO} mod download

build: ## build linux binary
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${GO} build -o ${BINARY} ./...

# Cross compilation
build-local: ## build mac binary
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 ${GO} build -o ${BINARY} ./...

docker-build: ## docker build
	docker build -f ${DOCKER_FILE} -t ${BINARY}:latest .

# **************************************
