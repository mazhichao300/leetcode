package main

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	return buildTree(grid, 0, 0, len(grid))
}

func buildTree(grid [][]int, x, y, length int) *Node {
	node := &Node{
		IsLeaf: true,
	}
	val := grid[x][y]
	for i := x; i < x+length; i++ {
		for j := y; j < y+length; j++ {
			if grid[i][j] != val {
				node.IsLeaf = false
				break
			}
		}
	}

	if val == 1 {
		node.Val = true
	}

	// leaf
	if node.IsLeaf {
		return node
	}

	// not leaf
	length /= 2
	node.TopLeft = buildTree(grid, x, y, length)
	node.TopRight = buildTree(grid, x, y+length, length)
	node.BottomLeft = buildTree(grid, x+length, y, length)
	node.BottomRight = buildTree(grid, x+length, y+length, length)
	return node
}

func main() {

}
