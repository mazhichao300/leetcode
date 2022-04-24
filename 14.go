package main

import "fmt"

func longestCommonPrefixOld(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	ans := []byte{}
	length := 300
	for _, s := range strs {
		if l := len(s); l < length {
			length = l
		}
	}

	for i := 0; i < length; i++ {
		mark := strs[0][i]
		flag := true
		for _, s := range strs {
			if s[i] != mark {
				flag = false
				break
			}
		}
		if flag {
			ans = append(ans, mark)
		} else {
			break
		}
	}
	return string(ans)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	max := strs[0]
	min := strs[0]

	for _, s := range strs {
		if s > max {
			max = s
		}
		if s < min {
			min = s
		}
	}

	length := len(max)
	if l := len(min); l < length {
		length = l
	}

	i := 0
	for i = 0; i < length; i++ {
		if max[i] != min[i] {
			break
		}
	}
	return max[:i]
}

func main() {
	s := []string{"acc", "acc"}
	ans := longestCommonPrefix(s)
	fmt.Println(ans)
}
