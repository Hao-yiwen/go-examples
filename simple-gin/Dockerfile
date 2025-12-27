# 构建阶段
FROM golang:1.25-alpine AS builder

# 安装必要的构建工具
RUN apk add --no-cache git ca-certificates tzdata

# 设置工作目录
WORKDIR /app

# 复制依赖文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 安装 swag 工具
RUN go install github.com/swaggo/swag/cmd/swag@latest

# 复制源代码
COPY . .

# 生成 Swagger 文档
RUN swag init -g cmd/simple-gin/main.go -o docs

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/bin/simple-gin ./cmd/simple-gin

# 运行阶段
FROM alpine:3.19

# 安装必要的运行时依赖
RUN apk add --no-cache ca-certificates tzdata

# 设置时区
ENV TZ=Asia/Shanghai

# 创建非 root 用户
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# 设置工作目录
WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /app/bin/simple-gin .

# 复制配置文件
COPY --from=builder /app/configs ./configs

# 切换到非 root 用户
USER appuser

# 暴露端口
EXPOSE 8080

# 启动命令
CMD ["./simple-gin"]
