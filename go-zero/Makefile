.PHONY: all build run clean tidy docker-up docker-down docker-build k8s-deploy k8s-delete

# 服务目录
SERVICES := common user-rpc auth-rpc role-rpc user-api

# Docker 镜像前缀和标签
DOCKER_REGISTRY ?= go-zero
IMAGE_TAG ?= latest

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

# 启动基础设施 (Docker Compose)
docker-up:
	docker-compose -f deploy/docker-compose.yaml up -d

# 停止基础设施 (Docker Compose)
docker-down:
	docker-compose -f deploy/docker-compose.yaml down

# 清理构建产物
clean:
	rm -rf bin/

# 生成 proto 代码
gen-user-rpc:
	cd user-rpc && protoc --go_out=. --go_opt=module=github.com/Hao-yiwen/go-examples/go-zero/user-rpc --go-grpc_out=. --go-grpc_opt=module=github.com/Hao-yiwen/go-examples/go-zero/user-rpc proto/user.proto

gen-auth-rpc:
	cd auth-rpc && protoc --go_out=. --go_opt=module=github.com/Hao-yiwen/go-examples/go-zero/auth-rpc --go-grpc_out=. --go-grpc_opt=module=github.com/Hao-yiwen/go-examples/go-zero/auth-rpc proto/auth.proto

gen-role-rpc:
	cd role-rpc && protoc --go_out=. --go_opt=module=github.com/Hao-yiwen/go-examples/go-zero/role-rpc --go-grpc_out=. --go-grpc_opt=module=github.com/Hao-yiwen/go-examples/go-zero/role-rpc proto/role.proto

gen-user-api:
	cd user-api && goctl api go -api user.api -dir .

# ===========================================
# Docker 镜像构建命令
# ===========================================

# 构建所有 Docker 镜像
docker-build: docker-build-user-rpc docker-build-auth-rpc docker-build-role-rpc docker-build-user-api
	@echo "All Docker images built successfully!"

# 构建 user-rpc 镜像
docker-build-user-rpc:
	@echo "Building user-rpc Docker image..."
	docker build -t $(DOCKER_REGISTRY)/user-rpc:$(IMAGE_TAG) -f user-rpc/Dockerfile .

# 构建 auth-rpc 镜像
docker-build-auth-rpc:
	@echo "Building auth-rpc Docker image..."
	docker build -t $(DOCKER_REGISTRY)/auth-rpc:$(IMAGE_TAG) -f auth-rpc/Dockerfile .

# 构建 role-rpc 镜像
docker-build-role-rpc:
	@echo "Building role-rpc Docker image..."
	docker build -t $(DOCKER_REGISTRY)/role-rpc:$(IMAGE_TAG) -f role-rpc/Dockerfile .

# 构建 user-api 镜像
docker-build-user-api:
	@echo "Building user-api Docker image..."
	docker build -t $(DOCKER_REGISTRY)/user-api:$(IMAGE_TAG) -f user-api/Dockerfile .

# 推送所有 Docker 镜像
docker-push: docker-push-user-rpc docker-push-auth-rpc docker-push-role-rpc docker-push-user-api
	@echo "All Docker images pushed successfully!"

docker-push-user-rpc:
	docker push $(DOCKER_REGISTRY)/user-rpc:$(IMAGE_TAG)

docker-push-auth-rpc:
	docker push $(DOCKER_REGISTRY)/auth-rpc:$(IMAGE_TAG)

docker-push-role-rpc:
	docker push $(DOCKER_REGISTRY)/role-rpc:$(IMAGE_TAG)

docker-push-user-api:
	docker push $(DOCKER_REGISTRY)/user-api:$(IMAGE_TAG)

# ===========================================
# Kubernetes 部署命令
# ===========================================

# 部署到 Kubernetes (使用 Kustomize)
k8s-deploy:
	@echo "Deploying to Kubernetes..."
	kubectl apply -k deploy/k8s/

# 从 Kubernetes 删除
k8s-delete:
	@echo "Deleting from Kubernetes..."
	kubectl delete -k deploy/k8s/ --ignore-not-found=true

# 仅部署基础设施 (MySQL + Redis)
k8s-deploy-infra:
	@echo "Deploying infrastructure..."
	kubectl apply -f deploy/k8s/namespace.yaml
	kubectl apply -f deploy/k8s/secret.yaml
	kubectl apply -f deploy/k8s/mysql.yaml
	kubectl apply -f deploy/k8s/redis.yaml

# 部署微服务 (需要先部署基础设施)
k8s-deploy-services:
	@echo "Deploying microservices..."
	kubectl apply -f deploy/k8s/configmap.yaml
	kubectl apply -f deploy/k8s/user-rpc.yaml
	kubectl apply -f deploy/k8s/auth-rpc.yaml
	kubectl apply -f deploy/k8s/role-rpc.yaml
	kubectl apply -f deploy/k8s/user-api.yaml
	kubectl apply -f deploy/k8s/ingress.yaml

# 查看部署状态
k8s-status:
	@echo "=== Namespaces ==="
	kubectl get ns go-zero
	@echo ""
	@echo "=== Pods ==="
	kubectl get pods -n go-zero
	@echo ""
	@echo "=== Services ==="
	kubectl get svc -n go-zero
	@echo ""
	@echo "=== Ingress ==="
	kubectl get ingress -n go-zero

# 查看 Pod 日志
k8s-logs-user-api:
	kubectl logs -n go-zero -l app=user-api -f --tail=100

k8s-logs-user-rpc:
	kubectl logs -n go-zero -l app=user-rpc -f --tail=100

k8s-logs-auth-rpc:
	kubectl logs -n go-zero -l app=auth-rpc -f --tail=100

k8s-logs-role-rpc:
	kubectl logs -n go-zero -l app=role-rpc -f --tail=100

# 重启部署
k8s-restart:
	kubectl rollout restart deployment -n go-zero user-api user-rpc auth-rpc role-rpc

# 端口转发 (用于本地测试)
k8s-port-forward:
	@echo "Forwarding user-api to localhost:8888..."
	kubectl port-forward -n go-zero svc/user-api-svc 8888:8888

# ===========================================
# 完整部署流程
# ===========================================

# 完整部署: 构建镜像 + 部署到 K8s
deploy-all: docker-build k8s-deploy
	@echo "Deployment completed!"

# 帮助信息
help:
	@echo "可用命令:"
	@echo ""
	@echo "  开发命令:"
	@echo "  make tidy           - 整理所有模块Go依赖"
	@echo "  make sync           - 同步 go.work"
	@echo "  make build          - 构建所有服务"
	@echo "  make docker-up      - 启动本地 MySQL 和 Redis (Docker Compose)"
	@echo "  make docker-down    - 停止本地 MySQL 和 Redis"
	@echo "  make run-user-rpc   - 运行用户 RPC 服务"
	@echo "  make run-auth-rpc   - 运行认证 RPC 服务"
	@echo "  make run-role-rpc   - 运行角色 RPC 服务"
	@echo "  make run-user-api   - 运行 API 网关服务"
	@echo "  make clean          - 清理构建产物"
	@echo ""
	@echo "  Docker 命令:"
	@echo "  make docker-build           - 构建所有 Docker 镜像"
	@echo "  make docker-build-user-api  - 构建 user-api 镜像"
	@echo "  make docker-build-user-rpc  - 构建 user-rpc 镜像"
	@echo "  make docker-build-auth-rpc  - 构建 auth-rpc 镜像"
	@echo "  make docker-build-role-rpc  - 构建 role-rpc 镜像"
	@echo "  make docker-push            - 推送所有镜像到镜像仓库"
	@echo ""
	@echo "  Kubernetes 命令:"
	@echo "  make k8s-deploy           - 部署到 Kubernetes"
	@echo "  make k8s-delete           - 从 Kubernetes 删除"
	@echo "  make k8s-deploy-infra     - 仅部署基础设施 (MySQL + Redis)"
	@echo "  make k8s-deploy-services  - 部署微服务"
	@echo "  make k8s-status           - 查看部署状态"
	@echo "  make k8s-logs-user-api    - 查看 user-api 日志"
	@echo "  make k8s-restart          - 重启所有微服务"
	@echo "  make k8s-port-forward     - 端口转发到本地"
	@echo ""
	@echo "  完整流程:"
	@echo "  make deploy-all     - 构建镜像并部署到 K8s"
	@echo ""
	@echo "  环境变量:"
	@echo "  DOCKER_REGISTRY     - Docker 镜像仓库前缀 (默认: go-zero)"
	@echo "  IMAGE_TAG           - 镜像标签 (默认: latest)"
