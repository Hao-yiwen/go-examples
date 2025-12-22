# GO 示例项目

Go 语言学习示例代码集合。

## 项目结构

- [data-access/](data-access/) - 数据库访问示例
- [web-service-gin/](web-service-gin/) - 基于 Gin 框架的 REST API 示例
- [workspace/](workspace/) - 其他 Go 编程示例（HTTP 服务器、模板、RAG 服务器、类型系统等）
- [docs/](docs/) - Go 官方中文文档（本地版）

## 快速开始

进入任意示例目录，运行 `go run main.go` 即可启动。部分示例需要配置 `.env` 环境变量文件。

### 启动文档服务器

```bash
# 使用 Go 启动文档
cd docs && go run ./serve.go

# 或使用 Python
cd docs && python3 -m http.server 8080
```

访问 http://localhost:8080/doc/ 查看文档。