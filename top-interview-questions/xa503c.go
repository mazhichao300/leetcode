package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]bool)
	for _, d := range wordDict {
		dict[d] = true
	}
	// return dfs(s, 0, len(s), dict)
	arr := []int{}
	for i := 0; i < len(s); i++ {
		if _, ok := dict[s[0:i+1]]; ok {
			arr = append(arr, i)
			continue
		}
		for _, j := range arr {
			if j+1 <= i {
				if _, ok := dict[s[j+1:i+1]]; ok {
					arr = append(arr, i)
					break
				}
			}
		}
	}
	return len(arr) > 0 && arr[len(arr)-1] == len(s)-1
}

func dfs(s string, start, len int, dict map[string]bool) bool {
	if start >= len {
		return true
	}
	for i := start + 1; i <= len; i++ {
		tmp := s[start:i]
		// fmt.Println(tmp)
		if _, ok := dict[tmp]; ok && dfs(s, i, len, dict) {
			return true
		}
	}
	return false
}

func main() {
	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	wordDict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	// s := "applepenapplea"
	// wordDict := []string{"apple", "pen"}
	ans := wordBreak(s, wordDict)
	fmt.Println(ans)
}
