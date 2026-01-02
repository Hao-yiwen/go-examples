package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
	}
	return string(result), nil
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func main2() {
	content, err := Contents("test.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(content)
	}

	// for i := 0; i < 5; i++ {
	// 	defer fmt.Printf("%d\n", i)
	// }

	// b()

	// type SyncedBuffer struct {
	// 	lock   sync.Mutex
	// 	buffer buffer.Buffer
	// }

	// p := new(SyncedBuffer)
	// var v SyncedBuffer

	// p.lock.Lock()
	// p.buffer.AppendString("Hello, ")
	// p.lock.Unlock()
	// v.lock.Lock()
	// v.buffer.AppendString("World!")
	// v.lock.Unlock()

	// fmt.Printf("p: %v\n", p)
	// fmt.Printf("v: %v\n", v)

	// var p *[]int = new([]int)
	// var v []int = make([]int, 10)
	// fmt.Printf("p: %v\n", p)
	// fmt.Printf("v: %v\n", v)
	// var p *[]int = new([]int)
	// *p = make([]int, 10)
	// (*p)[0] = 1

	// fmt.Printf("*p: %v\n", *p)

	// v := make([]int, 10)
	// fmt.Printf("v: %v\n", v)

	// // --- 演示 Slice 和 Map 的引用特性 ---
	// fmt.Println("\n=== Slice & Map Reference Test ===")

	// // 1. Slice
	// originalSlice := []int{1, 2, 3}
	// modifySlice(originalSlice)
	// fmt.Printf("Original Slice after modify: %v (受影响)\n", originalSlice)

	// // 2. Map
	// originalMap := map[string]int{"a": 1}
	// modifyMap(originalMap)
	// fmt.Printf("Original Map after modify: %v (受影响)\n", originalMap)

	list := []int{1, 2, 3}
	addOne(list)
	fmt.Println(list) // 输出还是 [1 2 3]，如果是指针传递，应该能看到 100

	list1 := []int{1, 2, 3}
	modifySlice(list1)
	fmt.Println(list1)

	// Allocate the top-level slice, the same as before.
	// picture := make([][]uint8, 10) // One row per unit of y.
	// // Allocate one large slice to hold all the pixels.
	// pixels := make([]uint8, 10*10) // Has type []uint8 even though picture is [][]uint8.
	// // Loop over the rows, slicing each row from the front of the remaining pixels slice.
	// for i := range picture {
	// 	picture[i] = make([]uint8, 10)
	// }
	// picture[0][0] = 1
	// picture[0][1] = 2
	// picture[0][2] = 3
	// picture[0][3] = 4
	// picture[0][4] = 5
	// picture[0][5] = 6
	// picture[0][6] = 7
	// picture[0][7] = 8
	// picture[0][8] = 9
	// picture[0][9] = 10
	// fmt.Printf("picture: %v\n", picture)
	// fmt.Printf("pixels: %v\n", pixels)

	// var timeZone = map[string]int{
	// 	"UTC": 0 * 60 * 60,
	// 	"EST": -5 * 60 * 60,
	// 	"CST": -6 * 60 * 60,
	// 	"MST": -7 * 60 * 60,
	// 	"PST": -8 * 60 * 60,
	// }
	// fmt.Println(timeZone)

	// if value, ok := timeZone["UTC"]; ok {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("UTC not found")
	// }

	// delete(timeZone, "UTC")
	// fmt.Println(timeZone)

	// fmt.Printf("Hello %d\n", 23)
	// fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
	// fmt.Println("Hello", 23)
	// fmt.Println(fmt.Sprint("Hello ", 23))

	// t := &T{7, -2.35, "abc\tdef"}
	// fmt.Printf("%v\n", t)
	// fmt.Printf("%+v\n", t)
	// fmt.Printf("%#v\n", t)
	// fmt.Printf("%#v\n", timeZone)
	// fmt.Printf("%T\n", timeZone)
	// fmt.Printf("%v\n", t)

	// a := 1.3
	// str := fmt.Sprintf("%f", a)
	// fmt.Println("float as string:", str)

	// b := 1
	// fmt.Printf("v= %v, s= %s\n", b, strconv.Itoa(b))

	// c := "121312"
	// num, err := strconv.Atoi(c)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("Number:", num)
	// }
	fmt.Println(gopath)
}

type ByteSize float64

const (
	KB ByteSize = 1 << (10 * (iota + 1)) // 1 << (10*1) = 1024
	MB                                   // 1 << (10*2) = 1048576
	GB                                   // 1 << (10*3) = 1073741824
	TB                                   // 1 << (10*4) = 1099511627776
	PB                                   // 1 << (10*5)
	EB                                   // 1 << (10*6)
	ZB                                   // 1 << (10*7)
	YB                                   // 1 << (10*8)
)

func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}

var user string
var home string
var gopath string

func init2() {
	// if user == "" {
	// 	log.Fatal("$USER not set")
	// }
	// if home == "" {
	// 	home = "/home/" + user
	// }
	// if gopath == "" {
	// 	gopath = home + "/go"
	// }
	// 这行代码的作用是允许用户通过命令行参数 --gopath 来覆盖默认的 gopath 路径。
	// flag.StringVar 函数将命令行参数 "gopath" 绑定到 gopath 变量。
	// 如果运行程序时指定了 --gopath 参数，gopath 变量的值就会被更新为传入的值；
	// 否则 gopath 保持其原有的默认值。
	gopath = "嘿嘿"
	flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
	flag.Parse()
	fmt.Println(gopath)
}

type T struct {
	a int
	b float64
	c string
}

func (t *T) String() string {
	return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
}

func addOne(s []int) {
	s = append(s, 100) // s 的长度变成了 4，但这是函数内部那个 s 副本的长度变化
}

func modifySlice(s []int) {
	s[0] = 999       // 修改底层数组，外部可见
	s = append(s, 4) // 发生扩容或修改了局部长度，外部不可见
	fmt.Println("Inside modifySlice:", s)
}
