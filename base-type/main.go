package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}
	return x / y, nil
}

type Person struct {
	Name string
	Age  int
}

func (p Person) SayHello() {
	fmt.Println("Hello, my name is", p.Name)
}

type Speaker interface {
	Speak()
}

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Dog Speak")
}

func ReadFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return scanner.Err()
}

func main() {
	// var a int = 10
	// var b float64 = 0.4
	// var c string = "您好"
	// var d bool = true
	// var e rune = 'a'

	// fmt.Println(a, b, c, d, e)

	// var f []rune = []rune("您好")
	// for i, v := range f {
	// 	fmt.Printf("%d: %c\n", i, v)
	// }

	// for i, v := range c {
	// 	fmt.Printf("%d: %c\n", i, v)
	// }

	// for i := 0; i < len(c); i++ {
	// 	fmt.Printf("%c\n", c[i])
	// }

	// for i := 0; i < len(f); i++ {
	// 	fmt.Printf("%c\n", f[i])
	// }

	// type User struct {
	// 	Name string
	// 	Age  int
	// }

	// u := User{
	// 	Name: "John",
	// 	Age:  20,
	// }

	// fmt.Printf("%v\n", u)
	// fmt.Printf("%+v\n", u)
	// fmt.Printf("%#v\n", u)
	// fmt.Printf("%10v\n", u)
	// fmt.Printf("%T\n", u)
	// fmt.Println(fmt.Sprintf("%v", u))

	// n := 65
	// fmt.Printf("%c\n", n)
	// fmt.Printf("%d\n", n)
	// fmt.Printf("%b\n", n)

	// f := 123.467
	// fmt.Printf("%f\n", f)
	// fmt.Printf("%g\n", f)
	// fmt.Printf("%v\n", f)
	// fmt.Printf("%.2f\n", f)
	// fmt.Printf("%.2g\n", f)
	// fmt.Printf("%e\n", f)

	// s := "Hi\tGood"
	// fmt.Printf("%s\n", s)
	// fmt.Printf("%v\n", s)
	// fmt.Printf("%q\n", s)

	// x := 10
	// p := &x
	// fmt.Printf("%p\n", p)
	// fmt.Printf("%v\n", p)
	// fmt.Printf("%v\n", *p)

	// var f1 float64

	// fmt.Printf("%f\n", f1)
	// fmt.Printf("%v\n", &f1)

	// f1 = 0.0
	// fmt.Printf("%f\n", f1)
	// fmt.Printf("%v\n", &f1)

	// var p1 *float64
	// fmt.Printf("%p\n", p1)

	// tmp := 64.0
	// p1 = &tmp
	// fmt.Printf("%p\n", p1)
	// fmt.Printf("%f\n", *p1)

	// type test struct {
	// 	Name string
	// 	age  int
	// }

	// t := test{
	// 	Name: "John",
	// 	age:  20,
	// }

	// fmt.Printf("%v\n", t)
	// fmt.Printf("%v\n", t.Name)
	// fmt.Printf("%v\n", t.age)

	// var nums []int
	// lastCap := 0

	// // 添加 2000 个元素
	// for i := 0; i < 2000; i++ {
	// 	nums = append(nums, i)

	// 	// 如果容量发生了变化，打印出来
	// 	if cap(nums) != lastCap {
	// 		// 防止除以零
	// 		growth := 0.0
	// 		if lastCap != 0 {
	// 			growth = float64(cap(nums)) / float64(lastCap)
	// 		}
	// 		fmt.Printf("长度: %-4d -> 新容量: %-4d (增长比例: %.2f倍)\n",
	// 			len(nums), cap(nums), growth)
	// 		lastCap = cap(nums)
	// 	}
	// }

	fmt.Printf("Hello %s\n", "World")

	var a int = 10
	var b = 10

	var (
		name string = "John"
		age  int    = 20
	)

	fmt.Println(a, b, name, age)

	const Pi = 3.1415926
	const (
		StatusOk string = "200"
		NotFound string = "404"
	)

	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	var h bool = true
	var c string = "Hello"
	var d int = 10
	var e float64 = 0
	var f rune = 'a'
	var g byte = 'b'
	var i string = "a"
	fmt.Println(h, c, d, e, f)
	fmt.Println(strconv.FormatBool(h))
	// 这里输出的是98，而不是'b'，因为strconv.Itoa 是把整数转成字符串，而 g 是 byte 类型（本质是uint8），它的数值是'b'的ASCII值98。
	fmt.Println(string(g)) // 这样才会输出 "b"
	fmt.Println(string(f)) // 这样才会输出 "a"
	fmt.Println(i)
	var j []rune = []rune("hello")
	fmt.Println(string(j))

	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	silce1 := []int{1, 2, 3, 4, 5}
	fmt.Println(silce1)

	slice2 := make([]int, 5, 10)
	fmt.Println(slice2, len(slice2), cap(slice2))

	slice2 = append(slice2, 6, 7, 8, 9, 10)
	fmt.Println(slice2, len(slice2), cap(slice2))

	subslice := slice2[7:]
	fmt.Println(subslice)
	fmt.Println(slice2)

	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	fmt.Println(m)
	fmt.Println(m["a"])
	fmt.Println(m["b"])
	fmt.Println(m["c"])
	fmt.Println(m["d"])

	if value, ok := m["d"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("d not found")
	}

	result, err := div(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	changeValue := func(val *int) {
		*val = 100
	}

	k := 10
	changeValue(&k)
	fmt.Println(k)

	p := Person{
		Name: "John",
		Age:  20,
	}
	fmt.Println(p)
	p.SayHello()

	var s Speaker = Dog{}
	s.Speak()

	err1 := ReadFile("test.txt")
	if err1 != nil {
		fmt.Println(err1, 123123)
	}

	msgChan := make(chan string)
	go worker(msgChan)
	for {
		msg := <-msgChan
		fmt.Println(msg)
		if msg == "" {
			break
		}
		fmt.Println(msg)
	}
}

func worker(c chan string) {
	for i := 0; i < 10; i++ {
		c <- "hello"
	}
	close(c)
	// c <- "dsa"
}
