.PHONY: build run clean

# 编译 Go 代码为 WebAssembly
build:
	GOOS=js GOARCH=wasm go build -o main.wasm main.go

# 启动开发服务器
run: build
	go run server.go

# 清理生成的文件
clean:
	rm -f main.wasm
