package main

func canConstruct(ransomNote string, magazine string) bool {
	var m [26]int
	for _, r := range magazine {
		m[r-'a']++
	}
	for _, r := range ransomNote {
		m[r-'a']--
		if m[r-'a'] < 0 {
			return false
		}
	}
	return true
}

func canConstruct2(ransomNote string, magazine string) bool {
	m := make(map[rune]int)
	for _, r := range magazine {
		m[r]++
	}
	for _, r := range ransomNote {
		m[r]--
		if m[r] < 0 {
			return false
		}
	}
	return true
}
