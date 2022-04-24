package main

import "fmt"

var ans [][]int

func combinationSum(candidates []int, target int) [][]int {
	ans = make([][]int, 0)
	f(candidates, 0, target, 0, []int{})
	return ans
}

func f(candidates []int, i, target, current int, tmp []int) {
	if current == target {
		c := make([]int, len(tmp))
		copy(c, tmp)
		ans = append(ans, c)
		return
	}
	if current > target || i >= len(candidates) {
		return
	}
	v := candidates[i]
	cnt := (target - current) / v
	for x := 0; x <= cnt; x++ {
		add := v * x
		if add > 0 {
			tmp = append(tmp, v)
		}
		f(candidates, i+1, target, current+add, tmp)
	}
}

func main() {
	//ans := combinationSum([]int{2, 7, 6, 3, 5, 1}, 9)
	ans := combinationSum([]int{2, 3}, 9)
	fmt.Println(ans)
}
