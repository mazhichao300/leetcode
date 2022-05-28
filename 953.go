package main

func isAlienSorted(words []string, order string) bool {
	alphaOrder := make(map[byte]int)
	for i := 0; i < len(order); i++ {
		alphaOrder[order[i]] = i
	}
	for i := 1; i < len(words); i++ {
		if isBig(words[i-1], words[i], alphaOrder) {
			return false
		}
	}
	return true
}

func isBig(w1, w2 string, order map[byte]int) bool {
	for i := 0; i < len(w1) && i < len(w2); i++ {
		if order[w1[i]] > order[w2[i]] {
			return true
		} else if order[w1[i]] < order[w2[i]] {
			return false
		}
	}
	return len(w1) > len(w2)
}

func main() {

}
