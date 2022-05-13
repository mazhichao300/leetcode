package main

import (
	"fmt"
	"regexp"
)

func findSubstring(s string, words []string) []int {
	ans := []int{}
	wordsIndex := map[string][]int{}
	for _, x := range words {
		if _, ok := wordsIndex[x]; !ok {
			wordsIndex[x] = []int{}
		} else {
			continue
		}
		re := regexp.MustCompile(x)
		a := re.FindAllStringIndex(s, -1)
		if len(a) == 0 {
			return ans
		}

		for _, i := range a {
			wordsIndex[x] = append(wordsIndex[x], i[0])
		}
	}
	fmt.Println(wordsIndex)
	return []int{}
}

func main() {
	words := []string{
		"foo",
		"bar",
		"foo",
	}
	findSubstring("barfoothefoobarman", words)
}
