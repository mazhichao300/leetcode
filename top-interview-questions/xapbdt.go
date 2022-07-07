package main

func reverseString(s []byte) {
	size := len(s)
	for i := 0; i < size/2; i++ {
		tail := size - 1 - i
		s[i], s[tail] = s[tail], s[i]
	}
}
