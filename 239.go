package main

func maxSlidingWindow(nums []int, k int) []int {
	ans := []int{}
	queue := []int{}
	for idx := 0; idx < len(nums); idx++ {
		for len(queue) > 0 && queue[len(queue)-1] < nums[idx] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[idx])

		if idx >= k-1 {
			ans = append(ans, queue[0])

			pre := idx - k + 1
			if nums[pre] == queue[0] {
				queue = queue[1:]
			}
		}
	}
	return ans
}
