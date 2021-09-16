FROM golang:1.15-alpine

# 作者
LABEL maintainer="mall"

# 镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct" \
    TZ=Asia/Shanghai

# 执行镜像的工作目录
WORKDIR /go/src/mall

# 将目录拷贝到镜像里
COPY . .

RUN go build -o mall .

EXPOSE 9052

# 启动执行命令
ENTRYPOINT ["/go/src/mall/mall"]

# 1. build image: docker build -t mall:latest -f Dockerfile .
# 2. start: docker run --rm -it -p 9052:9052 mall:latest
# 启动时设置 --rm 选项，这样在容器退出时就能够自动清理容器内部的文件系统