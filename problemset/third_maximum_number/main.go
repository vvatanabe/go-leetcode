package main

func thirdMax(nums []int) int {
	min := ^(1 << 32) + 1
	rank := []int{min, min, min}
	for _, num := range nums {
		switch {
		case rank[0] < num:
			rank[2], rank[1], rank[0] = rank[1], rank[0], num
		case rank[1] < num && num < rank[0]:
			rank[2], rank[1] = rank[1], num
		case rank[2] < num && num < rank[1]:
			rank[2] = num
		}
	}
	if rank[2] == min {
		return rank[0]
	}
	return rank[2]
}
