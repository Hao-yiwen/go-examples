package main

import "fmt"

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

	type User struct {
		Name string
		Age  int
	}

	u := User{
		Name: "John",
		Age:  20,
	}

	fmt.Printf("%v\n", u)
	fmt.Printf("%+v\n", u)
	fmt.Printf("%#v\n", u)
	fmt.Printf("%10v\n", u)
	fmt.Printf("%T\n", u)
	fmt.Println(fmt.Sprintf("%v", u))

	n := 65
	fmt.Printf("%c\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%b\n", n)

	f := 123.467
	fmt.Printf("%f\n", f)
	fmt.Printf("%g\n", f)
	fmt.Printf("%v\n", f)
	fmt.Printf("%.2f\n", f)
	fmt.Printf("%.2g\n", f)
	fmt.Printf("%e\n", f)

	s := "Hi\tGood"
	fmt.Printf("%s\n", s)
	fmt.Printf("%v\n", s)
	fmt.Printf("%q\n", s)

	x := 10
	p := &x
	fmt.Printf("%p\n", p)
	fmt.Printf("%v\n", p)
	fmt.Printf("%v\n", *p)

	var f1 float64

	fmt.Printf("%f\n", f1)
	fmt.Printf("%v\n", &f1)

	f1 = 0.0
	fmt.Printf("%f\n", f1)
	fmt.Printf("%v\n", &f1)

	var p1 *float64
	fmt.Printf("%p\n", p1)

	tmp := 64.0
	p1 = &tmp
	fmt.Printf("%p\n", p1)
	fmt.Printf("%f\n", *p1)

	type test struct {
		Name string
		age  int
	}

	t := test{
		Name: "John",
		age:  20,
	}

	fmt.Printf("%v\n", t)
	fmt.Printf("%v\n", t.Name)
	fmt.Printf("%v\n", t.age)

	var nums []int
	lastCap := 0

	// 添加 2000 个元素
	for i := 0; i < 2000; i++ {
		nums = append(nums, i)

		// 如果容量发生了变化，打印出来
		if cap(nums) != lastCap {
			// 防止除以零
			growth := 0.0
			if lastCap != 0 {
				growth = float64(cap(nums)) / float64(lastCap)
			}
			fmt.Printf("长度: %-4d -> 新容量: %-4d (增长比例: %.2f倍)\n",
				len(nums), cap(nums), growth)
			lastCap = cap(nums)
		}
	}
}
