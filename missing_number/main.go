package main

func missingNumber(nums []int) int {
	var max int
	var hasZero bool
	a := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		a[nums[i]] = struct{}{}
		if max < nums[i] {
			max = nums[i]
		}
		if nums[i] == 0 {
			hasZero = true
		}
	}
	for i := 0; i < len(nums); i++ {
		if max == nums[i] {
			continue
		}
		if _, ok := a[nums[i]+1]; !ok {
			return nums[i] + 1
		}
	}
	if !hasZero {
		return 0
	}
	return max + 1
}
