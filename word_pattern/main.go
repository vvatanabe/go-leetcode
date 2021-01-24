package main

import "strings"

func wordPattern(pattern string, s string) bool {
	arr := strings.Split(s, " ")
	if len(arr) != len(pattern) {
		return false
	}
	r2s := make(map[rune]string)
	s2r := make(map[string]rune)
	for i, r := range pattern {
		s, ok := r2s[r]
		if ok && s != arr[i] {
			return false
		}
		r2s[r] = arr[i]
		s2r[arr[i]] = r
	}
	return len(r2s) == len(s2r)
}
