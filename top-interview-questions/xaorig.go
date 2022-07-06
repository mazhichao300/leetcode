package main

import (
	"fmt"
	"time"
)

func findWords(board [][]byte, words []string) []string {
	trie := Constructor()
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Println(i, j)
			dfs(&board, i, j, &trie)
		}
	}
	ans := []string{}
	for _, w := range words {
		if trie.StartsWith(w) {
			ans = append(ans, w)
		}
	}
	return ans
}

func dfs(board *[][]byte, i, j int, tree *Trie) {
	if i < 0 || i >= len(*board) || j < 0 || j >= len((*board)[0]) || (*board)[i][j] == '0' {
		return
	}
	for x := 0; x < len(*board); x++ {
		fmt.Println(string((*board)[x]))
	}
	fmt.Println("")
	time.Sleep(time.Second * 1)
	// fmt.Println(i, j, string(board[i][j]))
	ch := (*board)[i][j]
	tree.InsertByte(ch)
	next := tree.Son[ch-'a']

	(*board)[i][j] = '0'

	points := [][]int{
		[]int{i, j + 1},
		[]int{i, j - 1},
		[]int{i - 1, j},
		[]int{i + 1, j},
	}

	for _, p := range points {
		x, y := p[0], p[1]
		dfs(board, x, y, next)
	}
	(*board)[i][j] = ch
}

type Trie struct {
	Son   [26]*Trie
	IsEnd bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) InsertByte(w byte) {
	cur := this
	idx := int(w - 'a')
	if cur.Son[idx] == nil {
		cur.Son[idx] = &Trie{}
	}
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

func main() {
	m := [][]byte{}
	for i := 0; i < 12; i++ {
		y := []byte{}
		for j := 0; j < 12; j++ {
			y = append(y, 'a')
		}
		m = append(m, y)
	}
	s := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	ans := findWords(m, s)
	fmt.Println(ans)
}
