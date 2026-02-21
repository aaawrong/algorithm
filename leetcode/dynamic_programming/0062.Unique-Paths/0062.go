package dynamic_programming

// 不同路径
// 力扣：
// https://leetcode.cn/problems/unique-paths/description/
// 随想录：
// https://programmercarl.com/0062.%E4%B8%8D%E5%90%8C%E8%B7%AF%E5%BE%84.html#%E7%AE%97%E6%B3%95%E5%85%AC%E5%BC%80%E8%AF%BE

// 时间复杂度：O(m*n)
// 空间复杂度：O(m*n)
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)   // dp表示dp[i][j]有多少种路径
	for i := 0; i < m; i++ { // i 表示行
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ { // j 表示列
			if i == 0 || j == 0 { // 初始化第一行，第一列
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}
