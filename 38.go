package main

import (
	"fmt"
	"strconv"
)

var cache []string

func countAndSay(n int) string {
	if cache == nil {
		cache = make([]string, 31)
	}
	cache[1] = "1"
	for i := 2; i <= 30; i++ {
		cache[i] = countString(cache[i-1])
	}
	return cache[n]
}

func countString(s string) string {
	if s == "" {
		return s
	}
	pre := s[0]
	cnt := 1
	ans := []byte{}
	for i := 1; i < len(s); i++ {
		if s[i] == pre {
			cnt++
		} else {
			cntStr := strconv.Itoa(cnt)
			ans = append(ans, []byte(cntStr)...)
			ans = append(ans, pre)
			cnt = 1
			pre = s[i]
		}
	}
	cntStr := strconv.Itoa(cnt)
	ans = append(ans, []byte(cntStr)...)
	ans = append(ans, pre)
	return string(ans)
}

func main() {
	ans := countAndSay(5)
	fmt.Println(ans)
}
