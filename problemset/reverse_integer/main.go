package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// pop operation:
//    pop = x % 10;
//    x /= 10;
//
// push operation:
//    temp = rev * 10 + pop;
//    rev = temp;
func reverse(x int) int {
	var out int64
	for x != 0 {
		pop := int64(x % 10)
		x /= 10
		out = (out * 10) + pop
	}
	if out < -math.MaxInt32 || math.MaxInt32 < out {
		return 0
	}
	return int(out)
}

func reverse2(x int) int {
	isNegative := x < 0
	str := fmt.Sprint(x)
	var reversed string
	if isNegative {
		str = strings.TrimPrefix(str, "-")
		reversed += "-"
	}
	for i := len(str) - 1; i >= 0; i-- {
		start := i
		end := i + 1
		reversed += str[start:end]
	}
	rx, _ := strconv.ParseInt(reversed, 10, 64)
	if rx < -math.MaxInt32 || math.MaxInt32 < rx {
		return 0
	}
	return int(rx)
}
