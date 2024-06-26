FROM golang:1.21-alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download


# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件 bluebell_app
RUN go build -o dog_app .

###################
# 接下来创建一个小镜像
###################
FROM debian:buster-slim

COPY ./conf /conf
COPY ./.env /
COPY ./wait-for.sh /

# 从builder镜像中把/build/app 拷贝到当前目录
COPY --from=builder /build/dog_app /

RUN set -eux; \
	apt-get update; \
	apt-get install -y \
		--no-install-recommends \
		netcat; \
        chmod 755 wait-for.sh

# 安装证书
RUN apt-get install -y --no-install-recommends ca-certificates

# 声明服务端口
EXPOSE 9001