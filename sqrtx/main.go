package main

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	low := 0
	high := x / 2
	for low <= high {
		mid := (low + high) / 2
		if x < mid*mid {
			high = mid - 1
		} else if x > mid*mid {
			low = mid + 1
		} else {
			return mid
		}
	}
	return high
}
