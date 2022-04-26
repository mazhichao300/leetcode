package main

import "fmt"

func firstMissingPositive(nums []int) int {
	for _, n := range nums {
		pos := n - 1
		for pos >= 0 && pos < len(nums) && nums[pos] != pos+1 {
			next := nums[pos] - 1
			nums[pos] = pos + 1
			pos = next
		}
	}

	for idx, n := range nums {
		if idx != n-1 {
			return idx + 1
		}
	}
	return nums[len(nums)-1] + 1
}

func main() {
	ans := firstMissingPositive([]int{3, 4, -1, 1})
	fmt.Println(ans)
}
