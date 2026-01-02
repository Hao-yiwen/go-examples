package main

import (
	"fmt"
	"io"
	"strings"
)

func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1
	}
	return 0
}

func testThridRes() (int, int, int, bool) {
	return 1, 2, 3, true
}

func part1() {
	i := 1
	i++
	fmt.Println(i)

	// 定义运行所需的变量
	src := []byte{1, 5, 2, 6, 8}
	var size int
	const sizeOne = 4
	const sizeTwo = 10
	validateOnly := false
	var err error
	var errShortInput = fmt.Errorf("input too short")
	const shift = 1
	update := func(v byte) {
		fmt.Printf("Updated value: %d\n", v)
	}

Loop:
	for n := 0; n < len(src); n += size {
		switch {
		case src[n] < sizeOne:
			if validateOnly {
				size = 1
				break
			}
			size = 1
			update(src[n])

		case src[n] < sizeTwo:
			if n+1 >= len(src) {
				err = errShortInput
				break Loop
			}
			if validateOnly {
				size = 2
				break
			}
			size = 2
			update(src[n] + src[n+1]<<shift)

		default:
			size = 1
		}
	}

	if err != nil {
		fmt.Println("Loop error:", err)
	}

	test1, test2 := []byte{1, 2, 4}, []byte{1, 2, 4}
	fmt.Println(Compare(test1, test2))

	var t interface{}
	t = Compare(test1, test2)
	switch t.(type) {
	default:
		fmt.Println("default")
	case bool:
		fmt.Println("bool")
	case int:
		fmt.Println("int")
	}

	testThridRes()
	fmt.Printf("=============\n")

	c := "123dsa23123dfsads12"
	b := []byte(c)
	for i := 0; i < len(b); {
		val, next := nextInt(b, i)
		i = next
		fmt.Println(val)
	}

	ReadFull(strings.NewReader(c), b)
	fmt.Println(string(b))
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func nextInt(b []byte, i int) (x, next int) {
	for ; i < len(b) && !isDigit(b[i]); i++ {
	}
	x = 0
	for ; i < len(b) && isDigit(b[i]); i++ {
		x = x*10 + int(b[i]) - '0'
	}
	next = i
	return
}

// ReadFull 函数会持续从给定的 io.Reader（r）中读取数据，
// 直到 buf 这个字节切片被完全填满或者出现读取错误。
// 具体做法是：每次调用 r.Read(buf)，读取尽可能多的数据到 buf 里，
// 并将已读的数据数量累加到 n 上，然后缩小 buf，使其指向未被填充的部分。
// 读到出错（err 非 nil）或者 buf 没空间时停止，最终返回总共读取的字节数和可能的错误。

func ReadFull(r io.Reader, buf []byte) (n int, err error) {
	for len(buf) > 0 && err == nil {
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
	}
	return
}
