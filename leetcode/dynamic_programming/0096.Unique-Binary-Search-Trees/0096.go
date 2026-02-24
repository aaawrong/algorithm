package dynamic_programming

// 不同的二叉搜索树
// 力扣：
// https://leetcode.cn/problems/unique-binary-search-trees/description/
// 随想录：
// https://programmercarl.com/0096.%E4%B8%8D%E5%90%8C%E7%9A%84%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91.html#%E6%80%9D%E8%B7%AF

// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n]
}
