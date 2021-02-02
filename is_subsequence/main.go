package main

func isSubsequence(s string, t string) bool {
	var i, j int
	sl, tl := len(s), len(t)
	for i < sl && j < tl {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == sl
}
