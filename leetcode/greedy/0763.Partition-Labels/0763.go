package greedy

// 划分字母区间
// 力扣：
// https://leetcode.cn/problems/partition-labels/description/
// 随想录：
// https://programmercarl.com/0763.%E5%88%92%E5%88%86%E5%AD%97%E6%AF%8D%E5%8C%BA%E9%97%B4.html#%E6%80%9D%E8%B7%AF

// 时间复杂度：O(n)
// 空间复杂度：O(1)，「空间复杂度」通常指辅助空间复杂度，因此这段代码的空间复杂度一般回答 O (1) 即可
func partitionLabels(s string) []int {
	var lastIndex [26]int
	// 存储每个字母最远的索引位置
	for i, ss := range s {
		lastIndex[ss-'a'] = i
	}

	var res []int
	left, right := 0, 0
	for i, ss := range s {
		right = max(right, lastIndex[ss-'a'])
		if i == right {
			res = append(res, right-left+1)
			left = i + 1
		}
	}

	return res
}
