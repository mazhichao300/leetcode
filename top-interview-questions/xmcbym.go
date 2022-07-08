package main

func intersect(nums1 []int, nums2 []int) []int {
	m := map[int]int{}

	for _, n := range nums1 {
		m[n]++
	}

	ans := []int{}
	for _, n := range nums2 {
		if m[n] > 0 {
			ans = append(ans, n)
			m[n]--
		}
	}
	return ans
}
