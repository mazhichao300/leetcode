package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m1 := getMap(nums1, nums2)
	m2 := getMap(nums3, nums4)
	ans := 0
	for k, v := range m1 {
		other := -k
		if n, ok := m2[other]; ok {
			ans += v * n
		}
	}
	return ans
}

func getMap(nums1, nums2 []int) map[int]int {
	ans := map[int]int{}
	for _, i := range nums1 {
		for _, j := range nums2 {
			ans[i+j]++
		}
	}
	return ans
}
