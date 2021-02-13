package main

func isPerfectSquare(num int) bool {
	low, high := 1, num
	for low <= high {
		mid := (low + high) / 2
		switch {
		case mid*mid == num:
			return true
		case mid*mid < num:
			low = mid + 1
		case mid*mid > num:
			high = mid - 1
		}
	}
	return false
}
