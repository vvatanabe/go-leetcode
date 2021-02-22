package main

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums) == 0 || r*c != len(nums)*len(nums[0]) {
		return nums
	}
	out := make([][]int, r)
	for i := range out {
		out[i] = make([]int, c)
	}
	var rows, cols int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[0]); j++ {
			out[rows][cols] = nums[i][j]
			cols++
			if cols == c {
				rows++
				cols = 0
			}
		}
	}
	return out
}
