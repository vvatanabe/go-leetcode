package main

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	length := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[length] {
			length++
			nums[length] = nums[j]
		}
	}
	return length + 1
}
