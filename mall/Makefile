# Go parameters
GOCMD=GO111MODULE=on go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

PROJECT_NAME := "mall"
PKG := "$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /example)

all: test build
build:
	rm -rf target/
	mkdir -p target/config
	cp config/local.yml target/config/default.yml
	$(GOBUILD) -o target/mall main.go

test:
	$(GOTEST) -v ${PKG_LIST}

clean:
	rm -rf target/
	rm -rf nohup.out

# 运行服务
run:
	nohup target/mall server -c target/config/default.yml &

# 停止服务
stop:
	pkill -f target/mall

abi:
	abigen --abi eth/Payment.abi --pkg eth --out eth/Payment.go --type Payment

#生成docker镜像，请确保已安装docker
docker:
	docker build -t mall_api:latest -f Dockerfile .

# 生成api文档
doc:
	@swag init
	echo "docs done"
	echo "see docs by: http://localhost:9052/swagger/index.html"

# 生成交互式的可视化Go程序调用图
graph:
	@echo "downloading go-callvis"
	@echo "generating graph"
	@go get -u github.com/ofabry/go-callvis
	@go-callvis ${PROJECT_NAME}

#检查代码规范
lint:
	@go get -u golang.org/x/lint/golint
	@golint -set_exit_status ${PKG_LIST}

#查看帮助
help:
	target/chat --help