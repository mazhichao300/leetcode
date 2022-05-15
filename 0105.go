package main

import "fmt"

var m map[string]int

func oneEditAway(first string, second string) bool {
	m = make(map[string]int)
	ans := dp(first, second, len(first)-1, len(second)-1)
	// fmt.Println(ans)
	return ans <= 1
}

func dp(first, second string, i, j int) int {
	key := fmt.Sprintf("%d_%d", i, j)
	if ans, ok := m[key]; ok {
		return ans
	}
	// fmt.Println(i, j)
	// if i == 0 && j == 0 {
	// 	if first[i] == second[j] {
	// 		return 0
	// 	} else {
	// 		return 1
	// 	}
	// }
	if i < 0 {
		return j + 1
	}
	if j < 0 {
		return i + 1
	}

	ans := dp(first, second, i-1, j-1)
	if first[i] == second[j] {
		m[key] = ans
		return ans
	}
	if tmp := dp(first, second, i-1, j); tmp < ans {
		ans = tmp
	}
	if tmp := dp(first, second, i, j-1); tmp < ans {
		ans = tmp
	}
	m[key] = ans + 1
	return ans + 1
}

func main() {
	first := ""
	second := ""

	ans := oneEditAway(first, second)
	fmt.Println(ans)
}
