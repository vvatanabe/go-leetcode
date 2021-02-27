package main

import "fmt"

func main() {
	nums := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	fmt.Println("nums:", nums)
	mergeSort(nums)
	fmt.Println("nums:", nums)
}

func mergeSort(arr []int) {
	copy(arr, sort(arr))
}

func sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	l := sort(arr[:mid])
	r := sort(arr[mid:])
	return merge(l, r)
}

func merge(left, right []int) []int {
	var arr []int
	for len(left) > 0 || len(right) > 0 {
		switch {
		case len(left) == 0:
			return append(arr, right...)
		case len(right) == 0:
			return append(arr, left...)
		case left[0] <= right[0]:
			arr = append(arr, left[0])
			left = left[1:]
		default:
			arr = append(arr, right[0])
			right = right[1:]
		}
	}
	return arr
}
