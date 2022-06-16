package main

import (
	"fmt"
)

func numUniqueEmails(emails []string) int {
	ans := make(map[string]bool)
	for _, s := range emails {
		t := []rune{}
		host := false
		skip := false
		for _, c := range s {
			if c == '@' || host {
				host = true
				t = append(t, c)
			} else if c == '.' || skip {
				continue
			} else if c == '+' {
				skip = true
			} else {
				t = append(t, c)
			}
		}
		ans[string(t)] = true
	}
	return len(ans)
}

func main() {
	ans := numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"})
	fmt.Println(ans)
}
