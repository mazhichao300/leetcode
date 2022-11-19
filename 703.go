type KthLargest struct {
	Heap     []int
	Capacity int
	Size     int
}

func adjustDown(idx int, arr []int) {
	l := len(arr)
	for {
		son := 2*idx + 1
		if son >= l {
			break
		}
		if right := 2*idx + 2; right < l && arr[right] < arr[son] {
			son = right
		}
		if arr[idx] <= arr[son] {
			break
		}
		arr[idx], arr[son] = arr[son], arr[idx]
		idx = son
	}
}

func makeHeap(arr []int) {
	for i := len(arr) / 2; i >= 0; i-- {
		adjustDown(i, arr)
	}
}

func Constructor(k int, nums []int) KthLargest {
	heap := make([]int, k)
	size := 0
	for i := 0; i < len(nums); i++ {
		if i < k {
			size++
			heap[i] = nums[i]
			if size == k {
				makeHeap(heap)
			}
		} else {
			if nums[i] > heap[0] {
				heap[0] = nums[i]
				adjustDown(0, heap)
			}
		}

	}
	return KthLargest{
		Heap:     heap,
		Capacity: k,
		Size:     size,
	}
}

func (this *KthLargest) Add(val int) int {
	if this.Size < this.Capacity {
		this.Heap[this.Size] = val
		makeHeap(this.Heap)
		this.Size++
	} else if val > this.Heap[0] {
		this.Heap[0] = val
		adjustDown(0, this.Heap)
	}
	return this.Heap[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
