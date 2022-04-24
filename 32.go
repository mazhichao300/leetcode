package main

import "fmt"

var ans int

func longestValidParentheses(s string) int {
	ans = 0
	f(s, 0, len(s)-1)
	return ans
}

func f(s string, i, j int) bool {
	if i >= j || (j-i+1)%2 == 1 {
		return false
	}
	if i+1 == j {
		if s[i] == '(' && s[j] == ')' {
			if j-i+1 > ans {
				ans = j - i + 1
			}
			return true
		}
		return false
	}
	if s[i] == '(' && s[j] == ')' && f(s, i+1, j-1) {
		if j-i+1 > ans {
			ans = j - i + 1
		}
		return true
	}
	if s[i] == '(' && s[i+1] == ')' && f(s, i+2, j) {
		if j-i+1 > ans {
			ans = j - i + 1
		}
		return true
	}
	if s[j-1] == '(' && s[j] == ')' && f(s, i, j-2) {
		if j-i+1 > ans {
			ans = j - i + 1
		}
		return true
	}
	return false
}

func main() {
	ans := longestValidParentheses("(()")
	fmt.Println(ans)
}
