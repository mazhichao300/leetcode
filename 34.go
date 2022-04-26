package main

import "fmt"

func findLeftOrRight(nums []int, target int, leftOrRight string) int {
	start := 0
	end := len(nums) - 1
	ans := -1

	for start <= end {
		mid := (start + end) / 2
		if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			ans = mid
			if leftOrRight == "left" {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}
	return ans
}

func searchRange(nums []int, target int) []int {
	// return []int{findStart(nums, target), findEnd(nums, target)}
	return []int{findLeftOrRight(nums, target, "left"), findLeftOrRight(nums, target, "right")}

}

func findStart(nums []int, target int) int {
	length := len(nums)
	start := 0
	end := length - 1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target && (mid == 0 || nums[mid-1] != target) {
			return mid
		} else if nums[mid] < target && mid+1 < length && nums[mid+1] == target {
			return mid + 1
		} else if target <= nums[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

func findEnd(nums []int, target int) int {
	length := len(nums)
	start := 0
	end := length - 1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target && (mid == length-1 || nums[mid+1] != target) {
			return mid
		} else if nums[mid] > target && mid-1 >= 0 && nums[mid-1] == target {
			return mid - 1
		} else if target >= nums[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

func main() {
	ans := searchRange([]int{5, 5, 7, 8, 10, 10}, 10)
	fmt.Println(ans)
}
