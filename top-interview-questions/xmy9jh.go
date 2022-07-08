package main

func moveZeroes(nums []int) {
	idx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[idx] = nums[i]
			idx++
		}
	}
	for idx < len(nums) {
		nums[idx] = 0
		idx++
	}
}

func moveZeroes1(nums []int) {
	l := len(nums)
	i, j := 0, 0
	for j < l {
		for i < l && nums[i] != 0 {
			i++
		}

		j = i + 1
		for j < l && nums[j] == 0 {
			j++
		}
		if i < l && j < l {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}
