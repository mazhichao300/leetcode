package main

import "fmt"

type DList struct {
	Key  int
	Val  int
	Pre  *DList
	Next *DList
}

type LRUCache struct {
	Capacity int
	Size     int
	Index    map[int]*DList
	Head     *DList
	Tail     *DList
}

func Constructor(capacity int) LRUCache {
	head := &DList{}
	tail := &DList{}
	head.Next = tail
	tail.Pre = head
	return LRUCache{
		Capacity: capacity,
		Index:    make(map[int]*DList, 0),
		Head:     head, // 首尾放两个实体节点，方便后续统一操作
		Tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if d, ok := this.Index[key]; ok {
		// 当前节点摘出来
		d.Pre.Next = d.Next
		d.Next.Pre = d.Pre
		// 放在队首
		d.Pre = this.Head
		d.Next = this.Head.Next
		this.Head.Next = d
		d.Next.Pre = d
		return d.Val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	// 当前key已经在缓存中，更新缓存
	if this.Get(key) != -1 {
		// 更新内容
		this.Head.Next.Val = value
		return
	}
	// 需要淘汰一个节点
	if this.Size >= this.Capacity {
		d := this.Tail.Pre
		this.Tail.Pre = d.Pre
		d.Pre.Next = this.Tail
		delete(this.Index, d.Key)
	} else { // 增加节点，实际容量加1
		this.Size++
	}
	// 增加一个节点
	d := &DList{
		Key:  key,
		Val:  value,
		Next: this.Head.Next,
		Pre:  this.Head,
	}
	d.Next.Pre = d
	this.Head.Next = d
	this.Index[key] = d
}

func (this *LRUCache) PrintDList() {
	fmt.Println("---------")
	for d := this.Head.Next; d.Next != nil; d = d.Next {
		fmt.Println(this.Size, d.Key, d.Val)
	}
	for d := this.Tail.Pre; d.Pre != nil; d = d.Pre {
		fmt.Println(this.Size, d.Key, d.Val)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1) // 缓存是 {1=1}
	fmt.Println(lRUCache.Index)
	lRUCache.PrintDList()
	lRUCache.Put(2, 2) // 缓存是 {1=1, 2=2}
	fmt.Println(lRUCache.Index)
	lRUCache.PrintDList()
	lRUCache.Get(1) // 返回 1
	lRUCache.PrintDList()
	lRUCache.Put(3, 3) // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	fmt.Println(lRUCache.Index)
	lRUCache.PrintDList()
	lRUCache.Get(2) // 返回 -1 (未找到)
	lRUCache.PrintDList()
	lRUCache.Put(4, 4) // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	fmt.Println(lRUCache.Index)
	lRUCache.PrintDList()
	lRUCache.Get(1) // 返回 -1 (未找到)
	lRUCache.PrintDList()
	lRUCache.Get(3) // 返回 3
	lRUCache.PrintDList()
	lRUCache.Get(4) // 返回 4
	lRUCache.PrintDList()

}
