package main

import "fmt"

func findSubstringInWraproundString(p string) int {
	dp := make([]int, 26)

	cnt := 0

	for i := 0; i < len(p); i++ {
		idx := int(p[i] - 'a')
		if i == 0 {
			dp[idx] = 1
			cnt = 1
		} else {
			// idxPre := int(p[i-1] - 'a')
			if p[i-1]+1 == p[i] || (p[i-1] == 'z' && p[i] == 'a') {
				cnt++
			} else {
				cnt = 1
			}
			if cnt > dp[idx] {
				dp[idx] = cnt
			}
		}
	}
	ans := 0
	for i := 0; i < len(dp); i++ {
		// n := dp[i]
		// ans += n * (n + 1) / 2
		ans += dp[i]
	}

	// fmt.Println(dp)

	return ans
}

func main() {
	ans := findSubstringInWraproundString("zab")
	fmt.Println(ans)
}
