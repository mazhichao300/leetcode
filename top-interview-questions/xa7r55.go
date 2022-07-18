package main

type MinStack struct {
	arr    []int
	min    []int
	length int
}

func Constructor() MinStack {
	return MinStack{
		arr:    []int{},
		min:    []int{},
		length: 0,
	}
}

func (this *MinStack) Push(val int) {
	this.arr = append(this.arr, val)
	min := val
	if this.length > 0 && this.min[this.length-1] < min {
		min = this.min[this.length-1]
	}
	this.min = append(this.min, min)
	this.length++
}

func (this *MinStack) Pop() {
	this.arr = this.arr[:this.length-1]
	this.min = this.min[:this.length-1]
	this.length--
}

func (this *MinStack) Top() int {
	return this.arr[this.length-1]
}

func (this *MinStack) GetMin() int {
	return this.min[this.length-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
