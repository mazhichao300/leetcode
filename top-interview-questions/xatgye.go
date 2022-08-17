package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	Idx int
	Val int
}

type PriorityQueue []*Item

func (this PriorityQueue) Len() int {
	return len(this)
}

func (this PriorityQueue) Less(i, j int) bool {
	return this[i].Val > this[j].Val
}

func (this PriorityQueue) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*this = append(*this, item)
}

func (this *PriorityQueue) Pop() interface{} {
	old := *this
	n := len(old)
	item := old[n-1]
	*this = old[0 : n-1]
	return item
}

func maxSlidingWindow(nums []int, k int) []int {
	pq := make(PriorityQueue, k)
	for i := 0; i < k; i++ {
		pq[i] = &Item{i, nums[i]}
	}
	heap.Init(&pq)
	item := pq[0]

	ans := []int{item.Val}

	for i := k; i < len(nums); i++ {
		heap.Push(&pq, &Item{i, nums[i]})
		for pq[0].Idx < i-k+1 {
			heap.Pop(&pq)
		}

		ans = append(ans, pq[0].Val)
	}
	return ans
}

func main() {
	a := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	ans := maxSlidingWindow(a, k)
	fmt.Println(ans)
}
