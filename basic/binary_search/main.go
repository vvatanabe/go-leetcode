package main

import "fmt"

func main() {
	nums := []int{1, 3, 8, 18, 31, 43, 68, 72, 101}
	k := 8
	fmt.Println("nums:", nums)
	fmt.Println("k:", k)
	fmt.Println("index:", binarySearch(nums, k))
}

func binarySearch(nums []int, k int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) / 2
		switch {
		case nums[mid] < k:
			low = mid + 1
		case nums[mid] > k:
			high = mid - 1
		default:
			return mid
		}
	}
	return -1
}
