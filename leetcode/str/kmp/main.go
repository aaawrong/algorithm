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
