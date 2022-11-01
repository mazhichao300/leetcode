package main

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	ans := []string{}
	back(s, 0, []string{}, &ans)
	return ans
}

func back(s string, idx int, ip []string, ans *[]string) {
	if idx >= len(s) && len(ip) == 4 {
		*ans = append(*ans, strings.Join(ip, "."))
		return
	}
	if idx >= len(s) || len(ip) > 4 {
		return
	}

	for i := idx; i < len(s); i++ {
		if t, _ := strconv.Atoi(s[idx : i+1]); t <= 255 {
			ip = append(ip, s[idx:i+1])
			back(s, i+1, ip, ans)
			ip = ip[:len(ip)-1]
			if t == 0 {
				break
			}
		} else {
			break
		}
	}
}
