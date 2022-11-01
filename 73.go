package main

func setZeroes(matrix [][]int) {
	row := len(matrix)
	col := len(matrix[0])
	zeroRow := make([]int, row)
	zeroCol := make([]int, col)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {
				zeroRow[i] = 1
				zeroCol[j] = 1
			}
		}
	}
	for i, v := range zeroRow {
		if v == 0 {
			continue
		}
		for j := 0; j < col; j++ {
			matrix[i][j] = 0
		}
	}
	for j, v := range zeroCol {
		if v == 0 {
			continue
		}
		for i := 0; i < row; i++ {
			matrix[i][j] = 0
		}
	}
}
