package main

// 0) 1
// 1) 1,1
// 2) 1,2,1
// 3) 1,3,3,1
// 4) 1,4,6,4,1
// 5) 1,5,10,10,5,1

func generate(numRows int) [][]int {
	var rows [][]int
	for i := 0; i < numRows; i++ {
		rows = append(rows, []int{1})
		if i == 0 {
			continue
		}
		for j := 1; j < i; j++ {
			rows[i] = append(rows[i], rows[i-1][j-1]+rows[i-1][j])
		}
		rows[i] = append(rows[i], 1)
	}
	return rows
}
