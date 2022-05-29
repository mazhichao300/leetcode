package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func findTheWinner(n int, k int) int {
	// root := &Node{}
	// p := root
	// for i
	list := make([]int, n)
	idx := 0
	fail := 0
	count := 0
	for fail != n-1 {
		count++
		if count == k {
			list[idx] = -1
			count = 0
			fail++
		}
		idx = getNext(list, idx)
	}
	return idx + 1
}

func getNext(list []int, idx int) int {
	idx++
	if idx >= len(list) {
		idx = 0
	}
	for list[idx] != 0 {
		idx++
		if idx >= len(list) {
			idx = 0
		}
	}
	return idx
}

func main() {
	fmt.Println(findTheWinner(6, 5))
}
