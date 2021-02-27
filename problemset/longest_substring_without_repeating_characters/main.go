package main

func lengthOfLongestSubstring(s string) int {
	mp := make(map[rune]struct{})
	var runes []rune
	var length int
	for _, r := range s {
		_, ok := mp[r]
		if ok {
			for i := len(runes) - 1; i >= 0; i-- {
				if runes[i] == r {
					for j := 0; j < i; j++ {
						delete(mp, runes[j])
					}
					runes = runes[i+1:]
					runes = append(runes, r)
					break
				}
			}
		} else {
			runes = append(runes, r)
			mp[r] = struct{}{}
			length = max(length, len(runes))
		}
	}
	return length
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
