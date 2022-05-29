package main

import (
	"fmt"
	"strings"
)

func validIPAddress(queryIP string) string {
	if isIpv4(queryIP) {
		return "IPv4"
	}
	if isIpv6(queryIP) {
		return "IPv6"
	}
	return "Neither"
}

func isIpv6(s string) bool {
	a := strings.Split(s, ":")
	if len(a) != 8 {
		return false
	}
	for _, p := range a {
		l := len(p)
		if l < 1 || l > 4 {
			return false
		}
		for _, c := range p {
			if (c < '0' || c > '9') && (c < 'a' || c > 'f') && (c < 'A' || c > 'F') {
				return false
			}
		}
	}
	return true
}

func isIpv4(s string) bool {
	a := strings.Split(s, ".")
	if len(a) != 4 {
		return false
	}

	for _, p := range a {
		l := len(p)
		if l == 0 || l > 3 || (l > 1 && p[0] == '0') {
			return false
		}
		for _, c := range p {
			if c < '0' || c > '9' {
				return false
			}
		}
		if p > "255" {
			return false
		}
	}
	return true
}

func main() {
	ip := "2001:0db8:85a3:0000:0000:8a2e:0370:7334"
	ans := validIPAddress(ip)
	fmt.Println(ans)
}
