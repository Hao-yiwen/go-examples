package main

import (
	"fmt"
	"math/rand"
)

func QuickSortTest3(arr []int) {
	QuickSortRecTest3(arr, 0, len(arr)-1)
}

func QuickSortRecTest3(arr []int, left int, right int) {
	if left < right {
		pivot := PartitionTest3(arr, left, right)
		QuickSortRecTest3(arr, left, pivot-1)
		QuickSortRecTest3(arr, pivot+1, right)
	}
}

func PartitionTest3(arr []int, start int, end int) int {
	rd := rand.Intn(end-start+1) + start
	arr[rd], arr[end] = arr[end], arr[rd]
	pivot := arr[end]
	j := start - 1
	for i := start; i < end; i++ {
		if arr[i] < pivot {
			j++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	arr[j+1], arr[end] = arr[end], arr[j+1]
	return j + 1
}

func main() {
	arr := []int{1, 4, 3, 2, 5, 0}
	QuickSortTest3(arr)
	fmt.Println(arr)
}
