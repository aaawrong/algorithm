package backtraking

import (
	"strconv"
	"strings"
)

// 复原IP地址
// 力扣：
// https://leetcode.cn/problems/restore-ip-addresses/description/
// 参考：
// https://programmercarl.com/0093.%E5%A4%8D%E5%8E%9FIP%E5%9C%B0%E5%9D%80.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC

// 时间复杂度: O(3^4)，IP地址最多包含4个数字，每个数字最多有3种可能的分割方式，则搜索树的最大深度为4，每个节点最多有3个子节点。
// 空间复杂度：O(n)

func restoreIpAddresses(s string) []string {
	path, result := make([]string, 0), make([]string, 0)
	dfs(s, 0, path, &result)
	return result
}
func dfs(s string, start int, path []string, result *[]string) {
	if len(path) == 4 {
		if start == len(s) {
			tmp := strings.Join(path, ".")
			*result = append(*result, tmp)
		}
	}

	for i := start; i < len(s); i++ {
		if i != start && s[start] == '0' { // 前导不为0
			break
		}

		str := s[start : i+1]
		num, _ := strconv.Atoi(str)
		if num < 0 || num > 255 {
			break
		}
		path = append(path, str)
		dfs(s, i+1, path, result)
		path = path[:len(path)-1]
	}
}
