package dynamic_programming

// 打家劫舍
// 力扣：
// https://leetcode.cn/problems/house-robber/description/
// 随想录：
//

// 时间复杂度: O(n)
// 空间复杂度：O(n)
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n+1) // dp[i]表示偷到第i家最高的金额
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i]) // dp[i-1]不偷i，dp[i-2]+nums[i]偷i
	}

	return dp[n-1]
}
