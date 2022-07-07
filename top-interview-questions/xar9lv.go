package main

func isAnagram(s string, t string) bool {
	m := [26]int{}
	for _, c := range s {
		m[c-'a']++
	}
	for _, c := range t {
		m[c-'a']--
	}
	for _, n := range m {
		if n != 0 {
			return false
		}
	}
	return true
}
