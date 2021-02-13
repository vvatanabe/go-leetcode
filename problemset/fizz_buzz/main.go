package main

import "strconv"

func fizzBuzz(n int) []string {
	if n == 0 {
		return nil
	}
	var s string
	if n%3 == 0 {
		s += "Fizz"
	}
	if n%5 == 0 {
		s += "Buzz"
	}
	if s == "" {
		s = strconv.Itoa(n)
	}
	return append(fizzBuzz(n-1), s)
}

func fizzBuzz2(n int) []string {
	var arr []string
	for i := 1; i <= n; i++ {
		var s string
		if i%3 == 0 {
			s += "Fizz"
		}
		if i%5 == 0 {
			s += "Buzz"
		}
		if s == "" {
			s = strconv.Itoa(i)
		}
		arr = append(arr, s)
	}
	return arr
}
