package main

//https://leetcode.cn/leetbook/read/top-interview-questions/xm0u83/
func singleNumber(nums []int) int {
	ans := 0
	for _, i := range nums {
		ans ^= i
	}
	return ans
}

func main() {

}
