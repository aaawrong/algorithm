package stackAndQueue

// 有效的括号
// 力扣：
// https://leetcode.cn/problems/valid-parentheses/description/
// 随想录：
//https://programmercarl.com/0020.%E6%9C%89%E6%95%88%E7%9A%84%E6%8B%AC%E5%8F%B7.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC

// 时间复杂度：O(n)
// 空间复杂度：O(n)
func isValid1(s string) bool {
	m := make(map[rune]rune)
	m[')'] = '('
	m['}'] = '{'
	m[']'] = '['

	stack := make([]rune, 0)
	for _, ss := range s {
		if ss == '(' || ss == '{' || ss == '[' {
			stack = append(stack, ss)
		} else {
			if len(stack) == 0 {
				return false
			}

			// 取栈顶
			peek := stack[len(stack)-1]
			if peek != m[ss] {
				return false
			}

			// 出栈
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
