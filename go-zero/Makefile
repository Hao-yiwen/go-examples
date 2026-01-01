.PHONY: all build run clean tidy docker-up docker-down

# 服务目录
SERVICES := common user-rpc auth-rpc role-rpc user-api

# 默认目标
all: tidy build

# 整理所有模块依赖
tidy:
	@for dir in $(SERVICES); do \
		echo "Tidying $$dir..."; \
		cd $$dir && go mod tidy && cd ..; \
	done

# 同步 workspace
sync:
	go work sync

# 构建所有服务
build: build-user-rpc build-auth-rpc build-role-rpc build-user-api

build-user-rpc:
	go build -o bin/user-rpc ./user-rpc/user.go

build-auth-rpc:
	go build -o bin/auth-rpc ./auth-rpc/auth.go

build-role-rpc:
	go build -o bin/role-rpc ./role-rpc/role.go

build-user-api:
	go build -o bin/user-api ./user-api/user.go

# 运行服务（需要在不同终端中分别运行）
run-user-rpc:
	go run ./user-rpc/user.go -f ./user-rpc/etc/user.yaml

run-auth-rpc:
	go run ./auth-rpc/auth.go -f ./auth-rpc/etc/auth.yaml

run-role-rpc:
	go run ./role-rpc/role.go -f ./role-rpc/etc/role.yaml

run-user-api:
	go run ./user-api/user.go -f ./user-api/etc/user-api.yaml

# 启动基础设施
docker-up:
	docker-compose -f deploy/docker-compose.yaml up -d

# 停止基础设施
docker-down:
	docker-compose -f deploy/docker-compose.yaml down

# 清理构建产物
clean:
	rm -rf bin/

# 生成 proto 代码
gen-user-rpc:
	cd user-rpc && goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=.

gen-auth-rpc:
	cd auth-rpc && goctl rpc protoc auth.proto --go_out=. --go-grpc_out=. --zrpc_out=.

gen-role-rpc:
	cd role-rpc && goctl rpc protoc role.proto --go_out=. --go-grpc_out=. --zrpc_out=.

gen-user-api:
	cd user-api && goctl api go -api user.api -dir .

# 帮助信息
help:
	@echo "可用命令:"
	@echo "  make tidy        - 整理所有模块Go依赖"
	@echo "  make sync        - 同步 go.work"
	@echo "  make build       - 构建所有服务"
	@echo "  make docker-up   - 启动MySQL和Redis"
	@echo "  make docker-down - 停止MySQL和Redis"
	@echo "  make run-user-rpc  - 运行用户RPC服务"
	@echo "  make run-auth-rpc  - 运行认证RPC服务"
	@echo "  make run-role-rpc  - 运行角色RPC服务"
	@echo "  make run-user-api  - 运行API网关服务"
	@echo "  make clean       - 清理构建产物"
