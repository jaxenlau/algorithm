package qsort

import (
	"fmt"
)

// QuickSort1 快排实现 1
func QuickSort1(arr []int) {
	implQuickSort1(arr, 0, len(arr)-1)
}

func implQuickSort1(arr []int, start, end int) {

	if start > end {
		return
	}

	left := start
	right := end
	key := arr[start]

	for left != right {
		for left < right && arr[right] >= key {
			right--
		}
		arr[left] = arr[right]

		for left < right && arr[left] <= key {
			left++
		}
		arr[right] = arr[left]
	}

	fmt.Printf("left: %d, right: %d, k: %d, %v ---- >> ", left, right, key, arr)

	arr[left] = key

	fmt.Printf("%v \n", arr)

	implQuickSort1(arr, start, left-1)
	implQuickSort1(arr, left+1, end)
}

// QuickSort2 快排实现 2
func QuickSort2(arr []int) {
	implQuickSort2(arr, 0, len(arr)-1)
}

func implQuickSort2(arr []int, start, end int) {

	if start > end {
		return
	}

	left := start
	right := end
	key := arr[start]

	for left != right {

		for left < right && arr[right] >= key {
			right--
		}

		for left < right && arr[left] <= key {
			left++
		}

		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}

	fmt.Printf("left: %d, right: %d, k: %d, %v ---- >> ", left, right, key, arr)

	arr[start] = arr[left]
	arr[left] = key

	fmt.Printf("%v \n", arr)

	implQuickSort2(arr, start, left-1)
	implQuickSort2(arr, left+1, end)
}

// QuickSort3 快排实现 3， 选择最后一个为枢轴
func QuickSort3(arr []int) {
	separateSort(arr, 0, len(arr)-1)
}

func separateSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	i := partition(arr, start, end)
	fmt.Printf("i=%d, arr: %v\n", i, arr)
	// separateSort(arr, start, i-1)
	// separateSort(arr, i+1, end)

}

func partition(arr []int, start, end int) int {
	pivot := arr[end]

	i := start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]

	return i
}
