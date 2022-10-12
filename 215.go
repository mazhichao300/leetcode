package main

import "fmt"

func adjustDown(heap []int, start, end int) {
	father := start
	son := 2*start + 1
	for son <= end {
		if son+1 <= end && heap[son+1] < heap[son] {
			son = son + 1
		}
		if heap[father] > heap[son] {
			heap[father], heap[son] = heap[son], heap[father]
		} else {
			break
		}
		father = son
		son = 2*father + 1
	}
}

func findKthLargestHeap(nums []int, k int) int {
	heap := make([]int, k)
	for i := 0; i < len(nums); i++ {
		if i < k {
			heap[i] = nums[i]
			if i == k-1 {
				for j := k / 2; j >= 0; j-- {
					adjustDown(heap, j, k-1)
				}
			}
		} else if nums[i] > heap[0] {
			heap[0] = nums[i]
			adjustDown(heap, 0, k-1)
		}
	}
	return heap[0]
}

func partition(nums []int, start, end int) int {
	p := nums[start]
	// fmt.Println(start, end, p, "------")
	for start < end {
		for start < end && nums[end] < p {
			end--
		}
		if start < end {
			nums[start] = nums[end]
		}
		for start < end && nums[start] >= p {
			start++
		}
		if start < end {
			nums[end] = nums[start]
		}
	}
	nums[start] = p
	// fmt.Println(nums)
	return start
}

func quickSort(nums []int, start, end, k int) int {
	if start >= end {
		return -1
	}

	if p := partition(nums, start, end); p == k-1 {
		return nums[p]
	} else if p > k-1 {
		return quickSort(nums, start, p-1, k)
	} else {
		return quickSort(nums, p+1, end, k)
	}
}

func findKthLargest(nums []int, k int) int {
	quickSort(nums, 0, len(nums)-1, k)
	return nums[k-1]
}

// func quickSort(nums []int, start, end int) {
// 	if start >= end {
// 		return
// 	}
// 	p := partition(nums, start, end)
// 	quickSort(nums, start, p-1)
// 	quickSort(nums, p+1, end)
// }

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	fmt.Println(findKthLargest(nums, k))
	fmt.Println(nums)
}
