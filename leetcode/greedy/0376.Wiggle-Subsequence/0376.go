package greedy

// 摆动序列
// 力扣：
// https://leetcode.cn/problems/wiggle-subsequence/
// 随想录：
// https://programmercarl.com/0376.%E6%91%86%E5%8A%A8%E5%BA%8F%E5%88%97.html#%E6%80%9D%E8%B7%AF

// 时间复杂度：O(n)
// 空间复杂度：O(1)

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	res := 1
	preDiff := nums[1] - nums[0]
	if preDiff != 0 {
		res = 2
	}

	for i := 2; i < n; i++ {
		diff := nums[i] - nums[i-1]
		if diff < 0 && preDiff >= 0 || diff > 0 && preDiff <= 0 {
			res++
			preDiff = diff
		}
	}
	return res
}
