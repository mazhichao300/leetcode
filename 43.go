package main

import "fmt"

func multiply(num1 string, num2 string) string {
	l1 := len(num1)
	l2 := len(num2)
	ans := make([]int, l1+l2)

	for i := l1 - 1; i >= 0; i-- {
		n1 := int(num1[i] - '0')
		for j := l2 - 1; j >= 0; j-- {
			n2 := int(num2[j] - '0')
			v := n1 * n2
			h := v / 10
			l := v % 10

			ans[i+j+1] += l
			ans[i+j] += h + ans[i+j+1]/10
			ans[i+j+1] %= 10
		}

	}
	t := make([]byte, l1+l2)
	idx := 0
	start := false
	for i := 0; i < len(ans); i++ {
		if ans[i] == 0 && !start {
			continue
		}
		start = true
		t[idx] = byte(ans[i] + '0')
		idx++
	}
	if start == false {
		t[idx] = '0'
	}
	return string(t)
}

func multiply1(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	ans := []byte{}
	for i := len(num2) - 1; i >= 0; i-- {
		tmp := []byte{}
		for cnt := 0; cnt < len(num2)-1-i; cnt++ {
			tmp = append(tmp, '0')
		}

		flag := 0
		for j := len(num1) - 1; j >= 0; j-- {
			flag = flag + (int(num1[j]-'0') * int(num2[i]-'0'))
			t := flag % 10
			flag /= 10
			tmp = append([]byte{byte(t + '0')}, tmp...)
		}
		if flag > 0 {
			tmp = append([]byte{byte(flag + '0')}, tmp...)
		}
		ans = add(ans, tmp)
	}
	return string(ans)
}

func add(a, b []byte) []byte {
	flag := 0
	ans := []byte{}

	i := len(a) - 1
	j := len(b) - 1

	for i >= 0 || j >= 0 {
		if i >= 0 {
			flag += int(a[i] - '0')
		}
		if j >= 0 {
			flag += int(b[j] - '0')
		}
		tmp := flag % 10
		flag /= 10
		ans = append([]byte{byte(tmp + '0')}, ans...)
		i--
		j--
	}
	if flag == 1 {
		ans = append([]byte{byte(flag + '0')}, ans...)
	}
	return ans
}

func main() {
	ans := multiply("2", "3")
	fmt.Println(ans)
}
