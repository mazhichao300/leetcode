package main

import (
	"fmt"
)

func totalNQueens(n int) int {
	ans := 0

	flag := make(map[int]int, 0)
	dfs(0, n, n, flag, &ans)

	return ans
}

func isValid(x, y int, flag map[int]int) bool {
	for k, v := range flag {
		if v == y {
			return false
		}
		x1 := x - k
		y1 := y - v
		if x1 == y1 || x1 == -y1 {
			return false
		}
	}
	return true
}

func dfs(x, n, left int, flag map[int]int, ans *int) {
	if left == 0 {
		*ans = *ans + 1
		return
	}
	if x < 0 || x >= n {
		return
	}

	for i := 0; i < n; i++ {
		if !isValid(x, i, flag) {
			continue
		}
		flag[x] = i
		dfs(x+1, n, left-1, flag, ans)
		delete(flag, x)
	}
}

func main() {
	ans := totalNQueens(4)
	fmt.Println(ans)
}
