package greedy

// 斐波那契数列
// 力扣：
// https://leetcode.cn/problems/fibonacci-number/
// 随想录：
// https://programmercarl.com/0509.%E6%96%90%E6%B3%A2%E9%82%A3%E5%A5%91%E6%95%B0.html#%E7%AE%97%E6%B3%95%E5%85%AC%E5%BC%80%E8%AF%BE

// 时间复杂度：O(n)
// 空间复杂度：O(1)
// 优化版动态规划
func fib(n int) int {
	if n <= 1 {
		return n
	}

	cur, prev1, prev2 := 0, 0, 1
	for i := 2; i <= n; i++ {
		cur = prev1 + prev2
		prev1 = prev2
		prev2 = cur
	}

	return cur
}

// 时间复杂度：O(n)
// 空间复杂度：O(n)
// 公式归纳
func fib1(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
