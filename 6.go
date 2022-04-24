package main

import "fmt"

func convertWithSpace(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	ans := make([][]byte, numRows)

	idx := 0
	next := 1
	for i := 0; i < len(s); i++ {
		c := s[i]
		if ans[idx] == nil {
			ans[idx] = []byte{c}
		} else {
			ans[idx] = append(ans[idx], c)
		}
		idx += next
		if idx == numRows-1 {
			next = -1
		}
		if idx == 0 {
			next = 1
		}
	}
	res := ""
	for _, a := range ans {
		res += string(a)
	}
	return res
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	length := len(s)
	// k个一组，分成n组
	k := 2*numRows - 2
	n := length / k
	if length%k != 0 {
		n++
	}
	ans := make([]byte, length)
	index := 0
	// 分别找出第i行的byte
	for i := 0; i < numRows; i++ {
		for j := 0; j < n; j++ {
			start := j * k
			// 第一行和最后一行，每个group中只能选取一个
			if i == 0 || i == numRows-1 {
				if idx := start + i; idx < length {
					//ans = append(ans, s[idx])
					ans[index] = s[idx]
					index++
				}
			} else { // 其他情况，每个group中选取两个
				if idx1 := start + i; idx1 < length {
					//ans = append(ans, s[idx1])
					ans[index] = s[idx1]
					index++
				}
				if idx2 := start + k - i; idx2 < length {
					//ans = append(ans, s[idx2])
					ans[index] = s[idx2]
					index++
				}
			}
		}
	}

	return string(ans)
}

func main() {
	s := "PAYPALISHIRING"
	ans := convert(s, 4)
	fmt.Println(ans)
}
