package main

import "fmt"

func initFlag() [][]int {
	arr := make([][]int, 10)
	for i, _ := range arr {
		arr[i] = make([]int, 10)
	}
	return arr
}

func solveSudoku(board [][]byte) {
	row := initFlag()
	col := initFlag()
	block := initFlag()

	for x, b := range board {
		for y, i := range b {
			if i != '.' {
				pos := int(i - '0')
				row[x][pos] = 1
				col[y][pos] = 1
				block[x/3*3+y/3][pos] = 1
			}
		}
	}
	dfs(board, row, col, block, 0, 0)
}

func dfs(board [][]byte, row [][]int, col [][]int, block [][]int, x, y int) bool {
	nextX := x
	nextY := y
	if x >= 9 {
		return true
	}
	if y < 8 {
		nextY++
	} else {
		nextX++
		nextY = 0
	}

	if board[x][y] != '.' {
		return dfs(board, row, col, block, nextX, nextY)
	} else {
		for i := 1; i < 10; i++ {
			c := byte(i + '0')
			if row[x][i] == 0 && col[y][i] == 0 && block[x/3*3+y/3][i] == 0 {
				board[x][y] = c
				row[x][i] = 1
				col[y][i] = 1
				block[x/3*3+y/3][i] = 1
				if dfs(board, row, col, block, nextX, nextY) == true {
					return true
				}

				board[x][y] = '.'
				row[x][i] = 0
				col[y][i] = 0
				block[x/3*3+y/3][i] = 0
			}
		}
	}
	return false
}

func main() {
	board := make([][]byte, 9)
	input := []string{
		"...3...4.",
		"....1.2..",
		".8...6..7",
		"...4...3.",
		"6....7..5",
		"....2.1..",
		"8.7..5...",
		"92.......",
		".56.....9",
	}
	for i := 0; i < 9; i++ {
		board[i] = make([]byte, 9)
		for j := 0; j < 9; j++ {
			board[i][j] = input[i][j]
		}
	}

	solveSudoku(board)
	// fmt.Println(board)
	for i := 0; i < len(board); i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(int(board[i][j]-'0'), " ")
		}
		fmt.Println("")
	}
}
