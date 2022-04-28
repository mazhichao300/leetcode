package main

import "fmt"

func sortArrayByParity(nums []int) []int {
	i := 0
	j := len(nums) - 1
	for i < j {
		if nums[i]&1 == 1 && nums[j]&1 == 0 {
			nums[i], nums[j] = nums[j], nums[i]
		}
		if nums[i]&1 == 0 {
			i++
		}
		if nums[j]&1 == 1 {
			j--
		}
	}
	return nums
}

func main() {
	ans := sortArrayByParity([]int{0, 2, 1, 4})
	fmt.Println(ans)
}
