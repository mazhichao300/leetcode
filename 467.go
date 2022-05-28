package main

import "fmt"

func findSubstringInWraproundString(p string) int {
	dp := make([]int, 26)

	for i := 0; i < len(p); i++ {
		idx := int(p[i] - 'a')
		if i == 0 {
			dp[idx] = 1
		} else {
			idxPre := int(p[i-1] - 'a')
			tmp := 1
			// dp[idx] = 1
			if p[i-1]+1 == p[i] || (p[i-1] == 'z' && p[i] == 'a') {
				tmp += dp[idxPre]
				// dp[idxPre] = 0
			}
			if tmp > dp[idx] {
				dp[idx] = tmp
			}
		}
	}
	ans := 0
	for i := 0; i < len(dp); i++ {
		// n := dp[i]
		// ans += n * (n + 1) / 2
		ans += dp[i]
	}

	fmt.Println(dp)

	return ans
}

func main() {
	ans := findSubstringInWraproundString("cdefghefghijklmnopqrstuvwxmnijklmnopqrstuvbcdefghijklmnopqrstuvwabcddefghijklfghijklmabcdefghijklmnopqrstuvwxymnopqrstuvwxyz")
	fmt.Println(ans)
}
