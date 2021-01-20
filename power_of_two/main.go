package main

// Bit Operation
// When n is a power of 2, the binary number of n has only one 1 at the top.
// 2  -> 10
// 4  -> 100
// 8  -> 1000
// 16 -> 10000
func isPowerOfTwo(n int) bool {
	return n != 0 && n&(n-1) == 0
}

// Divide by Two
func isPowerOfTwo2(n int) bool {
	for n%2 == 0 && n >= 2 {
		n /= 2
	}
	return n == 1
}
