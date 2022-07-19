package main

type Node struct {
	val int
	cnt int
}

func adjustHeapUp(arr []*Node, start int) {
	for start > 0 {
		father := (start - 1) / 2
		if arr[start].cnt < arr[father].cnt {
			arr[start], arr[father] = arr[father], arr[start]
			start = father
		} else {
			break
		}
	}
}

func adjustHeapDown(arr []*Node, len int) {
	root := 0
	for {
		left, right := 2*root+1, 2*root+2
		if left >= len {
			break
		}
		if right < len && arr[right].cnt < arr[left].cnt {
			left = right
		}
		if arr[left].cnt < arr[root].cnt {
			arr[left], arr[root] = arr[root], arr[left]
		} else {
			break
		}
		root = left
	}
}

func topKFrequent(nums []int, k int) []int {
	count := map[int]*Node{}
	for _, v := range nums {
		if count[v] != nil {
			count[v].cnt++
		} else {
			count[v] = &Node{
				v, 1,
			}
		}
	}

	idx := 0
	arr := make([]*Node, k)
	for _, n := range count {
		if idx < k {
			arr[idx] = n
			adjustHeapUp(arr, idx)
			idx++
		} else if n.cnt > arr[0].cnt {
			arr[0] = n
			adjustHeapDown(arr, k)
		}
	}

	ans := make([]int, k)
	for i, n := range arr {
		ans[i] = n.val
	}
	return ans
}
