package main

func reverseWords(s string) string {
	arr := []byte{}
	start := -1

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' && start != -1 {
			reverseByte(arr, start, len(arr)-1)

			start = -1
		} else if s[i] != ' ' {
			if start == -1 {
				l := len(arr)
				if l > 0 {
					arr = append(arr, ' ')
					l++
				}
				start = l
			}
			arr = append(arr, s[i])
		}
	}
	if start != -1 {
		reverseByte(arr, start, len(arr)-1)
	}

	reverseByte(arr, 0, len(arr)-1)

	return string(arr)
}

func reverseByte(s []byte, start, end int) {
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}
