package main

func maximumProduct(nums []int) int {
	max1, max2, max3 := -1001, -1001, -1001
	min1, min2 := 1001, 1001
	for _, num := range nums {
		switch {
		case max1 < num:
			max1, max2, max3 = num, max1, max2
		case max2 < num:
			max2, max3 = num, max2
		case max3 < num:
			max3 = num
		}
		switch {
		case min1 > num:
			min1, min2 = num, min1
		case min2 > num:
			min2 = num
		}
	}
	return max(max1*max2*max3, max1*min1*min2)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
