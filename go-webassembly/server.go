//go:build ignore

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := "8080"

	// 设置正确的 MIME 类型
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/main.wasm" {
			w.Header().Set("Content-Type", "application/wasm")
		}
		http.FileServer(http.Dir(".")).ServeHTTP(w, r)
	})

	fmt.Printf("服务器启动在 http://localhost:%s\n", port)
	fmt.Println("按 Ctrl+C 停止服务器")

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
