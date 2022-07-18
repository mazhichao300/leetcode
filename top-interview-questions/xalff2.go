package main

import "fmt"

type MedianFinder struct {
	max  []int
	min  []int
	len1 int
	len2 int
}

func Constructor() MedianFinder {
	return MedianFinder{
		max: make([]int, 100000),
		min: make([]int, 100000),
	}
}

// 从下向上调整
func adjustHeapUp(arr []int, t string, len int, start int) {
	for start > 0 {
		father := (start - 1) / 2
		if t == "max" && arr[start] > arr[father] {
			arr[start], arr[father] = arr[father], arr[start]
		} else if t == "min" && arr[start] < arr[father] {
			arr[start], arr[father] = arr[father], arr[start]
		}
		start = father
	}
}

// 从上向下调整
func adjustHeapDown(arr []int, t string, len int) {
	start := 0
	for {
		left := start*2 + 1
		right := left + 1
		if left > len {
			break
		}
		if right < len && t == "max" && arr[right] > arr[left] {
			left = right
		} else if right < len && t == "min" && arr[right] < arr[left] {
			left = right
		}

		if t == "max" && arr[left] > arr[start] {
			arr[left], arr[start] = arr[start], arr[left]
		} else if t == "min" && arr[left] < arr[start] {
			arr[left], arr[start] = arr[start], arr[left]
		} else {
			break
		}
		start = left
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.len1 == this.len2 {
		this.max[this.len1] = num
		adjustHeapUp(this.max, "max", this.len1+1, this.len1)

		this.min[this.len2] = this.max[0]
		adjustHeapUp(this.min, "min", this.len2+1, this.len2)
		this.len2++

		this.max[0] = this.max[this.len1]
		adjustHeapDown(this.max, "max", this.len1)
	} else {
		this.min[this.len2] = num
		adjustHeapUp(this.min, "min", this.len2+1, this.len2)

		this.max[this.len1] = this.min[0]
		adjustHeapUp(this.max, "max", this.len1+1, this.len1)
		this.len1++

		this.min[0] = this.min[this.len2]
		adjustHeapDown(this.min, "min", this.len2)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	// fmt.Println(this.arr)
	ans := 0.0
	ans += float64(this.min[0])
	if this.len1 == this.len2 {
		ans += float64(this.max[0])
		ans /= 2.0
	}
	return ans
}

func main() {
	o := Constructor()
	t := &o
	t.AddNum(1)
	t.AddNum(2)
	fmt.Println(t.FindMedian())
	t.AddNum(3)
	fmt.Println(t.FindMedian())
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
