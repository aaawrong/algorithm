package dynamic_programming

// 不同路径II
// 力扣：
// https://leetcode.cn/problems/unique-paths-ii/description/
// 随想录：
// https://programmercarl.com/0063.%E4%B8%8D%E5%90%8C%E8%B7%AF%E5%BE%84II.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC

// 时间复杂度：O(m*n)
// 空间复杂度：O(m*n)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// 起点为0，直接返回
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 || obstacleGrid[0][0] == 1 {
		return 0
	}

	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 第一列的每一行，如果存在障碍物，赋值0
	for i := 0; i < m && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = 1
	}

	// 第一行的每一列，如果存在障碍物，赋值0
	for j := 0; j < n && obstacleGrid[0][j] == 0; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] != 1 { // 不存在障碍物
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}
