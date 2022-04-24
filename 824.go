package main

import "fmt"

func toGoatLatin(sentence string) string {
	yuan := map[rune]int{
		'a': 1,
		'o': 1,
		'e': 1,
		'i': 1,
		'u': 1,
		'A': 1,
		'O': 1,
		'E': 1,
		'I': 1,
		'U': 1,
	}
	ans := []rune{}
	idx := 1

	word := []rune{}
	var tail []rune
	start := true

	sentence = sentence + " "

	for _, c := range sentence {
		// 一个单词结束，进行转换处理
		if c == ' ' {
			ans = append(ans, word...)
			ans = append(ans, tail...)
			for i := 0; i < idx; i++ {
				ans = append(ans, 'a')
			}
			ans = append(ans, ' ')
			idx++
			tail = nil
			word = []rune{}
			start = true
		} else {
			// 单词
			if start {
				if _, ok := yuan[c]; ok {
					word = append(word, c)
					tail = []rune{'m', 'a'}
				} else {
					tail = []rune{c, 'm', 'a'}
				}
				start = false
			} else {
				word = append(word, c)
			}
		}

	}

	return string(ans[:len(ans)-1])
}

func main() {
	ans := toGoatLatin("I speak Goat Latin")
	fmt.Println(len(ans))
	fmt.Println(ans)
}
