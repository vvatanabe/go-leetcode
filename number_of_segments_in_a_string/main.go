package main

func countSegments(s string) int {
	var cnt int
	var pre rune
	for i, r := range s {
		if (i == 0 || pre == ' ') && r != ' ' {
			cnt++
		}
		pre = r
	}
	return cnt
}

func countSegments2(s string) int {
	var cnt, seg int
	for _, r := range s {
		switch {
		case r != ' ':
			seg++
		case seg > 0:
			cnt++
			seg = 0
		}
	}
	if seg > 0 {
		cnt++
	}
	return cnt
}
