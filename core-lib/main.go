package main

import (
	"fmt"
	"strconv"
	"strings"
)

func test() {
	name := "John"
	age := 10

	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)

	msg := fmt.Sprintf("User[%s] is active", name)
	fmt.Println(msg)

	fmt.Print("Please enter text:")
	// reader := bufio.NewReader(os.Stdin)
	// input, _ := reader.ReadString('\n')
	// fmt.Println("You entered:", strings.Trim(input, " "))

	data := "apple,banana,cherry"

	parts := strings.Split(data, ",")
	fmt.Println(parts)

	numStr := "126"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Number:", num)
	}

	str := strconv.Itoa(42)
	fmt.Println(str)
}
