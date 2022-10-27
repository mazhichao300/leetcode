package main

import (
	"fmt"
)

func makeGrid(n int) [][]byte {
	ans := make([][]byte, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			ans[i][j] = '.'
		}
	}
	return ans
}
func solveNQueens(n int) [][]string {
	grid := makeGrid(n)
	ans := make([][]string, 0)

	flag := make(map[int]int, 0)
	dfs(grid, 0, n, n, flag, &ans)

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isValid(x, y int, flag map[int]int) bool {
	for k, v := range flag {
		if v == y {
			return false
		}
		x1 := abs(x - k)
		y1 := abs(y - v)
		if x1 == y1 {
			return false
		}
	}
	return true
}

func dfs(grid [][]byte, x, n, left int, flag map[int]int, ans *[][]string) {
	if left == 0 {
		tmp := make([]string, n)
		for i := 0; i < n; i++ {
			tmp[i] = string(grid[i])
		}
		*ans = append(*ans, tmp)
		return
	}
	if x < 0 || x >= n {
		return
	}

	for i := 0; i < n; i++ {
		if !isValid(x, i, flag) {
			continue
		}
		grid[x][i] = 'Q'
		flag[x] = i
		dfs(grid, x+1, n, left-1, flag, ans)
		delete(flag, x)
		grid[x][i] = '.'
	}
}

func main() {
	ans := solveNQueens(4)
	fmt.Println(ans)
}
