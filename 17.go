package main

import "fmt"

var tmp []byte
var ans []string
var dict map[byte][]byte

func letterCombinations(digits string) []string {
	if dict == nil {
		dict = map[byte][]byte{
			'2': []byte{'a', 'b', 'c'},
			'3': []byte{'d', 'e', 'f'},
			'4': []byte{'g', 'h', 'i'},
			'5': []byte{'j', 'k', 'l'},
			'6': []byte{'m', 'n', 'o'},
			'7': []byte{'p', 'q', 'r', 's'},
			'8': []byte{'t', 'u', 'v'},
			'9': []byte{'w', 'x', 'y', 'z'},
		}
	}
	tmp = make([]byte, len(digits))
	ans = []string{}
	f(digits, 0)

	return ans
}

func f(digits string, i int) {
	if len(digits) == 0 {
		return
	}
	if i >= len(digits) {
		ans = append(ans, string(tmp))
		return
	}
	d := dict[digits[i]]
	for _, c := range d {
		tmp[i] = c
		f(digits, i+1)
	}
}

func main() {
	a := "2"
	ans := letterCombinations(a)
	fmt.Println(ans)
	ans = letterCombinations("")
	fmt.Println(ans)
}
