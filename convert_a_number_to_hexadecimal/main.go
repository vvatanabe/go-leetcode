package main

func toHex(num int) string {
	n := uint32(num)
	table := "0123456789abcdef"
	var hex string
	for n > 0 {
		v := n % 16
		hex = table[v:v+1] + hex
		n /= 16
	}
	if hex == "" {
		hex = "0"
	}
	return hex
}
