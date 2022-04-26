package main

import "fmt"

func projectionArea(grid [][]int) int {
	length := len(grid)
	ans := 0

	for i := 0; i < length; i++ {
		maxRow := 0
		maxCol := 0
		for j := 0; j < length; j++ {
			if grid[i][j] > 0 { // xy面的投影面积
				ans += 1
			}
			if grid[i][j] > maxRow {
				maxRow = grid[i][j] // 每一行的最大值
			}
			if grid[j][i] > maxCol {
				maxCol = grid[j][i] // 每一列的最大值
			}
		}
		ans += maxRow + maxCol
	}

	return ans
}

func projectionArea1(grid [][]int) int {
	ans := 0
	// ansRow := make([]int, len(grid))
	ansCol := make([]int, len(grid))
	for _, row := range grid {
		max := 0
		for y, item := range row {
			if item > 0 { // xy面的投影面积
				ans += 1
			}
			if item > max {
				max = item // 每一行的最大值
			}
			if item > ansCol[y] {
				ansCol[y] = item // 每一列的最大值
			}
		}
		ans += max
	}
	for i := 0; i < len(grid); i++ {
		ans += ansCol[i]
	}
	return ans
}

func main() {
	input := [][]int{
		[]int{1, 0},
		[]int{0, 2},
	}
	ans := projectionArea(input)
	fmt.Println(ans)
}
