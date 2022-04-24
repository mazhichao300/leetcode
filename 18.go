package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	e := make(map[string]bool, 0)
	ans := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			start := j + 1
			end := len(nums) - 1

			for start < end {
				if nums[i]+nums[j]+nums[start]+nums[end] == target {
					key := fmt.Sprintf("%d_%d_%d_%d", nums[i], nums[j], nums[start], nums[end])
					if _, ok := e[key]; !ok {
						ans = append(ans, []int{nums[i], nums[j], nums[start], nums[end]})
						e[key] = true
						break
					}
					//start++
				} else if nums[i]+nums[j]+nums[start]+nums[end] < target {
					start++
				} else {
					end--
				}
			}
		}
	}
	return ans
}

func main() {
	ans := fourSum([]int{2, 2, 2, 2, 2}, 8)
	fmt.Println(ans)
}
