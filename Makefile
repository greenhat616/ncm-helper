PROJECT_NAME := "ncm-helper"
PROJECT_PATH := "github.com/a632079/ncm-helper"
PKG := "$(PROJECT_PATH)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

.PHONY: all dep get-tools lint vet test test-coverage build clean

all:
	build

get-tools:
	@echo Installing tools...
	go get -u github.com/mgechev/revive

dep: # get dependencies
	@echo Installing Dependencies...
	go mod download

lint: get-tools ## Lint Golang files
	@echo
	@echo Linting go codes with revive...
	@revive -config ./revive.toml -formatter friendly ${PKG_LIST}

vet:
	@echo Linting go codes with go vet...
	go vet ./...

build: dep
	@echo;
	@echo Building...;
	@mkdir -p dist;
	go build -v -o dist/${PROJECT_NAME} .;

test:
	@echo Testing...
	@go test -short ${PKG_LIST}

test-coverage:
	@go test -short -coverprofile cover.out -covermode=atomic ${PKG_LIST}
	@cat cover.out >> coverage.txt

clean:
	@rm -f coverage.txt
	@rm -f cover.out
