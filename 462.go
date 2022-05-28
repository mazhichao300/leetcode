package main

import (
	"fmt"
	"sort"
)

func minMoves2(nums []int) int {
	sort.Ints(nums)
	mid := len(nums) / 2
	median := nums[mid]
	if len(nums)%2 == 0 {
		median += nums[mid-1]
		median /= 2
	}
	ans := 0
	for _, v := range nums {
		diff := median - v
		if diff < 0 {
			diff = -diff
		}
		ans += diff
	}
	return ans
}

func main() {
	ans := minMoves2([]int{1, 2, 10, 9})
	fmt.Println(ans)
}
