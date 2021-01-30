package main

func findTheDifference(s string, t string) byte {
	var m [26]int
	for _, r := range s {
		m[r-'a']++
	}
	for _, r := range t {
		m[r-'a']--
		if m[r-'a'] < 0 {
			return byte(r)
		}
	}
	return byte(0)
}
