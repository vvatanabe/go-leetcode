package main

func findWords(words []string) []string {
	mp := make(map[rune]int)
	for i, row := range []string{"qwertyuiopQWERTYUIOP", "asdfghjklASDFGHJKL", "zxcvbnmZXCVBNM"} {
		for _, r := range row {
			mp[r] = i + 1
		}
	}
	var arr []string
	for _, word := range words {
		row := mp[rune(word[0])]
		for _, r := range word {
			if mp[r] != row {
				goto SKIP
			}
		}
		arr = append(arr, word)
	SKIP:
	}
	return arr
}
