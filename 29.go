package main

import "fmt"

func divide(dividend int, divisor int) int {
	flag := 1
	if dividend == -2147483648 && divisor == -1 {
		return 2147483647
	} else if dividend > 0 && divisor < 0 {
		flag = -1
		divisor = 0 - divisor
	} else if dividend < 0 && divisor > 0 {
		flag = -1
		dividend = 0 - dividend
	} else if dividend < 0 && divisor < 0 {
		dividend = 0 - dividend
		divisor = 0 - divisor
	}
	ans := 0

	d := divisor
	cnt := 1
	for dividend >= divisor {
		if dividend >= d {
			dividend -= d
			ans += cnt
			d = d << 1
			cnt = cnt << 1
		} else {
			d = d >> 1
			cnt = cnt >> 1
		}
	}

	if flag == -1 {
		ans = 0 - ans
	}

	return ans
}

func main() {
	ans := divide(1, 1)
	fmt.Println(ans)
}
