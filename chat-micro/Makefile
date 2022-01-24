
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -d github.com/envoyproxy/protoc-gen-validate
	mkdir doc

.PHONY: doc
doc:
	go install github.com/swaggo/swag/cmd/swag@latest
	@swag init
	echo "docs done"
	echo "see docs by: http://localhost:9050/swagger/index.html"

.PHONY: proto
proto:
	protoc -I ./proto \
    	  --go_out ./proto --go_opt paths=source_relative \
    	  --validate_out="lang=go,paths=source_relative:./proto" \
    	  --go-grpc_out ./proto --go-grpc_opt paths=source_relative \
    	  proto/*/*.proto

.PHONY: test
test:
	go test -v ./... -cover

