package main

import (
	"fmt"
	"strings"
)

// var ans []string

func wordBreak(s string, wordDict []string) []string {
	ans := make([]string, 0)
	dfs(s, 0, wordDict, []string{}, &ans)
	return ans
}

func dfs(s string, start int, wordDict []string, solve []string, ans *[]string) {
	length := len(s)
	if start >= length {
		return
	}
	for _, w := range wordDict {
		lenW := len(w)
		if lenW > length-start || s[start:start+lenW] != w {
			continue
		}
		solve = append(solve, w)
		if start+lenW == length {
			// ans = append(ans, append(append([]string{}, solve...), w))
			*ans = append(*ans, strings.Join(solve, " "))
		} else {
			// dfs(s, start+lenW, wordDict, append(solve, w))
			dfs(s, start+lenW, wordDict, solve, ans)
		}
		solve = solve[0 : len(solve)-1]
	}
}

func main() {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	ans := wordBreak(s, wordDict)
	fmt.Println(len(ans), ans)
	// 输出:["cats and dog","cat sand dog"]

}
