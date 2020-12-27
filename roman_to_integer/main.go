package main

func romanToInt(s string) int {
	dic := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	size := len(s)
	var out int
	for i := 0; i < size-1; i++ {
		curNum := dic[s[i:i+1]]
		if curNum < dic[s[i+1:i+2]] {
			out -= curNum
		} else {
			out += curNum
		}
	}
	out += dic[s[size-1:]]
	return out
}
