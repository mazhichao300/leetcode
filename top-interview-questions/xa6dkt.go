package main

func titleToNumber(columnTitle string) int {
	ans := 0

	for _, c := range columnTitle {
		ans = ans*26 + int(c-'A'+1)
	}

	return ans
}
