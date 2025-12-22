# Go 中文文档

本地 Go 官方文档的中文版本，包含教程、API 参考和各种指南。

## 启动文档服务器

### 方法一：使用 Python（推荐）

```bash
# Python 3
cd docs
python3 -m http.server 8080

# 访问 http://localhost:8080/doc/
```

### 方法二：使用 Go

```bash
# 在 docs 目录创建简单的 HTTP 服务器
cd docs
go run serve.go
```

创建 `serve.go` 文件（在项目根目录）：

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.Handle("/", http.FileServer(http.Dir("./docs")))
    fmt.Println("文档服务器启动: http://localhost:8080/doc/")
    http.ListenAndServe(":8080", nil)
}
```

### 方法三：使用 npx（需要 Node.js）

```bash
cd docs
npx serve -p 8080

# 访问 http://localhost:8080/doc/
```

## 访问文档

启动服务器后，在浏览器中访问：
- 主页：http://localhost:8080/doc/
- 数据库教程：http://localhost:8080/doc/database/
- Go 教程：http://localhost:8080/doc/tutorial/

## 文档内容

- 📖 教程和入门指南
- 🗄️ 数据库操作指南
- 🔧 工具和命令参考
- 📚 各版本发布说明
