package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	minus := false
	if n < 0 {
		minus = true
		n = -n
	}
	ans := x
	n--
	cnt := 1
	p := x
	for n > 0 {
		if cnt <= n {
			ans *= p
			n -= cnt
			p *= p
			cnt *= 2
		} else {
			p /= x
			cnt--
		}
	}
	if minus {
		ans = 1 / ans
	}
	return ans
}

// func binayPow(x float64, k, n int) float64 {

// }

func main() {
	ans := myPow(2.0, -2)
	fmt.Println(ans)
}
