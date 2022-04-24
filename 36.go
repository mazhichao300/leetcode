package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	col := [9][9]bool{}
	row := [9][9]bool{}
	grid := [9][9]bool{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			c := board[i][j]
			if c == '.' {
				continue
			}

			id := int(c - '0')

			if row[i][id] == true {
				return false
			}

			row[i][id] = true

			if col[j][id] == true {
				return false
			}
			col[j][id] = true

			idx := i/3*3 + j/3

			if grid[idx][id] == true {
				return false
			}
			grid[idx][id] = true
		}
	}

	return true
}

func main() {
	ans := isValidSudoku([][]byte{})
	fmt.Println(ans)
}
