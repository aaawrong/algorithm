package stackAndQueue

// 删除字符串中的所有相邻重复项
// 力扣：
// https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string/description/
// 随想录：
// https://programmercarl.com/1047.%E5%88%A0%E9%99%A4%E5%AD%97%E7%AC%A6%E4%B8%B2%E4%B8%AD%E7%9A%84%E6%89%80%E6%9C%89%E7%9B%B8%E9%82%BB%E9%87%8D%E5%A4%8D%E9%A1%B9.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC

// 时间复杂度：O(n)
// 空间复杂度：O(n)
func removeDuplicates(s string) string {
	var stack []rune
	for _, ss := range s {
		if len(stack) > 0 && stack[len(stack)-1] == ss {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ss)
		}
	}
	return string(stack)
}

func removeDuplicates2(s string) string {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
