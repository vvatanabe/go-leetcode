package main

func moveZeroes(nums []int) {
	var lastNoZeroIndex int
	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		if cur != 0 {
			last := nums[lastNoZeroIndex]
			nums[i] = last
			nums[lastNoZeroIndex] = cur
			lastNoZeroIndex++
		}
	}
}
