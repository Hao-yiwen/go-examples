package main

import "fmt"

func QuickSort(arr []int) {
	quickSortRec(arr, 0, len(arr)-1)
}

func quickSortRec(arr []int, low, high int) {
	if low < high {
		privotIndex := partition(arr, low, high)
		quickSortRec(arr, low, privotIndex-1)
		quickSortRec(arr, privotIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high] // 选择最后一个元素作为基准
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main1() {
	nums := []int{4, 5, 1, 2, 6, 9}
	QuickSort(nums)
	fmt.Println(nums)
}
