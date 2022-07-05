package main

import "runtime/debug"

func init() { debug.SetGCPercent(-1) }

type Trie struct {
	Son   [26]*Trie
	IsEnd bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	cur := this
	for _, c := range word {
		idx := int(c - 'a')
		if cur.Son[idx] == nil {
			cur.Son[idx] = &Trie{}
		}
		cur = cur.Son[idx]
	}
	cur.IsEnd = true
}

func (this *Trie) Search(word string) bool {
	cur := this
	for _, c := range word {
		idx := int(c - 'a')
		if cur.Son[idx] == nil {
			return false
		}
		cur = cur.Son[idx]
	}
	return cur.IsEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for _, c := range prefix {
		idx := int(c - 'a')
		if cur.Son[idx] == nil {
			return false
		}
		cur = cur.Son[idx]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
