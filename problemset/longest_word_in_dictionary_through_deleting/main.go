package main

import (
	"strings"
)

// Input: s = "abpcplea", dictionary = ["ale","apple","monkey","plea"]
// Output: "apple"

func findLongestWord(s string, dictionary []string) string {
	var match string
	for _, word := range dictionary {
		if isSubstr(s, word) {
			if len(match) < len(word) || (len(word) == len(match) && word < match) {
				match = word
			}
		}
	}
	return match

}

func isSubstr(s, word string) bool {
	for _, c := range word {
		p := strings.Index(s, string(c))
		if p == -1 {
			return false
		}
		s = s[p+1:]
	}
	return true
}
