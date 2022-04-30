package main

import "fmt"

func smallestRangeI(nums []int, k int) int {
	max := -1
	min := 20000

	for _, n := range nums {
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}
	if min+k >= max-k {
		return 0
	}
	return max - k - min - k
}

func main() {
	ans := smallestRangeI([]int{1, 3, 6}, 3)
	fmt.Println(ans)
}
