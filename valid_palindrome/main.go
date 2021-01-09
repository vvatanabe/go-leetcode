package main

import "strings"

// Input: "A man, a plan, a canal: Panama"
// Output: true
func isPalindrome(s string) bool {
	if s == "" {
		return true
	}

	min, max := 0, len(s)-1

	for min < max {
		if !isAlphanumeric(s[min]) {
			min++
			continue
		}
		if !isAlphanumeric(s[max]) {
			max--
			continue
		}

		if !strings.EqualFold(string(s[min]), string(s[max])) {
			return false
		}
		min++
		max--
	}
	return true
}

func isAlphanumeric(a byte) bool {
	return (a >= 'a' && a <= 'z') ||
		(a >= 'A' && a <= 'Z') ||
		(a >= '0' && a <= '9')
}
