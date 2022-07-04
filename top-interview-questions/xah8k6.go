package main

import "fmt"

func isAlDi(c byte) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}
	if c >= 'A' && c <= 'Z' {
		return true
	}
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func format(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return c - ('a' - 'A')
	}
	return c
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1

	for i < j {
		for i < j && !isAlDi(s[i]) {
			i++
		}
		for i < j && !isAlDi(s[j]) {
			j--
		}
		if i >= j {
			break
		}
		if format(s[i]) != format(s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	ans := isPalindrome(s)
	fmt.Println(ans)
}
