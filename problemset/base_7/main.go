package main

import (
	"strconv"
)

func convertToBase7(num int) string {
	var prefix string
	if num < 0 {
		prefix = "-"
		num *= -1
	}
	var out string
	for num > 0 {
		i := num % 7
		out = strconv.Itoa(i) + out
		num /= 7
	}
	if out == "" {
		return "0"
	}
	return prefix + out
}
