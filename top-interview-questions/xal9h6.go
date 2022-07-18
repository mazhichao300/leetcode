package main

import "fmt"

func makeHeap(nums []int, n int) {
	for i := (n - 1) / 2; i >= 0; i-- {
		start := i

		for {
			left := 2*start + 1
			right := left + 1
			if left >= n {
				break
			}
			if right < n && nums[right] < nums[left] {
				left = right
			}
			if nums[start] > nums[left] {
				nums[start], nums[left] = nums[left], nums[start]
			}
			start = left
		}
	}
}

func adjustHeap(nums []int, n int) {
	i := 0
	for {
		left := 2*i + 1
		right := left + 1
		if left >= n {
			break
		}
		if right < n && nums[right] < nums[left] {
			left = right
		}
		if nums[i] > nums[left] {
			nums[i], nums[left] = nums[left], nums[i]
			i = left
		} else {
			break
		}
	}
}

func findKthLargest(nums []int, k int) int {
	makeHeap(nums, k)
	for i := k; i < len(nums); i++ {
		if nums[i] > nums[0] {
			nums[0] = nums[i]
			adjustHeap(nums, k)
		}
	}
	return nums[0]
}

func main() {
	a := []int{3, 2, 1, 5, 6, 4}
	k := 2
	ans := findKthLargest(a, k)
	fmt.Println(ans)
}
