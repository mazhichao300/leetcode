package main

import (
	"fmt"
)

func findClosest(words []string, word1 string, word2 string) int {
	m := make(map[string][]int)
	for i, s := range words {
		if _, ok := m[s]; !ok {
			m[s] = []int{}
		}
		m[s] = append(m[s], i)
	}
	return getArrDist(m[word1], m[word2], len(words))
}

func getArrDist(a, b []int, ans int) int {
	if len(b) > ans {
		ans = len(b)
	}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		t := abs(a[i], b[j])
		if t < ans {
			ans = t
		}
		if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}
	return ans
}

func abs(a, b int) int {
	a -= b
	if a < 0 {
		a = -a
	}
	return a
}

func main() {
	words := []string{"I", "am", "a", "student", "from", "a", "university", "in", "a", "city"}
	word1 := "a"
	word2 := "student"
	ans := findClosest(words, word1, word2)
	fmt.Println(ans)
}
