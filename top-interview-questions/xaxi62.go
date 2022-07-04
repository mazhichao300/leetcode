package main

import "fmt"

var ans [][]string

func partition(s string) [][]string {
	ans = make([][]string, 0)
	dfs(s, 0, []string{})
	return ans
}

func dfs(s string, index int, strs []string) {
	if index >= len(s) {
		ans = append(ans, append([]string{}, strs...))
		return
	}
	for i := index; i < len(s); i++ {
		str := s[index : i+1]
		if isVlalid(str) {
			dfs(s, i+1, append(strs, str))
		}
	}
}

func isVlalid(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	s := "abbcac"
	a := partition(s)
	fmt.Println(a)
}
