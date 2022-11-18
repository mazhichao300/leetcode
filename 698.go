package main

import "sort"

func canPartitionKSubsets(nums []int, k int) bool {
	total := 0
	for _, n := range nums {
		total += n
	}
	if total%k != 0 {
		return false
	}
	target := total / k
	ans := false
	groups := make([]int, k)
	sort.Ints(nums)
	dfs(nums, groups, 0, target, &ans)
	return ans
}

func dfs(nums, groups []int, i, target int, ans *bool) {
	if *ans {
		return
	}
	if i >= len(nums) {
		// flag := true
		// for _, n := range groups {
		// 	if n != target {
		// 		flag = false
		// 		break
		// 	}
		// }
		// if flag {
		// 	*ans = true
		// }
		*ans = true
		return
	}
	for idx, n := range groups {
		if idx > 0 && groups[idx] == groups[idx-1] {
			continue
		}
		if n+nums[i] <= target {
			groups[idx] = n + nums[i]
			dfs(nums, groups, i+1, target, ans)
			groups[idx] = n
		}
	}
}
