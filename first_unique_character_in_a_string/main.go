package main

func firstUniqChar(s string) int {
	var m [26]int
	for _, r := range s {
		m[r-'a']++
	}
	for i, r := range s {
		if m[r-'a'] <= 1 {
			return i
		}
	}
	return -1
}

func firstUniqChar2(s string) int {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	for i, r := range s {
		if m[r] <= 1 {
			return i
		}
	}
	return -1
}
