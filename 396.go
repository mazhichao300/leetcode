package main

import "fmt"

func maxRotateFunction(nums []int) int {
	fPre := 0
	length := len(nums)
	sum := 0
	for i, n := range nums {
		sum += n
		fPre += i * n
	}
	ans := fPre
	for start := 1; start < length; start++ {
		lastStart := start - 1
		// 公示推导：f(n)-f(n-1)=length*nums[lastStart] - sum
		f := fPre + length*nums[lastStart] - sum

		if f > ans {
			ans = f
		}
		fPre = f
	}
	return ans
}

func main() {
	ans := maxRotateFunction([]int{18, 4, 13, -1, 2, 7, 19, 14, 11, 0, -9, 9, 4, 2, -2, -7, 18, 18, -7, -5, 9, 6, -8, 3, 13, 0, 15, 9, 10, -1, 17, 13, 13, -8, 3, 8, 4, 19, 10, -8, 18, 15, 11, 13, 11, 1, 14, 2, 10, 1, 11, 5, 18, 5, -8, 13, -10, 5, -8, -9, -5, 9, 10, -10, -3, -3, -4, -4, -8, -10})
	fmt.Println(ans)
}
