package main

import "fmt"

func rotate(nums []int, k int) {
	l := len(nums)
	k %= l
	if k == 0 {
		return
	}

	cnt := l

	idx := 0
	for {
		start := idx
		t := nums[idx]
		for {
			cnt--
			pre := (idx - k + l) % l
			// fmt.Println(idx, pre)
			nums[idx] = nums[pre]

			if pre == start {
				nums[idx] = t
				break
			}
			idx = pre

		}
		idx++
		if cnt <= 0 {
			break
		}
	}
}

func main() {
	a := []int{-1, -100, 3, 99, 3}
	rotate(a, 2)
	fmt.Println(a)
}
