package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	ans := 0.0
	lena := len(nums1)
	lenb := len(nums2)
	all := lena + lenb
	k := []int{}
	if all%2 == 0 {
		k = []int{all / 2, all/2 + 1}
	} else {
		k = []int{(all + 1) / 2}
	}

	for _, i := range k {
		ans += float64(find(nums1, nums2, i))
	}
	if all%2 == 0 {
		ans /= 2.0
	}
	return ans
}

func find(nums1 []int, nums2 []int, k int) int {
	lena := len(nums1)
	lenb := len(nums2)
	if lena == 0 {
		return nums2[k-1]
	}
	if lenb == 0 {
		return nums1[k-1]
	}
	if k == 1 {
		if nums1[0] < nums2[0] {
			return nums1[0]
		}
		return nums2[0]
	}

	t := k/2 - 1
	val1 := nums1[lena-1]
	val2 := nums2[lenb-1]
	if t < lena {
		val1 = nums1[t]
	}
	if t < lenb {
		val2 = nums2[t]
	}

	if val1 <= val2 {
		if t >= lena {
			nums1 = []int{}
			k -= lena
		} else {
			nums1 = nums1[t+1:]
			k -= t + 1
		}
	} else {
		if t >= lenb {
			nums2 = []int{}
			k -= lenb
		} else {
			nums2 = nums2[t+1:]
			k -= t + 1
		}
	}

	return find(nums1, nums2, k)
}

func main() {
	a := []int{2, 2, 4, 4}
	b := []int{2, 2, 4, 4}
	ans := findMedianSortedArrays(a, b)
	fmt.Println(ans)
}
