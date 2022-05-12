package main

import "fmt"

var ans int

var m []int

func longestValidParentheses(s string) int {
	ans := 0
	dp := make([]int, len(s))

	for i := 1; i < len(s); i++ {
		if s[i] == '(' {
			dp[i] = 0
		} else if s[i-1] == '(' {
			dp[i] = 2
			if i-2 > 0 {
				dp[i] += dp[i-2]
			}
		} else if idx := i - dp[i-1] - 1; idx >= 0 && s[idx] == '(' {
			dp[i] = 2 + dp[i-1]
			if idx-1 > 0 {
				dp[i] += dp[idx-1]
			}
		}
		if ans < dp[i] {
			ans = dp[i]
		}
	}
	return ans
}

func longestValidParenthesesRe(s string) int {
	m = make([]int, len(s))
	for i := 0; i < len(s); i++ {
		m[i] = -1
	}
	ans = 0
	for i := 0; i < len(s); i++ {
		f(s, i)
	}
	return ans
}

func f(s string, i int) int {
	if i < 0 {
		return 0
	}
	key := i
	if m[key] != -1 {
		return m[key]
	}

	if i == 0 || s[i] == '(' {
		m[i] = 0
		return 0
	}

	length := 0
	if s[i-1] == '(' {
		length = 2 + f(s, i-2)
	} else if idx := i - 1 - f(s, i-1); idx >= 0 && s[idx] == '(' {
		length = 2 + f(s, i-1)
		if idx-1 > 0 {
			length += f(s, idx-1)
		}
	}

	m[i] = length
	if ans < length {
		ans = length
	}
	return length
}

func main() {
	// s := ")(((((()())()()))()(()))("
	s := "()(()(((((())())((()))(()((())()(()()(()((()))()))))())))))())()())))(()()()())((()()()))()(()(()))())(((()))())(()((()))(()(()))()"
	// s = ")()())"
	ans := longestValidParentheses(s)
	fmt.Println(ans)
}
