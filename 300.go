package main

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	ans := 1
	for i := 1; i < len(nums); i++ {
		tmp := 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				if dp[j]+1 > tmp {
					tmp = dp[j] + 1
				}
			}
		}
		dp[i] = tmp
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
