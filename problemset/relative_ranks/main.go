package main

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	var temp [][]int
	for i, v := range score {
		temp = append(temp, []int{i, v})
	}
	sort.Slice(temp, func(i, j int) bool {
		return temp[i][1] > temp[j][1]
	})
	out := make([]string, len(score))
	for i, v := range temp {
		switch i {
		case 0:
			out[v[0]] = "Gold Medal"
		case 1:
			out[v[0]] = "Silver Medal"
		case 2:
			out[v[0]] = "Bronze Medal"
		default:
			out[v[0]] = strconv.Itoa(i + 1)
		}
	}
	return out
}
