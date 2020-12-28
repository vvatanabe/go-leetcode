package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	out := removeElement(nums, val)
	fmt.Println("out", out, nums)
}

func removeElement(nums []int, val int) int {
	length := 0
	for j := 0; j < len(nums); j++ {
		if val != nums[j] {
			nums[length] = nums[j]
			length++
		}
	}
	return length
}
