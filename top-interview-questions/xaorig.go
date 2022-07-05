package main

import "fmt"

func findWords(board [][]byte, words []string) []string {
	trie := Constructor()
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			dfs(board, i, j, &map[string]bool{}, &[]byte{}, &trie)
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

func dfs(board [][]byte, i, j int, path *map[string]bool, str *[]byte, tree *Trie) {
	key := fmt.Sprintf("%d_%d", i, j)

	(*path)[key] = true
	*str = append(*str, board[i][j])

	points := [][]int{
		[]int{i, j + 1},
		[]int{i, j - 1},
		[]int{i - 1, j},
		[]int{i + 1, j},
	}
	num := 0
	for _, p := range points {
		x, y := p[0], p[1]
		if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
			continue
		}
		k := fmt.Sprintf("%d_%d", x, y)
		if t, ok := (*path)[k]; ok && t {
			continue
		}
		dfs(board, x, y, path, str, tree)
		num++
	}
	// 找不到下一个路径，节点终结，压入trie
	if num == 0 {
		tree.Insert(string(*str))
	}
	(*path)[key] = false
	*str = (*str)[0 : len(*str)-1]
}

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

func main() {

}
