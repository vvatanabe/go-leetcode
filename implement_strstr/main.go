package main

import (
	"fmt"
)

func main() {
	haystack := "mississippi"
	needle := "issip"
	out := strStr(haystack, needle)
	fmt.Println("out", out)
}

func strStr(haystack string, needle string) int {
	if haystack == "" && needle == "" {
		return 0
	}
	if needle == "" {
		return 0
	}
	for i := 0; i < len(haystack); i++ {
		if i+len(needle) > len(haystack) {
			return -1
		}
		m := i
		for j := 0; j < len(needle); j++ {
			if haystack[m:m+1] == needle[j:j+1] {
				if j == len(needle)-1 {
					return i
				}
				m++
			} else {
				break
			}
		}
	}
	return -1
}
