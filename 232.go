package main

type MyQueue struct {
	StackIn  []int
	StackOut []int
}

func Constructor() MyQueue {
	return MyQueue{[]int{}, []int{}}
}

func (this *MyQueue) Push(x int) {
	this.StackIn = append(this.StackIn, x)
}

func (this *MyQueue) Pop() int {
	p := this.Peek()
	this.StackOut = this.StackOut[:len(this.StackOut)-1]
	return p
}

func (this *MyQueue) Peek() int {
	if len(this.StackOut) == 0 {
		for i := len(this.StackIn) - 1; i >= 0; i-- {
			this.StackOut = append(this.StackOut, this.StackIn[i])
		}
		this.StackIn = []int{}
	}
	return this.StackOut[len(this.StackOut)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.StackIn) == 0 && len(this.StackOut) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
