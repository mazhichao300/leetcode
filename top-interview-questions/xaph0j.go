package main

func firstUniqChar(s string) int {
	m := [26][2]int{}

	ans := len(s)

	for i := len(s) - 1; i >= 0; i-- {
		idx := int(s[i] - 'a')
		m[idx][0]++
		m[idx][1] = i
	}

	for _, a := range m {
		if a[0] == 1 && a[1] < ans {
			ans = a[1]
		}
	}
	if ans == len(s) {
		ans = -1
	}
	return ans
}
