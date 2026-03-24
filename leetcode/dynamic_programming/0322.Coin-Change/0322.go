package dynamic_programming

// 零钱兑换
// 力扣:
// https://leetcode.cn/problems/coin-change/description/
// 随想录：
// https://programmercarl.com/0322.%E9%9B%B6%E9%92%B1%E5%85%91%E6%8D%A2.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC

// 时间复杂度: O(mn)，其中 m 是amount，n 是 coins 的长度
// 空间复杂度：O(m)
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1) // dp[j]表示容量为j的最少硬币个数
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1 // 因为要求最小值
	}

	// 先遍历物品再背包（组合），先遍历背包再物品（排列），都可以
	// 完全背包正序，01背包倒序
	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] = min(dp[j-coin]+1, dp[j])
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}
