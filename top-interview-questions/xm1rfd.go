package main

func containsDuplicate(nums []int) bool {
	m := map[int]bool{}
	for _, n := range nums {
		if _, ok := m[n]; ok {
			return true
		}
		m[n] = true
	}
	return false
}
