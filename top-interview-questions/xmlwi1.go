package main

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	return searchMatrixBinary(matrix, target, 0, m-1, 0, n-1)
}

func searchMatrixBinary(matrix [][]int, target int, row1, row2, col1, col2 int) bool {
	if row1 > row2 || col1 > col2 {
		return false
	}
	if row1 == row2 && col1 == col2 {
		return matrix[row1][col1] == target
	}
	rowmid := (row1 + row2) / 2
	colmid := (col1 + col2) / 2
	if matrix[rowmid][colmid] == target {
		return true
	}
	if matrix[rowmid][colmid] > target {
		if searchMatrixBinary(matrix, target, row1, rowmid-1, col1, col2) {
			return true
		}
		if searchMatrixBinary(matrix, target, rowmid, row2, col1, colmid-1) {
			return true
		}
		return false
	}
	if searchMatrixBinary(matrix, target, row1, row2, colmid+1, col2) {
		return true
	}
	if searchMatrixBinary(matrix, target, rowmid+1, row2, col1, colmid) {
		return true
	}
	return false
}

func searchMatrix1(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	row := m - 1
	col := 0

	for row >= 0 && col < n {
		if matrix[row][col] == target {
			return true
		}

		if matrix[row][col] > target {
			row--
		} else {
			col++
		}
	}

	return false
}
