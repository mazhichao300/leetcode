package main

type Node struct {
	Idx int
	Val int
}

func dailyTemperatures1(temperatures []int) []int {
	stack := []*Node{}
	ans := make([]int, len(temperatures))
	for i, v := range temperatures {
		for len(stack) > 0 && stack[len(stack)-1].Val < v {
			n := stack[len(stack)-1]
			ans[n.Idx] = i - n.Idx
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, &Node{i, v})
	}
	return ans
}

func dailyTemperatures2(temperatures []int) []int {
	stack := []int{}
	ans := make([]int, len(temperatures))
	for i, v := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < v {
			n := stack[len(stack)-1]
			ans[n] = i - n
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, len(temperatures))
	ans := make([]int, len(temperatures))
	length := 0
	for i, v := range temperatures {
		for length > 0 && temperatures[stack[length-1]] < v {
			n := stack[length-1]
			ans[n] = i - n
			length--
		}
		stack[length] = i
		length++
	}
	return ans
}
