# Go parameters
GOCMD=GO111MODULE=on go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

PROJECT_NAME := "chat"
PKG := "$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /example)

all: test build
build:
	rm -rf target/
	mkdir -p target/config
	cp config/connect/default.yaml target/config/connect.yaml
	cp config/task/default.yaml target/config/task.yaml
	cp config/chat/default.yaml target/config/logic.yaml
	cp config/cebter/default.yaml target/config/center.yaml
	$(GOBUILD) -o target/chat main.go

test:
	$(GOTEST) -v ${PKG_LIST}

clean:
	rm -rf target/
	rm -rf nohup.out

# 运行服务
run:
	nohup target/chat connect -c target/config/connect.yaml &
	nohup target/chat task -c target/config/task.yaml &
	nohup target/chat chat -c target/config/logic.yaml &
	nohup target/chat center -c target/config/center.yaml &

# 停止服务
stop:
	pkill -f target/chat

#生成docker镜像，请确保已安装docker
docker:
	docker build -t chat_app:latest -f Dockerfile .

# 生成api文档
doc:
	@swag init
	echo "docs done"
	echo "see docs by: http://localhost:9050/swagger/index.html"

# 生成交互式的可视化Go程序调用图
graph:
	@echo "downloading go-callvis"
	@echo "generating graph"
	@go get -u github.com/ofabry/go-callvis
	@go-callvis ${PROJECT_NAME}

# 生成ca证书
ca:
	openssl req -new -nodes -x509 -out target/cert/server.crt -keyout target/cert/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=127.0.0.1/emailAddress=xxxxx@qq.com"

#检查代码规范
lint:
	@go get -u golang.org/x/lint/golint
	@golint -set_exit_status ${PKG_LIST}

#查看帮助
help:
	target/chat --help