package main

func numIslands(grid [][]byte) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != '1' {
				continue
			}
			ans++
			dfs(grid, i, j)
		}
	}
	return ans
}

func dfs(grid [][]byte, i, j int) {
	m := len(grid)
	n := len(grid[0])
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1' {
		return
	}
	grid[i][j] = '2'
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}
