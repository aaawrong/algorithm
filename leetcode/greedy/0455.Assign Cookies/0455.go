package array

import "sort"

// 分发饼干
// 力扣：
// https://leetcode.cn/problems/assign-cookies/description/
// 随想录：
// https://programmercarl.com/0455.%E5%88%86%E5%8F%91%E9%A5%BC%E5%B9%B2.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC

// 时间复杂度：O(n log n)
// 空间复杂度：

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	index := 0
	for i := 0; i < len(s); i++ {
		if index < len(g) && g[index] <= s[i] {
			index++
		}
	}
	return index
}
