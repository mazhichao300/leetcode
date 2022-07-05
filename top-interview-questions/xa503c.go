package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s))

	for i := 0; i < len(s); i++ {
		for _, w := range wordDict {
			if i+1 < len(w) {
				continue
			}
			if s[i+1-len(w):i+1] == w && (i-len(w) < 0 || dp[i-len(w)]) {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)-1]
}

func wordBreak1(s string, wordDict []string) bool {
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
	wordDict := []string{"a", "b", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	// s := "applepenapplea"
	// wordDict := []string{"apple", "pen"}
	ans := wordBreak(s, wordDict)
	fmt.Println(ans)
}
