package main

import (
	"fmt"
	"strconv"
)

func findComplement(num int) int {
	b := fmt.Sprintf("%b", num)
	var out string
	for _, v := range b {
		if v == '0' {
			out += "1"
		} else {
			out += "0"
		}
	}
	n, _ := strconv.ParseUint(out, 2, 32)
	return int(n)
}
