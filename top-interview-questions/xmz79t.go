package main

func majorityElement(nums []int) int {
	ans, cnt := -1, 0

	for _, n := range nums {
		if cnt == 0 || ans == n {
			ans = n
			cnt++
		} else {
			cnt--
		}
	}
	return ans
}
