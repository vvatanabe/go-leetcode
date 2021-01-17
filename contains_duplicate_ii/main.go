package main

func containsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < i+1+k; j++ {
			if j == len(nums) {
				break
			}
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}
