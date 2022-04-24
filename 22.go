package main

import "fmt"

var ans []string
var temp []byte

func generateParenthesis(n int) []string {
	ans = []string{}
	temp = make([]byte, 2*n)
	f(0, n*2, n, n)
	return ans
}

func f(i, length, left, right int) {
	if i >= length {
		ans = append(ans, string(temp))
		return
	}
	if left > 0 {
		temp[i] = '('
		f(i+1, length, left-1, right)
	}
	if right > left && right > 0 {
		temp[i] = ')'
		f(i+1, length, left, right-1)
	}
}

func main() {
	ans := generateParenthesis(0)
	fmt.Println(ans)
}
