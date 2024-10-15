# 使用 golang 官方镜像作为基础镜像
FROM golang:1.23 AS builder

# 设置工作目录
WORKDIR /app

# 复制 go 模块和 sum 文件
COPY go.mod go.sum ./

# 下载所有依赖
RUN go mod download

# 复制源代码到容器中
COPY . .

RUN ls -l .

# 构建应用
RUN CGO_ENABLED=0 go build -o main .

# 使用 scratch 作为基础镜像
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 从 builder 阶段复制构建好的二进制文件
COPY --from=builder /app/main .

RUN ls -l /root/

# 开放端口
EXPOSE 8080

# 运行二进制文件
CMD ["/root/main"]
