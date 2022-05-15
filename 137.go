package main

import "fmt"

func singleNumber(nums []int) int {
	arr := make([]int, 64)
	for _, x := range nums {
		for i := 0; i < 64; i++ {
			arr[i] += x & 1
			x = x >> 1
		}
	}
	ans := 0
	for i := 63; i >= 0; i-- {
		ans = ans << 1
		ans += arr[i] % 3
	}
	return ans
}

func main() {
	input := []int{2, 2, 3, 2}
	ans := singleNumber(input)
	fmt.Println(ans)
}
