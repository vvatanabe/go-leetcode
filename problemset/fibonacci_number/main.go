package main

func fib(n int) int {
	return helper(n, make(map[int]int))
}

func helper(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}
	if n, ok := memo[n]; ok {
		return n
	}
	l := helper(n-1, memo)
	r := helper(n-2, memo)
	memo[n-1] = l
	memo[n-2] = r
	memo[n] = l + r
	return l + r
}
