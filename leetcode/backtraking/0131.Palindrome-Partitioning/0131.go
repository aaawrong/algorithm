package backtraking

// 分割回文串
// 力扣：https://leetcode.cn/problems/palindrome-partitioning/description/
// 参考：https://programmercarl.com/0131.%E5%88%86%E5%89%B2%E5%9B%9E%E6%96%87%E4%B8%B2.html#%E6%80%9D%E8%B7%AF

//时间复杂度: O(n * 2^n)
//空间复杂度: O(n^2)

func partition(s string) [][]string {
	path, result := make([]string, 0), make([][]string, 0)
	dfs(s, 0, path, &result)
	return result
}

func dfs(s string, start int, path []string, result *[][]string) {
	if start == len(s) { // 如果起始位置等于s的大小，说明已经找到了一组分割方案了
		tmp := make([]string, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
		return
	}
	for i := start; i < len(s); i++ {
		str := s[start : i+1]
		if isPalindrome(str) {
			path = append(path, str)
			dfs(s, i+1, path, result) // 寻找i+1为起始位置的子串
			path = path[:len(path)-1] // 回溯过程，弹出本次已经添加的子串
		}
	}
}

// 回文判断
func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
