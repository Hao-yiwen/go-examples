package main

import "fmt"

func InsertSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main6() {
	arr := []int{1, 4, 5, 2, 3}
	InsertSort(arr)
	fmt.Println(arr)
}
