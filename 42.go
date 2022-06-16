package main

import "fmt"

func trap(height []int) int {
	left := make([]int, len(height))
	right := make([]int, len(height))
	for i := 0; i < len(height); i++ {
		if i == 0 {
			left[i] = height[i]
		} else if height[i] > left[i-1] {
			left[i] = height[i]
		} else {
			left[i] = left[i-1]
		}
	}
	ans := 0
	for i := len(height) - 1; i >= 0; i-- {
		if i == len(height)-1 {
			right[i] = height[i]
		} else if height[i] > right[i+1] {
			right[i] = height[i]
		} else {
			right[i] = right[i+1]
		}
		tmp := right[i]
		if tmp > left[i] {
			tmp = left[i]
		}
		ans += tmp - height[i]
	}
	return ans
}

// n平方的解法
func trapn2(height []int) int {
	ans := 0
	for i := 1; i < len(height)-1; i++ {
		left := 0
		right := 0
		for x := i - 1; x >= 0; x-- {
			if height[x] > left {
				left = height[x]
			}
		}
		for x := i + 1; x < len(height); x++ {
			if height[x] > right {
				right = height[x]
			}
		}
		if right < left {
			left = right
		}
		if left > height[i] {
			ans += left - height[i]
		}
	}
	return ans
}

func main() {
	input := []int{4, 2, 0, 3, 2, 5}
	ans := trap(input)
	fmt.Println(ans)
}
