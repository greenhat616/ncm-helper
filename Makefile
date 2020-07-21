PROJECT_NAME := "ncm-helper"
PROJECT_PATH := "github.com/a632079/ncm-helper"
PKG := "$(PROJECT_PATH)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

.PHONY: all dep get-tools lint vet test test-coverage build clean

all:
	build

get-tools:
	go get -u golang.org/x/lint/golint

dep: # get dependence
	@go mod download

lint: ## Lint Golang files
	@golint -set_exit_status ${PKG_LIST}

build: dep
	@echo Building...
	@go build -v .

test:
	@echo Testing...
	@go test -short ${PKG_LIST}

test-coverage:
	@go test -short -coverprofile cover.out -covermode=atomic ${PKG_LIST}
	@cat cover.out >> coverage.txt

clean:
	rm -f coverage.txt
	rm -f cover.out
