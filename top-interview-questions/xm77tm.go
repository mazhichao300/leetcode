//https://leetcode.cn/leetbook/read/top-interview-questions/xm77tm/
package main

func majorityElement(nums []int) int {
	ans := 0
	cnt := 0
	for _, i := range nums {
		if cnt == 0 {
			cnt = 1
			ans = i
		} else {
			if ans == i {
				cnt++
			} else {
				cnt--
			}
		}
	}
	return ans
}
