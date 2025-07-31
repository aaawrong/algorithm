package str

import "strings"

// 重复的子字符串
// 力扣：
// https://leetcode.cn/problems/repeated-substring-pattern
// 随想录：

// 时间复杂度：O(n) O(n^2)
// 空间复杂度：O(n)
func repeatedSubstringPatternSimple(s string) bool {
	//将字符串 s 拼接两次，得到 s + s。
	//去掉 s + s 的首尾字符（避免匹配原始字符串本身）。
	//如果 s 存在于这个新的字符串中，说明它是由重复子串组成的。
	ss := (s + s)[1 : 2*len(s)-1]
	return strings.Contains(ss, s)
}
