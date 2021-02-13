package main

func reverseString(s []byte) {
	var start, end = 0, len(s) - 1
	for start <= end {
		tmp := s[end]
		s[end] = s[start]
		s[start] = tmp
		start++
		end--
	}
}

func reverseString2(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}
