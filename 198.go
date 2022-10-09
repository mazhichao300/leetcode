package main

var dp []int

func rob(nums []int) int {
	if dp == nil {
		dp = make([]int, 101)
	}
	size := len(nums)
	dp[0] = nums[0]
	if size > 1 {
		dp[1] = max(dp[0], nums[1])
	}

	for i := 2; i < size; i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}
	return dp[size-1]
}

func rob1(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	dp := make([][2]int, size)
	dp[0][0] = 0
	dp[0][1] = nums[0]

	for i := 1; i < size; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = nums[i] + dp[i-1][0]
	}
	return max(dp[size-1][0], dp[size-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
