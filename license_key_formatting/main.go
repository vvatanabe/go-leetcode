package main

func licenseKeyFormatting(S string, K int) string {
	var runes []rune
	for _, r := range S {
		if r == '-' {
			continue
		}
		if 'a' <= r && r <= 'z' {
			r -= 32
		}
		runes = append(runes, r)
	}
	mod := len(runes) % K
	var (
		out  []rune
		size int
	)
	for i, r := range runes {
		out = append(out, r)
		size++
		if i+1 != len(runes) && (i+1 == mod || size == K) {
			out = append(out, '-')
			size = 0
		}
	}
	return string(out)
}
