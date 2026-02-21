package dynamic_programming

// 爬楼梯
// 力扣：
// https://leetcode.cn/problems/climbing-stairs/description/
// 随想录：
// https://programmercarl.com/0070.%E7%88%AC%E6%A5%BC%E6%A2%AF.html#%E6%80%9D%E8%B7%AF

// 时间复杂度：O(n)
// 空间复杂度：O(n)
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
