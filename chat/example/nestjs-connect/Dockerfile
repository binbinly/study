# 引用镜像
FROM node:lts-alpine
# 作者
LABEL maintainer="chat"

# 执行镜像的工作目录
WORKDIR /var/www/chat

# 将nestapi目录拷贝到镜像里，也可用COPY命令
COPY . .

RUN npm i --production

# 配置系统变量，指定端口
ENV HTTP_PORT=3005 WS_PORT=3010 NODE_ENV=prod \
  REDIS_HOST=127.0.0.1 CONSUL_HOST=127.0.0.1 \
  SERVICE_HOST=127.0.0.1 \
  SERVICE_NAME=chat-connect SERVICE_ID=conn-1 SERVICE_PORT=20005

# 开放端口
EXPOSE ${HTTP_PORT} ${WS_PORT} ${SERVICE_PORT}

# 容器启动命令
CMD ["node","dist/main.js"]