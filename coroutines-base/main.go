package main

import (
	"fmt"
	"time"
)

// 员工：专门负责从 jobs 管道里拿任务做
func worker(id int, jobs <-chan string) {
	// range 会一直阻塞等待，直到管道被关闭
	for job := range jobs {
		fmt.Printf("👷 员工 %d 接到了任务: %s\n", id, job)
		time.Sleep(500 * time.Millisecond) // 假装在忙
	}
	fmt.Println("👋 员工", id, "下班了")
}

func main() {
	// // 创建一个传送带，能放 3 个汉堡（缓冲区）
	// burgers := make(chan string, 3)

	// // 厨师 A：做汉堡 (生产者)
	// go func() {
	// 	for i := 0; i < 5; i++ {
	// 		randInt := rand.Intn(100)
	// 		var burger string
	// 		switch randInt % 4 {
	// 		case 0:
	// 			burger = "香辣鸡腿堡"
	// 		case 1:
	// 			burger = "原味鸡腿堡"
	// 		case 2:
	// 			burger = "麦辣鸡腿堡"
	// 		case 3:
	// 			burger = "劲脆鸡腿堡"
	// 		default:
	// 			burger = "蜜汁鸡腿堡"
	// 		}
	// 		fmt.Println("厨师 A 做好了一个汉堡", burger)
	// 		time.Sleep(time.Duration(100 * time.Millisecond))
	// 		burgers <- burger

	// 	}
	// 	close(burgers) // 做完了，关闭传送带
	// }()

	// // 厨师 B：打包汉堡 (消费者)
	// // 只要传送带上有东西，我就拿。如果空了，我就自动等。
	// for burger := range burgers {
	// 	fmt.Println("厨师 B 打包了：", burger)
	// 	time.Sleep(time.Duration(100 * time.Millisecond))
	// }
	jobs := make(chan string, 5)

	// 1. 先启动 2 个员工（协程），让它们等着接活
	// 注意：这时候 jobs 还是空的，它们会阻塞在 range 那里
	go worker(1, jobs)
	go worker(2, jobs)

	// 2. 主线程（老板）开始发号施令
	fmt.Println("👨‍💼 老板开始派活了...")

	jobs <- "打扫卫生" // 主线程发消息
	jobs <- "写代码"  // 主线程发消息
	jobs <- "修电脑"  // 主线程发消息
	jobs <- "订外卖"  // 主线程发消息

	fmt.Println("👨‍💼 活派完了，老板准备关门")

	// 3. 关键步骤：关闭通道
	// 告诉员工：“没活了，干完手头的就下班吧”
	close(jobs)

	// 为了演示效果，主线程稍微等一下，不然主线程退出了协程还没打印完
	time.Sleep(3 * time.Second)
}
