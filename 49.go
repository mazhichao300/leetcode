package main

import (
	"fmt"
	"sort"
)

var zhishu []int = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}

func groupAnagrams(strs []string) [][]string {
	m := map[int][]string{}
	for _, s := range strs {
		num := 1
		for _, i := range s {
			num *= zhishu[int(i-'a')]
		}
		if _, ok := m[num]; !ok {
			m[num] = []string{}
		}
		m[num] = append(m[num], s)
	}
	ans := make([][]string, len(m))
	idx := 0
	for _, v := range m {
		ans[idx] = v
		idx++
	}
	return ans
}

func groupAnagramsOld(strs []string) [][]string {
	m := map[string][]string{}
	for _, s := range strs {
		b := []byte(s)
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		sortS := string(b)
		if _, ok := m[sortS]; !ok {
			m[sortS] = []string{}
		}
		m[sortS] = append(m[sortS], s)
	}
	ans := make([][]string, len(m))
	idx := 0
	for _, v := range m {
		ans[idx] = v
		idx++
	}
	return ans
}

func main() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "baz"}
	ans := groupAnagrams(input)
	fmt.Println(ans)
}
