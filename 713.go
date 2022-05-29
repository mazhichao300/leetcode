package main

import "fmt"

func numSubarrayProductLessThanK(nums []int, k int) int {
	ans := 0
	start := 0
	end := 0
	// lastStart := -1
	lastEnd := -1
	multi := nums[start]
	for start < len(nums) && end < len(nums) {
		if multi < k {
			end++
			multi *= nums[end]
			if multi >= k || end == len(nums) {
				n := end - start
				ans += n * (n + 1) / 2
				if lastEnd >= start {
					n = lastEnd - start + 1
					ans -= n * (n + 1) / 2
				}
				lastEnd = end
				for multi >= k {
					multi /= nums[start]
					start++
				}
			}
		} else {
			start++
			end = start
			multi = nums[start]
		}
	}

	return ans
}

func numSubarrayProductLessThanK1(nums []int, k int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		multi := 1
		for j := i; j < len(nums); j++ {
			multi *= nums[j]
			if multi < k {
				ans++
			} else {
				break
			}
		}
	}
	return ans
}

func main() {
	ans := numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100)
	fmt.Println(ans)
}
