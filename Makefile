ifneq ("$(wildcard $(./config.d/.env))","")
	include ./config.d/.env
	export $(shell sed 's/=.*//' ./config.d/.env)
endif


SERVICE_NAME = $(shell basename "$(PWD)")

ROOT = $(shell "$(PWD)")
GO ?= go
OS = $(shell uname -s | tr A-Z a-z)
export GOBIN = ${ROOT}bin

PATH := "$(PATH)":"$(GOBIN)"

LINT = ${GOBIN}/golangci-lint
LINT_DOWNLOAD = curl --progress-bar -SfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s latest

TPARSE = $(GOBIN)/tparse
TPARSE_DOWNLOAD = $(GO) get -u github.com/mfridman/tparse

COMPILEDEAMON = ${GOBIN}/CompileDaemon
COMPILEDEAMON_DOWNLOAD = $(GO) get -u github.com/githubnemo/CompileDaemon

SWAG = ${GOBIN}/swag
SWAG_DOWNLOAD = $(GO) get -u github.com/swaggo/swag/cmd/swag


.PHONY: help
help: ## Display this help message
	@ cat $(MAKEFILE_LIST) | grep -e "^[a-zA-Z_\-]*: *.*## *" | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


.PHONY: release
release:create-env ## Build production binary file
	@ $(GO) build -a -o ./bin/${SERVICE_NAME} ./cmd/...


.PHONY: build
build:create-env ## Build development binary file
	@ $(GO) build -o ./bin/${SERVICE_NAME} ./cmd/...


.PHONY: run
run:create-env ## run as development reload if code changes
	@ test -e $(COMPILEDEAMON) || $(COMPILEDEAMON_DOWNLOAD)
	@ chmod +x $(COMPILEDEAMON)
	@ $(COMPILEDEAMON) --build="make build" --command="$(GOBIN)/$(SERVICE_NAME)"


.PHONY: create-env
create-env: ## Create .env and config.yml file
	@ test -e ./config.d/.env && echo ./config.d/.env exists || cp ./config.d/.env.example ./config.d/.env
	@ test -e ./config.d/config.yml && echo ./config.d/config.yml exists || cp ./config.d/config.yml.example ./config.d/config.yml


.PHONY: mod
mod: ## Get dependency packages
	@ $(GO) mod tidy


.PHONY: test
test:create-env ## Run unit tests
	echo $(TPARSE)
	@ test -e $(TPARSE) || $(TPARSE_DOWNLOAD)
	@ $(GO) test -failfast -count=1 ./... -json -cover | $(TPARSE) -all -smallscreen


.PHONY: race
race:create-env ## Run data race detector
	@ test -e $(TPARSE) || $(TPARSE_DOWNLOAD)
	@ $(GO) test -short -race ./... -json -cover | $(TPARSE) -all -smallscreen


.PHONY: coverage
coverage:create-env ## check coverage test code
	@ $(GO) test ./... -coverprofile=coverage.out
	@ $(GO) tool cover -func=coverage.out
	@ $(GO) tool cover -html=coverage.out -o coverage.html;


.PHONY: lint
lint: ## Lint the files
	@ test -e $(LINT) || $(LINT_DOWNLOAD)
	@ $(LINT) version
	@ $(LINT) --timeout 10m run


.PHONY: docs
docs: ## Create/Update documents using swagger tool
	@ test -e  $(SWAG) || $(SWAG_DOWNLOAD)
	@ swag init -g ./cmd/main.go -o ./docs --parseDependency


.PHONY: docker-build
docker-build: ## Build docker-compose
	@ docker-compose build --no-cache


.PHONY: docker-up
docker-up: ## Run with docker-compose auto reload
	@ docker-compose up -d


.PHONY: docker-down
docker-down: ## Stop docker-compose
	@ docker-compose down


.PHONY: docker-log
docker-log: ## Print docker log
	@ docker-compose logs --tail=300 -f
