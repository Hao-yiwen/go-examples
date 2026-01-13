package main

import (
	"fmt"
	"math/rand"
)

func quickSortTest2(arr []int) {
	quickSortRecTest2(arr, 0, len(arr)-1)
}

func quickSortRecTest2(arr []int, left, right int) {
	if left < right {
		pivot := PartitionTest2(arr, left, right)
		quickSortRecTest2(arr, left, pivot-1)
		quickSortRecTest2(arr, pivot+1, right)
	}
}

func PartitionTest2(arr []int, start, end int) int {
	rdnum := rand.Intn(end-start+1) + start
	arr[rdnum], arr[end] = arr[end], arr[rdnum]
	pivot := arr[end]
	j := start - 1
	for i := start; i < end; i++ {
		if arr[i] < pivot {
			j++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[j+1], arr[end] = arr[end], arr[j+1]
	return j + 1
}

func test1() {
	arr := []int{5, 4, 3, 2, 1}
	quickSortTest2(arr)
	fmt.Println(arr)
}
