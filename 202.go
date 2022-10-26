package main

func isHappy(n int) bool {
	path := make(map[int]bool, 0)
	for n != 1 {
		next := 0
		for n != 0 {
			single := n % 10
			next += single * single
			n /= 10
		}
		n = next
		if n == 1 {
			return true
		}
		if _, ok := path[n]; ok {
			return false
		}
		path[n] = false
	}
	return true
}
