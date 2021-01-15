package main

func trailingZeroes(n int) int {
	var v int
	for n > 0 {
		n /= 5
		v += n
	}
	return v
}

func trailingZeroes2(n int) int {
	var v int
	for {
		n = n / 5
		if n == 0 {
			break
		}
		v += n
	}
	return v
}
