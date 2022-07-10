package main

const INT_MAX = int(^uint(0) >> 1)

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	min := nums[0]
	max := INT_MAX
	for _, n := range nums {
		if n <= min {
			min = n
		} else if n <= max {
			max = n
		} else {
			return true
		}
	}
	return false
}
