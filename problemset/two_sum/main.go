package main

func twoSum(nums []int, target int) []int {
	max := len(nums)
	if max < 2 {
		return []int{}
	}
	memo := make(map[int]int, max)
	for i, k := range nums {
		if v, ok := memo[k]; ok {
			return []int{v, i}
		}
		memo[target-k] = i
	}
	return []int{}
}
