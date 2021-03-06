package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, s := range tokens {
		l := len(stack)
		n1, n2 := 0, 0
		if l >= 2 {
			n1 = stack[l-1]
			n2 = stack[l-2]
		}
		switch s {
		case "+":
			n2 += n1
			stack[l-2] = n2
			stack = stack[:l-1]
		case "-":
			n2 -= n1
			stack[l-2] = n2
			stack = stack[:l-1]
		case "*":
			n2 *= n1
			stack[l-2] = n2
			stack = stack[:l-1]
		case "/":
			n2 /= n1
			stack[l-2] = n2
			stack = stack[:l-1]
		default:
			n, _ := strconv.Atoi(s)
			stack = append(stack, n)
		}
	}
	// fmt.Println(stack)
	return stack[0]
}

func main() {
	s := []string{"2", "1", "+", "3", "*"}
	ans := evalRPN(s)
	fmt.Println(ans)
}
