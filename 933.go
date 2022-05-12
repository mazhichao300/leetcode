package main

type RecentCounter struct {
	Visit []int
}

func Constructor() RecentCounter {
	return RecentCounter{[]int{}}
}

func (this *RecentCounter) Ping(t int) int {
	ans := 0
	this.Visit = append(this.Visit, t)
	for i := len(this.Visit) - 1; i >= 0; i-- {
		if t-this.Visit[i] <= 3000 {
			ans++
		}
	}
	return ans
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
