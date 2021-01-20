package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	table := [26]int{}
	for i := 0; i < len(s); i++ {
		table[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		table[t[i]-'a']--
		if table[t[i]-'a'] < 0 {
			return false
		}
	}
	return true
}

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[int32]int)
	for _, r := range s {
		m[r]++
	}
	for _, r := range t {
		if cnt, ok := m[r]; !ok || cnt == 0 {
			return false
		}
		m[r]--
	}
	return true
}
