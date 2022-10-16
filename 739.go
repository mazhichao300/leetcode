package main

func dailyTemperatures(temperatures []int) []int {
	l := len(temperatures)
	maxIdx, max := l-1, temperatures[l-1]
	ans := make([]int, len(temperatures))
	for i := len(temperatures) - 2; i >= 0; i-- {
		if temperatures[i] < max {
			ans[i] = maxIdx - i
		}
	}
}
