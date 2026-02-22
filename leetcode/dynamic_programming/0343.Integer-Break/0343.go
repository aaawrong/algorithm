package dynamic_programming

// 整数拆分
// 力扣：
// https://leetcode.cn/problems/integer-break/description/
// 随想录：
// https://programmercarl.com/0343.%E6%95%B4%E6%95%B0%E6%8B%86%E5%88%86.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC

// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
func integerBreak(n int) int {
	dp := make([]int, n+1) // 把整数 i 拆分成 ≥2 个正整数后，能得到的最大乘积
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			// dp[i] = max(dp[i], j * (i - j), j*dp[i-j])
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j])) // dp表示继续拆
		}
	}

	return dp[n]
}
