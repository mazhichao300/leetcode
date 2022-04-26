package main

import "fmt"

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[start] <= nums[mid] && nums[mid] <= nums[end] { // 正常顺序的数组
			if nums[mid] < target {
				start = mid + 1
			} else {
				end = mid - 1
			}
		} else if nums[start] >= nums[mid] && nums[mid] <= nums[end] {
			if nums[mid] > target || target > nums[end] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if target > nums[mid] || target < nums[start] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}

func main() {
	ans := search([]int{4, 5, 6, 7, 0, 1, 2}, 44)
	fmt.Println(ans)
}
