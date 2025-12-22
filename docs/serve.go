package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))
	fmt.Println("📚 文档服务器启动成功！")
	fmt.Println("🌐 访问地址: http://localhost:8080/doc/")
	fmt.Println("按 Ctrl+C 停止服务器")
	http.ListenAndServe(":8080", nil)
}
