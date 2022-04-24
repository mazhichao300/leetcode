package main

import "fmt"

var ans [][]int
var tmp []int
var e map[int]int

func permuteUnique(nums []int) [][]int {
	ans = make([][]int, 0)
	e = make(map[int]int, 0)
	tmp = make([]int, len(nums))

	for _, v := range nums {
		if _, ok := e[v]; ok {
			e[v]++
		} else {
			e[v] = 1
		}
	}

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
	for k, v := range e {
		if v > 0 {
			tmp[i] = k
			e[k]--
			f(nums, i+1, tmp)
			e[k]++
		}
	}
}

func main() {
	ans := permuteUnique([]int{1, 2, 1})
	fmt.Println(ans)
}
