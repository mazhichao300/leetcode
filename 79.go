package main

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, x, y int, word string, idx int) bool {
	row := len(board)
	col := len(board[0])
	lenOfWord := len(word)
	if idx >= lenOfWord {
		return true
	}
	if x < 0 || x >= row || y < 0 || y >= col {
		return false
	}
	if board[x][y] == '1' {
		return false
	}
	current := word[idx]
	if board[x][y] != current {
		return false
	}
	tmp := board[x][y]
	board[x][y] = '1'
	ans := dfs(board, x-1, y, word, idx+1)
	if !ans {
		ans = dfs(board, x+1, y, word, idx+1)
	}
	if !ans {
		ans = dfs(board, x, y-1, word, idx+1)
	}
	if !ans {
		ans = dfs(board, x, y+1, word, idx+1)
	}
	board[x][y] = tmp
	return ans
}
