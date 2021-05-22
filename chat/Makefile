# Go parameters
GOCMD=GO111MODULE=on go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

all: test build
build:
	rm -rf target/
	mkdir -p target/config
	cp config/connect.default.yaml target/config/connect.yaml
	cp config/task.default.yaml target/config/task.yaml
	cp config/logic.default.yaml target/config/logic.yaml
	$(GOBUILD) -o target/chat main.go

test:
	$(GOTEST) -v ./...

clean:
	rm -rf target/
	rm -rf nohup.out

# 运行服务
run:
	nohup target/chat connect -c target/config/connect.yaml &
	nohup target/chat task -c target/config/task.yaml &
	nohup target/chat logic -c target/config/logic.yaml &

# 停止服务
stop:
	pkill -f target/chat

# 生成api文档
gen-docs:
	@swag init
	echo "gen-docs done"
	echo "see docs by: http://localhost:9050/swagger/index.html"

# 生成交互式的可视化Go程序调用图
gen-graph:
	echo "downloading go-callvis"
	echo "generating graph"
	@go-callvis chat

# 生成ca证书
ca:
	openssl req -new -nodes -x509 -out target/cert/server.crt -keyout target/cert/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=127.0.0.1/emailAddress=xxxxx@qq.com"

help:
	shell target/chat --help