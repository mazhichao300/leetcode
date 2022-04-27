package main

import "fmt"

var flag map[string]int // 0 pacific; 1 atlantic; 2 all;3 no where

func pacificAtlantic(heights [][]int) [][]int {
	ans := make([][]int, 0)
	flag = make(map[string]int, 0)
	m := len(heights)
	n := len(heights[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(heights, i, j) == 2 {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}

func dfs(heights [][]int, x, y int) int {
	m := len(heights)
	n := len(heights[0])
	key := fmt.Sprintf("%d_%d", x, y)
	if c, ok := flag[key]; ok && c != -1 {
		return flag[key]
	}
	if x == 0 && y == n-1 {
		flag[key] = 2
		return 2
	}
	if y == 0 && x == m-1 {
		flag[key] = 2
		return 2
	}

	ans := make([]int, 4)
	// up
	if x-1 >= 0 && heights[x-1][y] <= heights[x][y] {
		flag[key] = 3
		t := dfs(heights, x-1, y)
		flag[key] = -1
		ans[t] = 1
	} else if x == 0 {
		ans[0] = 1
	}
	// down
	if x+1 < m && heights[x+1][y] <= heights[x][y] {
		flag[key] = 3
		t := dfs(heights, x+1, y)
		flag[key] = -1
		ans[t] = 1
	} else if x+1 == m {
		ans[1] = 1
	}
	// left
	if y-1 >= 0 && heights[x][y-1] <= heights[x][y] {
		flag[key] = 3
		t := dfs(heights, x, y-1)
		flag[key] = -1
		ans[t] = 1
	} else if y == 0 {
		ans[0] = 1
	}
	// right
	if y+1 < n && heights[x][y+1] <= heights[x][y] {
		flag[key] = 3
		t := dfs(heights, x, y+1)
		flag[key] = -1
		ans[t] = 1
	} else if y+1 == n {
		ans[1] = 1
	}

	if ans[2] == 1 || (ans[0] == 1 && ans[1] == 1) {
		flag[key] = 2
		return 2
	}
	if ans[0] == 1 {
		// flag[key] = 0
		return 0
	}
	if ans[1] == 1 {
		// flag[key] = 1
		return 1
	}
	// flag[key] = 3
	return 3
}

func main() {
	input := [][]int{
		[]int{3, 3, 3, 3, 3, 3},
		[]int{3, 0, 3, 3, 0, 3},
		[]int{3, 3, 3, 3, 3, 3},
	}

	ans := pacificAtlantic(input)
	fmt.Println(ans)
}
