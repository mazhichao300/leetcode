package main

import "fmt"

// 中心扩散法
func longestPalindrome1(s string) string {
	ans := 0
	start := 0
	len := len(s)

	for i := 0; i < len; i++ {
		// 以i节点为中心计算
		left := i - 1
		right := i + 1
		num := 1
		for left >= 0 && right < len && s[left] == s[right] {
			num += 2
			left--
			right++
		}
		if num > ans {
			ans = num
			start = left + 1
		}
		// 以i,i+1为中心计算
		left = i
		right = i + 1
		num = 0
		for left >= 0 && right < len && s[left] == s[right] {
			num += 2
			left--
			right++
		}
		if num > ans {
			ans = num
			start = left + 1
		}
	}
	return s[start : start+ans]
}

type Item struct {
	start int
	end   int
}

var dp [][]*Item

// dp
func longestPalindrome(s string) string {
	length := len(s)
	dp = make([][]*Item, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]*Item, length)
	}
	ans := get(s, 0, length-1)
	return s[ans.start : ans.end+1]
}

func getLen(item *Item) int {
	if item == nil {
		return 0
	}
	return item.end + 1 - item.start
}

func get(s string, start, end int) *Item {
	if dp[start][end] != nil {
		return dp[start][end]
	}
	if start == end {
		dp[start][end] = &Item{start, end}
		return dp[start][end]
	}
	if start+1 == end {
		if s[start] == s[end] {
			dp[start][end] = &Item{start, end}
		} else {
			dp[start][end] = &Item{start, end - 1}
		}
		return dp[start][end]
	}
	if start > end {
		return nil
	}

	ans := get(s, start, end-1)
	ans2 := get(s, start+1, end)
	if getLen(ans2) > getLen(ans) {
		ans = ans2
	}

	if s[start] == s[end] && end-start >= 2 {
		ans3 := get(s, start+1, end-1)
		if getLen(ans3) == end-start-1 {
			ans = &Item{start, end}
		}
	}
	dp[start][end] = ans

	return ans
}

func main() {
	s := "rgczcpratwyqxaszbuwwcadruayhasynuxnakpmsyhxzlnxmdtsqqlmwnbxvmgvllafrpmlfuqpbhjddmhmbcgmlyeypkfpreddyencsdmgxysctpubvgeedhurvizgqxclhpfrvxggrowaynrtuwvvvwnqlowdihtrdzjffrgoeqivnprdnpvfjuhycpfydjcpfcnkpyujljiesmuxhtizzvwhvpqylvcirwqsmpptyhcqybstsfgjadicwzycswwmpluvzqdvnhkcofptqrzgjqtbvbdxylrylinspncrkxclykccbwridpqckstxdjawvziucrswpsfmisqiozworibeycuarcidbljslwbalcemgymnsxfziattdylrulwrybzztoxhevsdnvvljfzzrgcmagshucoalfiuapgzpqgjjgqsmcvtdsvehewrvtkeqwgmatqdpwlayjcxcavjmgpdyklrjcqvxjqbjucfubgmgpkfdxznkhcejscymuildfnuxwmuklntnyycdcscioimenaeohgpbcpogyifcsatfxeslstkjclauqmywacizyapxlgtcchlxkvygzeucwalhvhbwkvbceqajstxzzppcxoanhyfkgwaelsfdeeviqogjpresnoacegfeejyychabkhszcokdxpaqrprwfdahjqkfptwpeykgumyemgkccynxuvbdpjlrbgqtcqulxodurugofuwzudnhgxdrbbxtrvdnlodyhsifvyspejenpdckevzqrexplpcqtwtxlimfrsjumiygqeemhihcxyngsemcolrnlyhqlbqbcestadoxtrdvcgucntjnfavylip"
	ans := longestPalindrome(s)
	fmt.Println(ans)
}
