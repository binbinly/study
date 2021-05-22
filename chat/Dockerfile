FROM golang:1.15-alpine

# 作者
LABEL maintainer="chat"

# 镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct" \
    TZ=Asia/Shanghai

# 执行镜像的工作目录
WORKDIR /go/src/chat

# 将目录拷贝到镜像里，也可用COPY命令
COPY . .

RUN go build -o chat .

# 开放http ws端口
EXPOSE 9050 9070

# 启动执行命令
ENTRYPOINT ["/go/src/chat/chat"]

# 1. build image: docker build -t chat:latest -f Dockerfile .
# 2. start: docker run --rm -it -p 9050:9050 -p 9070:9070 chat:latest
# 启动时设置 --rm 选项，这样在容器退出时就能够自动清理容器内部的文件系统