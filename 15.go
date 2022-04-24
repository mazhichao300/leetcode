package main

import (
	"fmt"
	"sort"
)

func threeSumOld(nums []int) [][]int {
	ans := make([][]int, 0)
	exist := map[string]bool{}
	m := map[int]int{}
	length := len(nums)
	for _, i := range nums {
		if _, ok := m[i]; ok {
			m[i]++
		} else {
			m[i] = 1
		}
	}

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			target := 0 - nums[i] - nums[j]
			n := 0
			// target不能是i,j
			if target == nums[i] {
				n++
			}
			if target == nums[j] {
				n++
			}
			// 找到有效的
			if t, ok := m[target]; ok && t > n {
				a := nums[i]
				b := nums[j]
				c := target
				if a > b {
					a, b = b, a
				}
				if a > c {
					a, c = c, a
				}
				if b > c {
					b, c = c, b
				}
				key := fmt.Sprintf("%d_%d_%d", a, b, c)
				if _, ok := exist[key]; ok {
					continue
				}
				exist[key] = true
				ans = append(ans, []int{a, b, c})
			}
		}
	}
	return ans
}

func threeSum(nums []int) [][]int {
	length := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	for i := 0; i < length; i++ {
		left := i + 1
		right := length - 1
		target := 0 - nums[i]
		for left < right {
			if nums[left]+nums[right] == target {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if nums[left]+nums[right] > target {
				right--
			} else {
				left++
			}
		}
		for i+1 < length && nums[i] == nums[i+1] {
			i++
		}
	}
	return ans
}

func main() {
	s := []int{-1, 0, 1, 2, -1, -4}
	ans := threeSum(s)
	fmt.Println(ans)
}
