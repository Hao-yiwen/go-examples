# Go-Zero 微服务 Kubernetes 部署指南

本文档介绍如何将 Go-Zero 微服务用户管理系统部署到 Kubernetes 集群。

## 目录

- [前提条件](#前提条件)
- [快速部署](#快速部署)
- [详细部署步骤](#详细部署步骤)
- [配置说明](#配置说明)
- [生产环境部署](#生产环境部署)
- [常用命令](#常用命令)
- [故障排查](#故障排查)

## 前提条件

### 必需工具

- **kubectl**: Kubernetes 命令行工具
- **Docker**: 用于构建镜像
- **Kubernetes 集群**: minikube、k3s、或云服务商提供的 K8s 集群

### 可选工具

- **kustomize**: K8s 配置管理工具（kubectl 已内置）
- **Helm**: 包管理工具

### 验证环境

```bash
# 检查 kubectl 连接
kubectl cluster-info

# 检查节点状态
kubectl get nodes
```

## 快速部署

### 一键部署（开发环境）

```bash
# 1. 构建所有 Docker 镜像
make docker-build

# 2. 部署到 Kubernetes
make k8s-deploy

# 3. 查看部署状态
make k8s-status
```

### 验证部署

```bash
# 等待所有 Pod 就绪
kubectl wait --for=condition=Ready pods --all -n go-zero --timeout=300s

# 端口转发测试
make k8s-port-forward
# 在另一个终端测试
curl http://localhost:8888/api/user/list
```

## 详细部署步骤

### 步骤 1: 构建 Docker 镜像

```bash
# 构建所有服务镜像
make docker-build

# 或单独构建
make docker-build-user-api
make docker-build-user-rpc
make docker-build-auth-rpc
make docker-build-role-rpc
```

如果使用远程镜像仓库：

```bash
# 设置镜像仓库地址
export DOCKER_REGISTRY=your-registry.com/go-zero
export IMAGE_TAG=v1.0.0

# 构建并推送
make docker-build
make docker-push
```

### 步骤 2: 创建命名空间和基础资源

```bash
# 创建命名空间
kubectl apply -f deploy/k8s/namespace.yaml

# 创建 Secret（包含敏感配置）
kubectl apply -f deploy/k8s/secret.yaml

# 创建 ConfigMap
kubectl apply -f deploy/k8s/configmap.yaml
```

### 步骤 3: 部署基础设施

```bash
# 部署 MySQL
kubectl apply -f deploy/k8s/mysql.yaml

# 部署 Redis
kubectl apply -f deploy/k8s/redis.yaml

# 等待基础设施就绪
kubectl wait --for=condition=Ready pods -l app=mysql -n go-zero --timeout=120s
kubectl wait --for=condition=Ready pods -l app=redis -n go-zero --timeout=60s
```

### 步骤 4: 部署微服务

```bash
# 部署 RPC 服务
kubectl apply -f deploy/k8s/user-rpc.yaml
kubectl apply -f deploy/k8s/auth-rpc.yaml
kubectl apply -f deploy/k8s/role-rpc.yaml

# 等待 RPC 服务就绪
kubectl wait --for=condition=Ready pods -l app=user-rpc -n go-zero --timeout=60s
kubectl wait --for=condition=Ready pods -l app=auth-rpc -n go-zero --timeout=60s
kubectl wait --for=condition=Ready pods -l app=role-rpc -n go-zero --timeout=60s

# 部署 API 网关
kubectl apply -f deploy/k8s/user-api.yaml
```

### 步骤 5: 配置外部访问

```bash
# 部署 Ingress（需要 Ingress Controller）
kubectl apply -f deploy/k8s/ingress.yaml

# 或使用 LoadBalancer（云环境）
kubectl patch svc user-api-svc -n go-zero -p '{"spec": {"type": "LoadBalancer"}}'

# 或使用 NodePort
kubectl patch svc user-api-svc -n go-zero -p '{"spec": {"type": "NodePort"}}'
```

## 配置说明

### 服务架构

```
┌─────────────────────────────────────────────────────────────────┐
│                        Kubernetes Cluster                        │
│  ┌─────────────────────────────────────────────────────────┐    │
│  │                   Namespace: go-zero                     │    │
│  │                                                          │    │
│  │  ┌──────────────┐                                       │    │
│  │  │   Ingress    │ ◄── 外部流量 (api.go-zero.local)     │    │
│  │  └──────┬───────┘                                       │    │
│  │         │                                                │    │
│  │         ▼                                                │    │
│  │  ┌──────────────┐                                       │    │
│  │  │  user-api    │ (:8888)                               │    │
│  │  │  (2 replicas)│                                       │    │
│  │  └──────┬───────┘                                       │    │
│  │         │ gRPC                                          │    │
│  │  ┌──────┴──────┬──────────────┐                        │    │
│  │  │             │              │                         │    │
│  │  ▼             ▼              ▼                         │    │
│  │ ┌────────┐  ┌────────┐  ┌────────┐                     │    │
│  │ │user-rpc│  │auth-rpc│  │role-rpc│                     │    │
│  │ │(:9001) │  │(:9002) │  │(:9003) │                     │    │
│  │ └───┬────┘  └───┬────┘  └───┬────┘                     │    │
│  │     │           │           │                           │    │
│  │     ▼           ▼           ▼                           │    │
│  │  ┌──────────────────────────────────────┐              │    │
│  │  │            MySQL (:3306)              │              │    │
│  │  │            Redis (:6379)              │              │    │
│  │  └──────────────────────────────────────┘              │    │
│  └─────────────────────────────────────────────────────────┘    │
└─────────────────────────────────────────────────────────────────┘
```

### 资源配置

| 服务 | 副本数 | CPU 请求/限制 | 内存 请求/限制 |
|------|--------|---------------|----------------|
| user-api | 2 | 100m/200m | 128Mi/256Mi |
| user-rpc | 2 | 100m/200m | 128Mi/256Mi |
| auth-rpc | 2 | 100m/200m | 128Mi/256Mi |
| role-rpc | 2 | 100m/200m | 128Mi/256Mi |
| MySQL | 1 | 250m/500m | 512Mi/1Gi |
| Redis | 1 | 100m/200m | 128Mi/256Mi |

### Secret 配置项

| 配置项 | 说明 | 默认值 |
|--------|------|--------|
| MYSQL_HOST | MySQL 服务地址 | mysql-svc |
| MYSQL_PORT | MySQL 端口 | 3306 |
| MYSQL_USER | MySQL 用户名 | root |
| MYSQL_PASSWORD | MySQL 密码 | 123456 |
| MYSQL_DATABASE | 数据库名 | go_zero_user |
| REDIS_HOST | Redis 服务地址 | redis-svc |
| REDIS_PORT | Redis 端口 | 6379 |
| JWT_ACCESS_SECRET | JWT 访问密钥 | go-zero-user-access-secret-key |
| JWT_REFRESH_SECRET | JWT 刷新密钥 | go-zero-user-refresh-secret-key |

## 生产环境部署

### 使用 Kustomize Overlay

```bash
# 1. 修改生产环境配置
vim deploy/k8s/overlays/production/secret.yaml

# 2. 修改镜像地址
vim deploy/k8s/overlays/production/kustomization.yaml

# 3. 部署到生产环境
kubectl apply -k deploy/k8s/overlays/production/
```

### 生产环境检查清单

- [ ] 修改 Secret 中的所有默认密码
- [ ] 配置正确的镜像仓库地址
- [ ] 调整副本数和资源限制
- [ ] 配置 TLS/SSL 证书
- [ ] 配置持久化存储（PV/PVC）
- [ ] 配置监控和告警（Prometheus/Grafana）
- [ ] 配置日志收集（ELK/Loki）
- [ ] 配置网络策略（NetworkPolicy）
- [ ] 配置 Pod 反亲和性
- [ ] 配置 HPA 自动扩缩容

### 配置 HPA（水平自动扩缩容）

```yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: user-api-hpa
  namespace: go-zero
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: user-api
  minReplicas: 2
  maxReplicas: 10
  metrics:
    - type: Resource
      resource:
        name: cpu
        target:
          type: Utilization
          averageUtilization: 70
```

### 配置 Ingress TLS

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: user-api-ingress-tls
  namespace: go-zero
  annotations:
    cert-manager.io/cluster-issuer: "letsencrypt-prod"
spec:
  ingressClassName: nginx
  tls:
    - hosts:
        - api.your-domain.com
      secretName: go-zero-api-tls
  rules:
    - host: api.your-domain.com
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: user-api-svc
                port:
                  number: 8888
```

## 常用命令

### 部署管理

```bash
# 查看所有资源
kubectl get all -n go-zero

# 查看 Pod 详情
kubectl describe pod <pod-name> -n go-zero

# 查看 Pod 日志
kubectl logs -f <pod-name> -n go-zero

# 进入 Pod
kubectl exec -it <pod-name> -n go-zero -- /bin/sh

# 重启 Deployment
kubectl rollout restart deployment <deployment-name> -n go-zero

# 回滚 Deployment
kubectl rollout undo deployment <deployment-name> -n go-zero
```

### 使用 Makefile

```bash
make k8s-deploy          # 部署到 K8s
make k8s-delete          # 删除部署
make k8s-status          # 查看状态
make k8s-logs-user-api   # 查看 user-api 日志
make k8s-restart         # 重启所有微服务
make k8s-port-forward    # 端口转发到本地
```

## 故障排查

### Pod 无法启动

```bash
# 查看 Pod 状态
kubectl get pods -n go-zero

# 查看 Pod 事件
kubectl describe pod <pod-name> -n go-zero

# 查看 Pod 日志
kubectl logs <pod-name> -n go-zero --previous
```

### 常见问题

1. **ImagePullBackOff**: 镜像拉取失败
   - 检查镜像名称和标签是否正确
   - 检查镜像仓库认证

2. **CrashLoopBackOff**: Pod 启动后崩溃
   - 查看日志排查错误原因
   - 检查配置文件是否正确

3. **Pending**: Pod 无法调度
   - 检查节点资源是否充足
   - 检查 PVC 是否绑定成功

4. **连接数据库失败**
   - 确认 MySQL Pod 已就绪
   - 检查 Secret 配置是否正确
   - 检查服务名称是否正确

### 调试模式

```bash
# 启动调试 Pod
kubectl run debug --rm -it --image=alpine -n go-zero -- /bin/sh

# 在调试 Pod 中测试连接
apk add mysql-client redis
mysql -h mysql-svc -u root -p
redis-cli -h redis-svc ping
```

## 清理资源

```bash
# 删除所有资源
make k8s-delete

# 或手动删除
kubectl delete namespace go-zero
```

---

如有问题，请查看项目 README.md 或提交 Issue。

