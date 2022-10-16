package main

func addStrings(num1 string, num2 string) string {
	var ans []byte
	i, j := len(num1)-1, len(num2)-1
	if i > j {
		ans = make([]byte, i+1)
	} else {
		ans = make([]byte, j+1)
	}
	flag := 0
	idx := len(ans) - 1
	for i >= 0 || j >= 0 {
		t := flag
		if i >= 0 {
			t += int(num1[i] - '0')
		}
		if j >= 0 {
			t += int(num2[j] - '0')
		}
		flag = t / 10
		t %= 10
		ans[idx] = byte(t + '0')
		idx--
		i--
		j--
	}
	if flag > 0 {
		ans = append([]byte{'1'}, ans...)
	}
	return string(ans)
}
