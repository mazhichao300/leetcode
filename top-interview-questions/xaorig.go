package main

import (
	"fmt"
)

func findWords(board [][]byte, words []string) []string {
	trie := Constructor()
	for _, w := range words {
		trie.Insert(w)
	}
	ans := []string{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			dfs(&board, i, j, &ans, &[]byte{}, &trie)
		}
	}

	return ans
}

func dfs(board *[][]byte, i, j int, ans *[]string, path *[]byte, tree *Trie) {
	if i < 0 || i >= len(*board) || j < 0 || j >= len((*board)[0]) || (*board)[i][j] == '0' {
		return
	}

	ch := (*board)[i][j]
	idx := int(ch - 'a')
	if tree.Son[idx] == nil {
		return
	}

	next := tree.Son[idx]

	*path = append(*path, ch)

	if next.Cnt > 0 {
		*ans = append(*ans, string(*path))
		next.Cnt--
	}

	(*board)[i][j] = '0'

	points := [][]int{
		[]int{i, j + 1},
		[]int{i, j - 1},
		[]int{i - 1, j},
		[]int{i + 1, j},
	}

	for _, p := range points {
		x, y := p[0], p[1]
		dfs(board, x, y, ans, path, next)
	}
	(*board)[i][j] = ch
	*path = (*path)[0 : len(*path)-1]
}

type Trie struct {
	Son [26]*Trie
	Cnt int
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
	cur.Cnt++
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
	return cur.Cnt > 0
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
