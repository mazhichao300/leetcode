package main

import "fmt"

var dp []int

const MAX_INT = 20000

func jump(nums []int) int {
	dp = make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		dp[i] = -1
	}

	return f(nums, len(nums)-1)
}

func f(nums []int, i int) int {
	if dp[i] != -1 {
		return dp[i]
	}
	min := MAX_INT
	for x := 0; x < i; x++ {
		if x+nums[x] >= i {
			tmp := f(nums, x)
			if tmp < min {
				min = tmp
			}
		}
	}
	if min != MAX_INT {
		dp[i] = min + 1
	} else {
		dp[i] = MAX_INT
	}
	return dp[i]
}

func main() {
	ans := jump([]int{2})
	fmt.Println(ans)
}
