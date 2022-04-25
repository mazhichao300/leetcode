package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Solution struct {
	Data map[int][]int
}

func Constructor(nums []int) Solution {
	ans := Solution{}
	ans.Data = make(map[int][]int, 0)
	for idx, n := range nums {
		if _, ok := ans.Data[n]; !ok {
			ans.Data[n] = make([]int, 0)
		}
		ans.Data[n] = append(ans.Data[n], idx)
	}
	return ans
}

func (this *Solution) Pick(target int) int {
	arr := this.Data[target]
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(arr))
	return arr[idx]
}

func main() {
	nums := []int{1, 2, 3, 3, 3}
	solution := Constructor(nums)
	// pick(3) 应该返回索引 2,3 或者 4。每个索引的返回概率应该相等。
	a := solution.Pick(3)
	fmt.Println(a)

	// pick(1) 应该返回 0。因为只有nums[0]等于1。
	a = solution.Pick(1)

	fmt.Println(a)

}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
