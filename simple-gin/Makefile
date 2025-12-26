# Simple Gin Makefile
# 使用方法: make help

.PHONY: all build run dev test test-v test-cover bench clean deps docs lint fmt help

# 变量定义
APP_NAME := simple-gin
BUILD_DIR := bin
MAIN_FILE := ./cmd/simple-gin
DOCS_DIR := docs

# 默认目标
all: deps docs build

# ==================== 构建相关 ====================

## build: 编译项目
build:
	@echo ">>> 编译项目..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)
	@echo ">>> 编译完成: $(BUILD_DIR)/$(APP_NAME)"

## run: 直接运行（不编译到文件）
run:
	go run $(MAIN_FILE)

## dev: 开发模式（生成文档 + 运行）
dev: docs run

# ==================== 测试相关 ====================

## test: 运行所有测试
test:
	go test ./...

## test-v: 运行所有测试（详细输出）
test-v:
	go test -v ./...

## test-cover: 运行测试并显示覆盖率
test-cover:
	go test -cover ./...

## test-cover-html: 生成覆盖率 HTML 报告
test-cover-html:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo ">>> 覆盖率报告: coverage.html"

## test-unit: 只运行单元测试（pkg）
test-unit:
	go test -v ./pkg/...

## test-integration: 只运行集成测试
test-integration:
	go test -v ./test/integration/...

## bench: 运行性能测试
bench:
	go test -bench=. -benchmem ./pkg/...

# ==================== 文档相关 ====================

## docs: 生成 Swagger 文档
docs:
	@echo ">>> 生成 Swagger 文档..."
	@which swag > /dev/null || (echo "安装 swag..." && go install github.com/swaggo/swag/cmd/swag@latest)
	swag init -g cmd/simple-gin/main.go -o $(DOCS_DIR)
	@echo ">>> 文档生成完成: $(DOCS_DIR)/"

## docs-fmt: 格式化 Swagger 注释
docs-fmt:
	swag fmt

# ==================== 依赖相关 ====================

## deps: 下载依赖
deps:
	go mod download

## deps-update: 更新依赖
deps-update:
	go get -u ./...
	go mod tidy

## deps-tidy: 清理依赖
deps-tidy:
	go mod tidy

# ==================== 代码质量 ====================

## fmt: 格式化代码
fmt:
	go fmt ./...

## lint: 代码检查（需要安装 golangci-lint）
lint:
	@which golangci-lint > /dev/null || (echo "请安装 golangci-lint: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest" && exit 1)
	golangci-lint run

## vet: go vet 检查
vet:
	go vet ./...

# ==================== 清理相关 ====================

## clean: 清理构建产物
clean:
	@echo ">>> 清理构建产物..."
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html
	@echo ">>> 清理完成"

## clean-all: 清理所有（包括文档）
clean-all: clean
	rm -rf $(DOCS_DIR)

# ==================== 工具安装 ====================

## tools: 安装开发工具
tools:
	@echo ">>> 安装开发工具..."
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/air-verse/air@latest
	@echo ">>> 工具安装完成"

# ==================== Docker 相关（可选）====================

## docker-build: 构建 Docker 镜像
docker-build:
	docker build -t $(APP_NAME):latest .

## docker-run: 运行 Docker 容器
docker-run:
	docker run -p 8080:8080 $(APP_NAME):latest

# ==================== 帮助信息 ====================

## help: 显示帮助信息
help:
	@echo ""
	@echo "Simple Gin - 可用命令:"
	@echo ""
	@echo "构建:"
	@echo "  make build          - 编译项目到 bin/"
	@echo "  make run            - 直接运行项目"
	@echo "  make dev            - 开发模式（生成文档 + 运行）"
	@echo ""
	@echo "测试:"
	@echo "  make test           - 运行所有测试"
	@echo "  make test-v         - 运行测试（详细输出）"
	@echo "  make test-cover     - 显示测试覆盖率"
	@echo "  make test-cover-html- 生成覆盖率 HTML 报告"
	@echo "  make test-unit      - 只运行单元测试"
	@echo "  make test-integration - 只运行集成测试"
	@echo "  make bench          - 运行性能测试"
	@echo ""
	@echo "文档:"
	@echo "  make docs           - 生成 Swagger 文档"
	@echo "  make docs-fmt       - 格式化 Swagger 注释"
	@echo ""
	@echo "依赖:"
	@echo "  make deps           - 下载依赖"
	@echo "  make deps-update    - 更新依赖"
	@echo "  make deps-tidy      - 清理依赖"
	@echo ""
	@echo "代码质量:"
	@echo "  make fmt            - 格式化代码"
	@echo "  make lint           - 代码检查"
	@echo "  make vet            - go vet 检查"
	@echo ""
	@echo "清理:"
	@echo "  make clean          - 清理构建产物"
	@echo "  make clean-all      - 清理所有（包括文档）"
	@echo ""
	@echo "工具:"
	@echo "  make tools          - 安装开发工具（swag, lint, air）"
	@echo ""
