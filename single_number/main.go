package main

//Input: nums = [4,1,2,1,2]
//Output: 4

func singleNumber(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	for i := 0; i < len(nums); i++ {
		_, ok := m[nums[i]]
		if ok {
			delete(m, nums[i])
			continue
		}
		m[nums[i]] = struct{}{}
	}
	for i := range m {
		return i
	}
	return 0
}
