package main

import (
	"fmt"
)

func buildNext(s string) []int {
	ans := make([]int, len(s))
	ans[0] = 0
	for i := 1; i < len(s); i++ {
		pre := ans[i-1]
		for pre >= 0 {
			if s[i] == s[pre] {
				ans[i] = pre + 1
				break
			}
			if pre == 0 {
				break
			}
			pre = ans[pre-1]
		}
	}
	return ans
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	next := buildNext(needle)
	tar := 0
	pos := 0
	for tar < len(haystack) {
		if haystack[tar] == needle[pos] {
			tar++
			pos++
		} else if pos > 0 {
			pos = next[pos-1]
		} else {
			tar++
		}

		if pos == len(needle) {
			return tar - pos
		}
	}
	return -1
}

func main() {
	a := "a"
	b := "ll"
	ans := strStr(a, b)
	fmt.Println(ans)
}
