package main

func longestPalindrome(s string) int {
	mp := make(map[rune]int)
	for _, r := range s {
		mp[r]++
	}
	var cnt int
	for _, v := range mp {
		cnt += v - v%2
	}
	if cnt < len(s) {
		cnt++
	}
	return cnt
}
