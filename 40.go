package main

import "fmt"

type item struct {
	val int
	cnt int
}

var ans [][]int

func combinationSum2(candidates []int, target int) [][]int {
	m := map[int]int{}
	for _, i := range candidates {
		if v, ok := m[i]; ok {
			m[i] = v + 1
		} else {
			m[i] = 1
		}
	}
	candi := []item{}

	for k, v := range m {
		candi = append(candi, item{k, v})
	}
	ans = make([][]int, 0)
	f(candi, 0, target, 0, []int{})
	return ans
}

func f(candidates []item, i, target, current int, tmp []int) {
	if current == target {
		t := make([]int, len(tmp))
		copy(t, tmp)
		ans = append(ans, t)
		return
	}
	if current > target || i >= len(candidates) {
		return
	}

	v := candidates[i].val
	cnt := candidates[i].cnt
	for x := 0; x <= cnt; x++ {
		add := v * x
		if add+current <= target {
			if add > 0 {
				tmp = append(tmp, v)
			}
			f(candidates, i+1, target, current+add, tmp)
		} else {
			break
		}
	}
}

func main() {
	ans := combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	fmt.Println(ans)
}
