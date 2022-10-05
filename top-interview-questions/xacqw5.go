package main

func findPeakElementLinar(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return len(nums) - 1
}

func findPeakElement(nums []int) int {
	start, end := 0, len(nums)-1
	for start < end {
		mid := (start + end) / 2
		if nums[mid] < nums[mid+1] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}
