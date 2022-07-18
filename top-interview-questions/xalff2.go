package main

import "fmt"

type MedianFinder struct {
	arr []int
	l   int
}

func Constructor() MedianFinder {
	return MedianFinder{
		arr: []int{},
		l:   0,
	}
}

func (this *MedianFinder) AddNum(num int) {
	start, end := 0, this.l-1

	for start <= end {
		mid := (start + end) / 2
		if this.arr[mid] < num {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	// fmt.Println(this.arr, num, start)
	prefix := this.arr[0:start]
	tail := this.arr[start:]
	// fmt.Println(prefix,num,tail)
	tmp := append(prefix, num)
	tmp = append(tmp, tail...)
	this.arr = make([]int, len(tmp))
	copy(this.arr, tmp)

	this.l++
}

func (this *MedianFinder) FindMedian() float64 {
	// fmt.Println(this.arr)
	ans := 0.0
	mid := this.l >> 1
	ans += float64(this.arr[mid])
	if this.l&1 == 0 {
		ans += float64(this.arr[mid-1])
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
