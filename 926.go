package main

// TODO

func minFlipsMonoIncr(s string) int {
	dp := make([]int, len(s))

	for i, c := range s {
		if i == 0 {
			dp[i] = 0
		} else if c == '1' {
			dp[i] = dp[i-1]
		} else {

		}
	}
}

func main() {

}
