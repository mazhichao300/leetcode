package main

import (
	"fmt"
	"strconv"
)

func calculate(s string) int {
	s += "+0"
	nums := []int{}
	op := []rune{}
	tmp := []rune{}
	for _, c := range s {
		if c == ' ' {
			continue
		}
		if c >= '0' && c <= '9' {
			tmp = append(tmp, c)
			continue
		}
		num, _ := strconv.Atoi(string(tmp))

		if len(op) > 0 && op[len(op)-1] == '*' {
			nums[len(nums)-1] *= num
			op[len(op)-1] = c
		} else if len(op) > 0 && op[len(op)-1] == '/' {
			nums[len(nums)-1] /= num
			op[len(op)-1] = c
		} else {
			nums = append(nums, num)
			op = append(op, c)
		}

		tmp = []rune{}
	}
	ans, _ := strconv.Atoi(string(tmp))
	nums = append(nums, ans)
	ans = nums[0]

	// fmt.Println(nums)
	// fmt.Println(string(op))

	for i := 1; i < len(nums); i++ {
		if op[i-1] == '+' {
			ans += nums[i]
		} else if op[i-1] == '-' {
			ans -= nums[i]
		} else if op[i-1] == '*' {
			ans *= nums[i]
		} else {
			ans /= nums[i]
		}
	}
	return ans
}

func main() {
	s := " 3+2*2"
	ans := calculate(s)
	fmt.Println(ans)
}
