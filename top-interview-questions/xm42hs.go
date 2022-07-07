package main

import "fmt"

func rotate(nums []int, k int) {
	l := len(nums)
	k %= l
	if k == 0 {
		return
	}
	t := nums[0]
	idx := 0
	for i := 1; i < l; i++ {
		pre := (idx - k + l) % l
		fmt.Println(idx, pre)
		nums[idx] = nums[pre]
		idx = pre
	}
	nums[idx] = t
}

func main() {
	a := []int{-1, -100, 3, 99}
	rotate(a, 2)
	fmt.Println(a)
}
