package main

func convertToTitle(n int) string {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var title string
	for n > 0 {
		d := n % len(alphabet)
		n /= len(alphabet)
		if d == 0 { // division remainder is zero
			d = len(alphabet)
			n--
		}
		title = alphabet[d-1:d] + title
	}
	return title
}
