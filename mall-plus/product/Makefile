# Go parameters
PROJECT_NAME := "product"
PKG := "$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /example)

GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install github.com/golang/protobuf/protoc-gen-go@latest
	@go install github.com/asim/go-micro/cmd/protoc-gen-micro/v4@latest

.PHONY: dev
dev:
	go run main.go --config local.yaml \
		--registry consul --registry_address 192.168.8.76:8500 \
		--tracer_address http://192.168.8.76:14268/api/traces

.PHONY: build
build:
	go build -o product *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t product:latest

#检查代码规范
.PHONY: lint
lint:
	@go get -u golang.org/x/lint/golint
	@golint -set_exit_status ${PKG_LIST}