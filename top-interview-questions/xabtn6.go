package main

func findDuplicate(nums []int) int {
	i := 0
	for i < len(nums) {
		if nums[i] == 0 {
			return i
		}
		next := nums[i]
		nums[i] = 0
		i = next
	}
	return 0
}
