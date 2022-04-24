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
	board[0] = []byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'}
	board[1] = []byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'}
	board[2] = []byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'}
	board[3] = []byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'}
	board[4] = []byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'}
	board[5] = []byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'}
	board[6] = []byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'}
	board[7] = []byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'}
	board[8] = []byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'}
	fmt.Println(board)
	solveSudoku(board)
	fmt.Println(board)
}
