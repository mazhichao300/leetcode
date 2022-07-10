package main

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))

	multi := 1
	for i := 0; i < len(nums); i++ {
		ans[i] = multi
		multi *= nums[i]
	}

	multi = 1
	for i := len(nums) - 1; i >= 0; i-- {
		ans[i] *= multi
		multi *= nums[i]
	}
	return ans
}
