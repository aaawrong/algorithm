package dynamic_programming

// 目标和
// 力扣：
// https://leetcode.cn/problems/coin-change-ii/description/
// 随想录：
// https://programmercarl.com/0518.%E9%9B%B6%E9%92%B1%E5%85%91%E6%8D%A2II.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC

// 时间复杂度: O(mn)，其中 m 是amount，n 是 coins 的长度
// 空间复杂度：O(m)
func change(amount int, coins []int) int {
	dp := make([]int, amount+1) // dp表示装满背包j有dp[j]中方法
	dp[0] = 1

	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] += dp[j-coin]
		}
	}

	return dp[amount]
}
