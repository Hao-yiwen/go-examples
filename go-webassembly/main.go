package main

import (
	"fmt"
	"syscall/js"
)

// 计数器
var counter int

// 更新显示
func updateDisplay() {
	document := js.Global().Get("document")
	element := document.Call("getElementById", "counter")
	element.Set("innerText", fmt.Sprintf("%d", counter))
}

// 增加计数
func increment(this js.Value, args []js.Value) interface{} {
	counter++
	updateDisplay()
	return nil
}

// 减少计数
func decrement(this js.Value, args []js.Value) interface{} {
	counter--
	updateDisplay()
	return nil
}

// 重置计数
func reset(this js.Value, args []js.Value) interface{} {
	counter = 0
	updateDisplay()
	return nil
}

// 显示问候语
func greet(this js.Value, args []js.Value) interface{} {
	document := js.Global().Get("document")
	input := document.Call("getElementById", "nameInput")
	name := input.Get("value").String()

	if name == "" {
		name = "世界"
	}

	message := fmt.Sprintf("你好, %s! 欢迎使用 Go WebAssembly!", name)

	output := document.Call("getElementById", "greeting")
	output.Set("innerText", message)

	return nil
}

func main() {
	fmt.Println("Go WebAssembly 已加载!")

	// 注册 JavaScript 函数
	js.Global().Set("goIncrement", js.FuncOf(increment))
	js.Global().Set("goDecrement", js.FuncOf(decrement))
	js.Global().Set("goReset", js.FuncOf(reset))
	js.Global().Set("goGreet", js.FuncOf(greet))

	// 保持程序运行
	select {}
}
