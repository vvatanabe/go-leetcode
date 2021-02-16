package main

// lower 97 - 122
// upper 65 - 90
func detectCapitalUse(word string) bool {
	runes := []rune(word)
	allLow := runes[0] >= 97 && runes[0] <= 122
	if allLow && len(runes) == 2 {
		return runes[1] >= 97 && runes[1] <= 122
	}
	for i := 1; i < len(runes)-1; i++ {
		if allLow && (!(runes[i] >= 97 && runes[i] <= 122) || !(runes[i+1] >= 97 && runes[i+1] <= 122)) {
			return false
		}
		if !allLow && (runes[i] >= 97 && runes[i] <= 122) != (runes[i+1] >= 97 && runes[i+1] <= 122) {
			return false
		}
	}
	return true
}
