package main

import "fmt"

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		length := n - 2*i
		step := length - 1
		start := i
		end := n - i

		for j := i; j < end-1; j++ {
			x := j
			y := i
			value := matrix[x][y]

			for cnt := 0; cnt < 4; cnt++ {
				nextX, nextY := getNext(x, y, step, start, end)
				value, matrix[nextX][nextY] = matrix[nextX][nextY], value
				x, y = nextX, nextY
			}
		}
	}
}

func getNext(x, y int, step, start, end int) (int, int) {
	for step > 0 {
		step--
		if x == start && y < end-1 { // 第一行,未走到最后，y++
			y++
		} else if y == end-1 && x < end-1 { // 最右一列，x未走到最后，x++
			x++
		} else if x == end-1 && y > start { // 最后一行，y未走到最前
			y--
		} else {
			x--
		}
	}
	return x, y
}

func main() {
	input := [][]int{
		[]int{5, 1, 9, 11},
		[]int{2, 4, 8, 10},
		[]int{13, 3, 6, 7},
		[]int{15, 14, 12, 16},
	}
	rotate(input)
	fmt.Println(input)

}
