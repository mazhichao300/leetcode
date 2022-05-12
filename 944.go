package main

func minDeletionSize(strs []string) int {
	ans := 0
	col := len(strs[0])
	for j := 0; j < col; j++ {
		for i := 1; i < len(strs); i++ {
			if strs[i][j] < strs[i-1][j] {
				ans++
				break
			}
		}
	}
	return ans
}

func main() {

}
