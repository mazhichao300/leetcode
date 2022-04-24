package main

import "fmt"

var ans [][]int
var tmp []int
var e []bool

func permute(nums []int) [][]int {
	ans = make([][]int, 0)
	e = make([]bool, len(nums))
	tmp = make([]int, len(nums))

	f(nums, 0, tmp)

	return ans
}

func f(nums []int, i int, tmp []int) {
	if i >= len(nums) {
		dst := make([]int, len(tmp))
		copy(dst, tmp)
		ans = append(ans, dst)
		return
	}
	for idx := 0; idx < len(nums); idx++ {
		if e[idx] == false {
			tmp[i] = nums[idx]
			e[idx] = true
			f(nums, i+1, tmp)
			e[idx] = false
		}
	}
}

func main() {
	ans := permute([]int{})
	fmt.Println(ans)
}
