package main

func getVersion(v string) []int {
	ans := []int{}
	n := 0
	for _, c := range v {
		if c != '.' {
			n = n*10 + int(c-'0')
		} else {
			ans = append(ans, n)
			n = 0
		}
	}
	if n > 0 {
		ans = append(ans, n)
	}

	end := len(ans) - 1

	for end > 0 && ans[end] == 0 {
		end--
	}
	return ans[:end+1]
}

func compareVersion(version1 string, version2 string) int {
	v1 := getVersion(version1)
	v2 := getVersion(version2)
	l1 := len(v1)
	l2 := len(v2)
	for i := 0; i < l1 && i < l2; i++ {
		if v1[i] > v2[i] {
			return 1
		} else if v1[i] < v2[i] {
			return -1
		}
	}
	if l1 > l2 {
		return 1
	} else if l1 < l2 {
		return -1
	} else {
		return 0
	}
}

func main() {
	a := compareVersion("7.5.2.4", "7.5.3")
}
