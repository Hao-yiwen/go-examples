package main

import "fmt"

func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func main5() {
	arr := []int{1, 5, 6, 9, 3, 2}
	BubbleSort(arr)
	fmt.Println(arr)
}
