package main

func isPowerOfThree(n int) bool {
	for n%3 == 0 && n >= 3 {
		n /= 3
	}
	return n == 1
}

func isPowerOfThree2(n int) bool {
	return n > 0 && 1162261467%n == 0
}
