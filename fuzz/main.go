package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	fmt.Printf("input: %q\n", s)
	b := []rune(s)
	fmt.Printf("runes: %q\n", b)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b), nil
}

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, err := Reverse(input)
	if err != nil {
		fmt.Printf("Cannot reverse: %v\n", err)
		return
	}
	doubleRev, err := Reverse(rev)
	if err != nil {
		fmt.Printf("Cannot reverse: %v\n", err)
		return
	}
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)
}
