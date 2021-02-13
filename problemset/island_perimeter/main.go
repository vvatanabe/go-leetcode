package main

func islandPerimeter(grid [][]int) int {
	var cnt int
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				cnt += 4
				if i > 0 && grid[i-1][j] == 1 {
					cnt -= 2
				}
				if j > 0 && grid[i][j-1] == 1 {
					cnt -= 2
				}
			}
		}
	}
	return cnt
}

func islandPerimeter2(grid [][]int) int {
	return helper(grid, 0, 0)
}

func helper(grid [][]int, i, j int) int {
	if i == len(grid) {
		return 0
	}
	var cnt int
	if grid[i][j] == 1 {
		cnt += 4
		if i > 0 && grid[i-1][j] == 1 {
			cnt -= 2
		}
		if j > 0 && grid[i][j-1] == 1 {
			cnt -= 2
		}
	}
	if j == len(grid[i])-1 {
		i++
		j = -1
	}
	j++
	return cnt + helper(grid, i, j)
}
