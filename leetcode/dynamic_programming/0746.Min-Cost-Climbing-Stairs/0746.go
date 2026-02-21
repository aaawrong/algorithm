package greedy

// 使用最小花费爬楼梯
// 力扣：
// https://leetcode.cn/problems/min-cost-climbing-stairs/
// 随想录：
// https://programmercarl.com/0746.%E4%BD%BF%E7%94%A8%E6%9C%80%E5%B0%8F%E8%8A%B1%E8%B4%B9%E7%88%AC%E6%A5%BC%E6%A2%AF.html#%E7%AE%97%E6%B3%95%E5%85%AC%E5%BC%80%E8%AF%BE

// 时间复杂度：O(n)
// 空间复杂度：O(n)
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1) // 到达i需要的最小花费
	dp[0], dp[1] = 0, 0
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[len(cost)]
}
