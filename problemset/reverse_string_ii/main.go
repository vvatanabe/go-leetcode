package main

func reverseStr(s string, k int) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i += 2 * k {
		l, r := i, i+k-1
		if len(runes)-1 < r {
			r = len(runes) - 1
		}
		for l < r {
			runes[l], runes[r] = runes[r], runes[l]
			l++
			r--
		}
	}
	return string(runes)
}
