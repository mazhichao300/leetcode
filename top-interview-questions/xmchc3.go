package main

import "math/rand"

type Solution struct {
	Origin []int
	Arr    []int
	Len    int
}

func Constructor(nums []int) Solution {
	arr := make([]int, len(nums))
	copy(arr, nums)
	return Solution{
		Origin: nums,
		Arr:    arr,
		Len:    len(arr),
	}
}

func (this *Solution) Reset() []int {
	return this.Origin
}

func (this *Solution) Shuffle() []int {
	for i := this.Len - 1; i >= 0; i-- {
		idx := rand.Intn(i + 1)
		this.Arr[i], this.Arr[idx] = this.Arr[idx], this.Arr[i]
	}
	return this.Arr
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
