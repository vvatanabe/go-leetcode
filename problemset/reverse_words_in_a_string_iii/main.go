package main

import (
	"strings"
)

// Input: "Let's take LeetCode contest"
// Output: "s'teL ekat edoCteeL tsetnoc"
func reverseWords(s string) string {
	if len(s) == 0 {
		return ""
	}
	words, start, end := []rune(s), 0, -1
	for i, r := range words {
		if r == ' ' {
			for start < end {
				words[start], words[end] = words[end], words[start]
				start++
				end--
			}
			start, end = i+1, i
		} else {
			end++
		}
	}
	for start < end {
		words[start], words[end] = words[end], words[start]
		start++
		end--
	}
	return string(words)
}

func reverseWords2(s string) string {
	if len(s) == 0 {
		return ""
	}
	var words strings.Builder
	var word string
	for _, r := range s {
		if r != ' ' {
			word = string(r) + word
		} else {
			words.WriteString(word)
			words.WriteString(" ")
			word = ""
		}
	}
	words.WriteString(word)
	return words.String()
}
