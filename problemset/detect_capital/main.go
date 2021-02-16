package main

func detectCapitalUse(word string) bool {
	runes := []rune(word)
	if len(runes) == 2 {
		return !isLower(runes[0]) || isLower(runes[1])
	}
	allLower := isLower(runes[0])
	for i := 1; i < len(runes)-1; i++ {
		if allLower {
			if !isLower(runes[i]) || !isLower(runes[i+1]) {
				return false
			}
		} else {
			if isLower(runes[i]) != isLower(runes[i+1]) {
				return false
			}
		}
	}
	return true
}

func isLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}
