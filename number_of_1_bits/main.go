package main

func hammingWeight(num uint32) int {
	var c uint32
	for num != 0 {
		c += num & 1
		num >>= 1
	}
	return int(c)
}
