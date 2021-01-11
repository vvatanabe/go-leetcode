package main

func twoSum(numbers []int, target int) []int {
	max := len(numbers)
	if max < 2 {
		return []int{}
	}
	memo := make(map[int]int, max)
	for i, k := range numbers {
		if v, ok := memo[k]; ok {
			return []int{v + 1, i + 1}
		}
		memo[target-k] = i
	}
	return []int{}
}
