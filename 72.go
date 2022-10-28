package main

func minDistance(word1 string, word2 string) int {
	m := make([][]int, len(word1))
	for i := 0; i < len(word1); i++ {
		m[i] = make([]int, len(word2))
		for j := 0; j < len(word2); j++ {
			m[i][j] = -1
		}
	}
	return dp(word1, word2, len(word1)-1, len(word2)-1, m)
}

func dp(word1, word2 string, i, j int, m [][]int) int {
	if i < 0 {
		return j + 1
	}
	if j < 0 {
		return i + 1
	}
	if m[i][j] != -1 {
		return m[i][j]
	}
	if word1[i] == word2[j] {
		m[i][j] = dp(word1, word2, i-1, j-1, m)
		return m[i][j]
	}
	a := dp(word1, word2, i-1, j-1, m)
	if b := dp(word1, word2, i, j-1, m); b < a {
		a = b
	}
	if c := dp(word1, word2, i-1, j, m); c < a {
		a = c
	}
	m[i][j] = a + 1
	return a + 1
}
