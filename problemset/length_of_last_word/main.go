package main

import "fmt"

func main() {
	s := "Hello World"
	out := lengthOfLastWord(s)
	fmt.Println("out", out)
}

func lengthOfLastWord(s string) int {
	var cnt int
	var ps bool
	for i := 0; i < len(s); i++ {
		if s[i:i+1] != " " {
			if ps {
				cnt = 0
			}
			cnt++
			ps = false
		} else {
			ps = true
		}
	}
	return cnt
}
