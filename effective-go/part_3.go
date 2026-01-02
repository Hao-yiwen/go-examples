package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type ByteSlice []byte

func (p *ByteSlice) Append(data []byte) []byte {
	slice := *p
	slice = append(slice, data...)
	*p = slice
	return slice
}

func (p *ByteSlice) Write(data []byte) (n int, err error) {
	slice := *p
	slice = append(slice, data...)
	*p = slice
	return len(data), nil
}

func main3() {
	var b ByteSlice
	fmt.Fprintf(&b, "This hour has %d days\n", 7)
	fmt.Println(string(b))

	// a := 65
	b1 := 1.1
	c := true
	// fmt.Printf(string(a))
	// a := 1
	// b1 := strconv.Itoa(a)
	// fmt.Println(b1)
	// fmt.Printf("%d\n", []uint8(b1))
	str := fmt.Sprintf("%f", b1)
	fmt.Println(str)
	str12 := fmt.Sprintf("%v", c)
	fmt.Println(str12)
	str3 := strconv.FormatFloat(b1, 'f', -1, 64)
	fmt.Println(str3)
	str4 := strconv.FormatBool(c)
	fmt.Println(str4)

	// http.Handle("/args", HandlerFunc(ArgServer))
	// http.ListenAndServe(":8080", nil)
}

func ArgServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, os.Args, 21321)
}

type HandlerFunc func(http.ResponseWriter, *http.Request)

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}
