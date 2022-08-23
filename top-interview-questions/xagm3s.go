package main

import "math/rand"

type RandomizedSet struct {
	m map[int]int
	a []int
	l int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		map[int]int{},
		[]int{},
		0,
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.m[val] = this.l
	this.a = append(this.a, val)
	this.l++
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if idx, ok := this.m[val]; ok {
		this.a[idx] = this.a[this.l-1]
		this.m[this.a[idx]] = idx
		this.a = this.a[:this.l-1]
		this.l--
		delete(this.m, val)
		return true
	} else {
		return false
	}
}

func (this *RandomizedSet) GetRandom() int {
	idx := rand.Intn(this.l)
	return this.a[idx]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
