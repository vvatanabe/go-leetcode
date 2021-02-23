package main

import "fmt"

func main() {
	nums := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	fmt.Println("nums:", nums)
	insertionSort(nums)
	fmt.Println("nums:", nums)
}

func insertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		j := i
		for j > 0 {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
			j = j - 1
		}
	}
}
