package main

import (
	"fmt"
	"math/rand"
)

func quickSort(arr []int) {
	quickSortRec1(arr, 0, len(arr)-1)
}

func quickSortRec1(arr []int, left, right int) {
	if left < right {
		parition1 := parition2(arr, left, right)
		quickSortRec1(arr, left, parition1-1)
		quickSortRec1(arr, parition1, right)
	}
}

func parition2(arr []int, left, right int) int {
	// 产生 left 到 right 之间的随机下标
	randIdx := rand.Intn(right-left+1) + left
	// 交换 pivot 到最右侧
	arr[randIdx], arr[right] = arr[right], arr[randIdx]
	pivot := arr[right]
	i := left - 1
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// 把 pivot 放到正确的位置
	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}

func mergeSortTest(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSortTest(arr[:mid])
	right := mergeSortTest(arr[mid:])
	return merge(left, right)
}

func MergeTest(left, right []int) []int {
	result := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	if i < len(left) {
		result = append(result, left[i:]...)
	}
	if j < len(right) {
		result = append(result, right[j:]...)
	}
	return result
}

func HeapSortTest(arr []int) {
	n := len(arr)

	// 成堆
	for i := n/2 - 1; i >= 0; i-- {
		HeapifyTest(arr, n, i)
	}

	// 堆排序
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		HeapifyTest(arr, i, 0)
	}
}

func HeapifyTest(arr []int, n int, i int) {
	largest := i
	left := i*2 + 1
	right := i*2 + 2

	if left < n && arr[largest] < arr[left] {
		largest = left
	}
	if right < n && arr[largest] < arr[right] {
		largest = right
	}

	if largest != i {
		arr[largest], arr[i] = arr[i], arr[largest]
		HeapifyTest(arr, n, largest)
	}

}

func BubbleSortTest(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swap := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
}

func InsertSortTest(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		temp := arr[i]
		j := i - 1
		for ; j > 0; j-- {
			if arr[j] > temp {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = temp
	}
}

func test() {
	arr := []int{1, 5, 4, 2, 6, 8, 2}
	InsertSortTest(arr)
	fmt.Println(arr)
}
