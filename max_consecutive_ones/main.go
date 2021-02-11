package main

func findMaxConsecutiveOnes(nums []int) int {
	var max, cur int
	for _, n := range nums {
		cur += n
		cur *= n
		if max < cur {
			max = cur
		}
	}
	return max
}

func findMaxConsecutiveOnes2(nums []int) int {
	var max, cur int
	for _, n := range nums {
		if n == 1 {
			cur++
			continue
		}
		if max < cur {
			max = cur
		}
		cur = 0
	}
	if max < cur {
		max = cur
	}
	return max
}
