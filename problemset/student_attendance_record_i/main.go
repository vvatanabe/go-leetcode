package main

//'A' : Absent.
//'L' : Late.
//'P' : Present.

func checkRecord(s string) bool {
	var a, l int
	for _, r := range s {
		switch r {
		case 'A':
			a++
		case 'L':
			l++
		case 'P':
			l = 0
		}
		if a > 1 || l > 2 {
			return false
		}
	}
	return true
}
