package main

import "fmt"

type Stack struct {
	Val []byte
	Len int
}

func NewStack() *Stack {
	return &Stack{
		Val: make([]byte, 10005),
		Len: 0,
	}
}

func (s *Stack) Top() byte {
	if s.Len == 0 {
		return ' '
	}
	return s.Val[s.Len-1]
}

func (s *Stack) Pop() byte {
	if s.Len == 0 {
		return ' '
	}
	ans := s.Val[s.Len-1]
	s.Len--
	return ans
}

func (s *Stack) Push(a byte) {
	s.Val[s.Len] = a
	s.Len++
}

func (s *Stack) Length() int {
	return s.Len
}

func removeOuterParentheses(s string) string {
	cnt := 0
	ans := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			if cnt > 0 {
				ans = append(ans, s[i])
			}
			cnt++
		} else {
			cnt--
			if cnt != 0 {
				ans = append(ans, s[i])
			}
		}
	}
	return string(ans)
}

func main() {
	ans := removeOuterParentheses("(()())(())(()(()))")
	fmt.Println(ans)
}
