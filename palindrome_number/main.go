package main

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	var rev int
	for x > rev {
		pop := x % 10
		rev = (rev * 10) + pop
		x /= 10
	}
	return x == rev || x == rev/10
}

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	y := x
	var rev int
	for y != 0 {
		pop := y % 10
		y /= 10
		rev = (rev * 10) + pop
	}
	return x == rev
}
