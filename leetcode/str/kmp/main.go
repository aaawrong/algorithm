package main

// https://blog.csdn.net/yearn520/article/details/6729426

// getNext 构造前缀表next
//
//	next 前缀表数组
//	s 模式串，i表示前缀末尾，j表示后缀末尾
//
// 比如 aacabf 一次匹配a, aa, aac, abca,...,
// 那么 					ji,  ji,
// next                 01, 010, 0001,...
//最坏情况：
//在最坏的情况下，主串 S 和模式串 T 的字符几乎完全相同，但只有最后一个字符不匹配。例如，主串为 "aaaaaa"，模式串为 "aaaaab"。
//在这种情况下，算法会遍历主串的每一个字符，并在每次比较时重置模式串索引 j，导致主串的每个字符都被比较多次。
//时间复杂度为 O(n⋅m)，其中 n 是主串的长度，m 是模式串的长度。
//最好情况：
//在最好情况下，主串的第一个字符就与模式串的第一个字符匹配，且模式串的所有字符都能在主串的开头找到。
//在这种情况下，算法只需进行一次比较，时间复杂度为 O(n)。
//平均情况：
//平均情况下，算法的性能取决于主串和模式串的字符分布。一般情况下，假设字符是随机分布的，时间复杂度大约为O(n+m)，因为每个字符大约会被比较一次。

// 空间复杂度O(n)
func getNext(next []int, s string) {
	j := 0
	next[0] = j
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1] // next[j-1]表示当前最大相等前后缀长，见上面blog
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
}
func strStr(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}
	j := 0
	next := make([]int, n)
	getNext(next, needle)
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1] // 回退到j的前一位
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}
	}
	return -1
}

// ---------- 测试 ----------
//func main() {
//	text := "ABABDABACDABABCABAB"
//	pattern := "ABABCABAB"
//	idx := kmpSearch(text, pattern)
//	fmt.Printf("Pattern found at index: %d\n", idx) // 输出 10
//}
