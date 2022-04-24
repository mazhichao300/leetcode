package main

import (
	"fmt"
	"sort"
)

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	length := len(nums)
	ans := 10000000
	for i := 0; i < length; i++ {
		start := i + 1
		end := length - 1

		for start < end {
			sum := nums[i] + nums[start] + nums[end]

			if abs(sum-target) < abs(ans-target) {
				ans = sum
			}
			if sum == target {
				return sum
			} else if sum > target {
				end--
			} else {
				start++
			}
		}
	}
	return ans
}

func main() {
	ans := threeSumClosest([]int{1, 2, 5, 10, 11}, 12)
	fmt.Println(ans)
}
