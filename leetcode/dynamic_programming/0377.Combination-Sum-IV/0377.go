package dynamic_programming

// 组合总和 Ⅳ 和 0518题目相似但是遍历方式不同！！！
// 力扣：
// https://leetcode.cn/problems/combination-sum-iv/description/
// 随想录：
// https://programmercarl.com/0377.%E7%BB%84%E5%90%88%E6%80%BB%E5%92%8C%E2%85%A3.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC

// 时间复杂度: O(mn)
// 空间复杂度：O(m)
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	// 求的是排列，所以要先遍历背包，再遍历物品
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i-num >= 0 {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}
