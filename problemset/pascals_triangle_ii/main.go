package main

func getRow(rowIndex int) []int {
	var rows [][]int
	for i := 0; i < rowIndex+1; i++ {
		rows = append(rows, []int{1})
		if i == 0 {
			continue
		}
		for j := 1; j < i; j++ {
			rows[i] = append(rows[i], rows[i-1][j-1]+rows[i-1][j])
		}
		rows[i] = append(rows[i], 1)
	}
	return rows[rowIndex]
}
