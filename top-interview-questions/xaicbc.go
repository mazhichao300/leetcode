package main

func adjustHeapUp(heap []int, idx int) {
	for {
		left, right, max := idx*2, idx*2+1, idx
		if left < len(heap) && heap[left] > heap[max] {
			max = left
		}
		if right < len(heap) && heap[right] > heap[max] {
			max = right
		}
		if idx != max {
			heap[idx], heap[max] = heap[max], heap[idx]
		}
		if idx == 0 {
			break
		}
		idx /= 2
	}
}

func adjustHeapDown(heap []int) {
	idx := 0
	for {
		left, right, max := idx*2, idx*2+1, idx
		if left >= len(heap) && right >= len(heap) {
			break
		}
		if heap[left] > heap[max] {
			max = left
		}
		if right < len(heap) && heap[right] > heap[max] {
			max = right
		}
		if idx != max {
			heap[idx], heap[max] = heap[max], heap[idx]
		} else {
			break
		}
		idx = max
	}
}

func kthSmallest(matrix [][]int, k int) int {
	heap := []int{}
	n := len(matrix)
	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n && j < k; j++ {
			if cnt < k {
				heap = append(heap, matrix[i][j])
				adjustHeapUp(heap, cnt)
				cnt++
			} else if matrix[i][j] < heap[0] {
				heap[0] = matrix[i][j]
				adjustHeapDown(heap)
			}
		}
	}
	return heap[0]
}
