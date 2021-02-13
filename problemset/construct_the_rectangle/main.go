package main

import "math"

func constructRectangle(area int) []int {
	w := int(math.Sqrt(float64(area)))
	for area%w != 0 {
		w--
	}
	return []int{area / w, w}
}

func constructRectangle2(area int) []int {
	w := int(sqrt(float64(area)))
	for area%w != 0 {
		w--
	}
	return []int{area / w, w}
}

func sqrt(x float64) float64 {
	l, m, r := 0.0, 0.0, x
	for i := 0; i < 1000; i++ {
		m = (l + r) / 2
		if m*m-x < 0 {
			l = m
		} else {
			r = m
		}
	}
	return r
}
