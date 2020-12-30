package main

func climbStairs(n int) int {
	return f(n, make(map[int]int))
}

func f(n int, memo map[int]int) int {
	if n <= 2 {
		return n
	}
	if v, ok := memo[n]; ok {
		return v
	}
	memo[n] = f(n-2, memo) + f(n-1, memo)
	return memo[n]
}

// Time Limit Exceeded
func climbStairs2(n int) int {
	if n <= 2 {
		return n
	}
	return climbStairs(n-2) + climbStairs(n-1)
}
