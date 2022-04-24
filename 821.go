package main

import "fmt"

func shortestToChar(s string, c byte) []int {
	ans := make([]int, len(s))
	cArr := []int{}
	for i := 0; i < len(ans); i++ {
		ans[i] = len(s) + 1
		if s[i] == c {
			cArr = append(cArr, i)
		}
	}

	for i := 0; i < len(s); i++ {
		if s[i] == c {
			cnt := len(s) - i
			if i > cnt {
				cnt = i
			}
			for step := 0; step <= cnt; step++ {
				if i+step < len(s) && ans[i+step] > step {
					ans[i+step] = step
				}
				if i-step >= 0 && ans[i-step] > step {
					ans[i-step] = step
				}
			}
		}
	}

	return ans
}

func shortestToCharOld(s string, c byte) []int {
	ans := make([]int, len(s))
	for i := 0; i < len(ans); i++ {
		ans[i] = len(s) + 1
	}

	for i := 0; i < len(s); i++ {
		if s[i] == c {
			cnt := len(s) - i
			if i > cnt {
				cnt = i
			}
			for step := 0; step <= cnt; step++ {
				if i+step < len(s) && ans[i+step] > step {
					ans[i+step] = step
				}
				if i-step >= 0 && ans[i-step] > step {
					ans[i-step] = step
				}
			}
		}
	}

	return ans
}

func main() {
	ans := shortestToChar("aaab", 'b')
	fmt.Println(ans)
}
