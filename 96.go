package main

var ans []int

func numTrees(n int) int {
	if ans == nil {
		ans = make([]int, 20)
		ans[0] = 1
		ans[1] = 1
		for i := 2; i < 20; i++ {
			cnt := 0
			for mid := 1; mid <= i; mid++ {
				cnt += ans[mid-1] * ans[i-mid]
			}
			ans[i] = cnt
		}
	}
	return ans[n]
}
