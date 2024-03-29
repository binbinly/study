FROM golang:1.16-alpine as build

# 作者
LABEL maintainer="mall"

# 镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct" \
    TZ=Asia/Shanghai \
    MICRO_SERVER_ADDRESS=:50200 \
    MICRO_CONFIG="/app/docker.yaml" \
    MICRO_REGISTRY=consul \
    MICRO_REGISTRY_ADDRESS="127.0.0.1:8500" \
    MICRO_TRACER_ADDRESS="http://127.0.0.1:14268/api/traces"

# 执行镜像的工作目录
WORKDIR /go/src/mall

# 将目录拷贝到镜像里
COPY center/ ./center
COPY common/ ./common
COPY pkg/ ./pkg

RUN cd center && go mod tidy && go build -o center .

# 引入alphine最小linux镜像
FROM alpine

WORKDIR /app
# 复制生成的可执行命令和一些配置文件
COPY --from=build /go/src/mall/center/center .
COPY --from=build /go/src/mall/center/docker.yaml .

EXPOSE 50200

# 启动执行命令
ENTRYPOINT ["/app/center"]

# 1. build image: docker build -t mall/center:latest -f Dockerfile .
# 2. start: docker run --rm -it mall/center:latest
# 启动时设置 --rm 选项，这样在容器退出时就能够自动清理容器内部的文件系统