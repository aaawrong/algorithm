package dynamic_programming

import "strings"

// 一和零
// 力扣：
// https://leetcode.cn/problems/ones-and-zeroes/description/
// 随想录：
// https://programmercarl.com/0474.%E4%B8%80%E5%92%8C%E9%9B%B6.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC

// 时间复杂度：O(K∗M∗N)，K为strs的长度
// 空间复杂度：O(M∗N)
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for _, s := range strs { // 遍历物品
		zeroCount := strings.Count(s, "0")
		oneCount := len(s) - zeroCount
		if zeroCount > m || oneCount > n {
			continue
		}
		for i := m; i >= zeroCount; i-- { // 遍历背包容量
			for j := n; j >= oneCount; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zeroCount][j-oneCount]+1)
			}
		}
	}

	return dp[m][n]
}
