package main

import "fmt"

func getMaxAndMin(a, b, c int) (int, int) {
	max, min := a, a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return max, min
}

func maxProduct(nums []int) int {
	max, min, ans := 1, 1, nums[0]

	for _, n := range nums {
		t1 := max * n
		t2 := min * n

		max, min = getMaxAndMin(t1, t2, n)
		if ans < max {
			ans = max
		}
	}
	return ans
}

func main() {
	a := []int{2, 3, -2, 4}
	ans := maxProduct(a)
	fmt.Println(ans)
}
