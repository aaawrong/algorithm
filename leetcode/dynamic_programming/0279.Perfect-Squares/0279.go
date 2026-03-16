package dynamic_programming

// 目标和
// 力扣:
// https://leetcode.cn/problems/perfect-squares/description/
// 随想录：
// https://programmercarl.com/0279.%E5%AE%8C%E5%85%A8%E5%B9%B3%E6%96%B9%E6%95%B0.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC

// 时间复杂度: O(mn)，其中 m 是amount，n 是 coins 的长度
// 空间复杂度：O(m)
func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = n + 1
	}

	for i := 1; i*i <= n; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}

	if dp[n] > n {
		return n
	}

	return dp[n]
}
