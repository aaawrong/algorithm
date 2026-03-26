package dynamic_programming

// 打家劫舍II
// 力扣：
// https://leetcode.cn/problems/house-robber-ii/description/
// 随想录：
//

// 时间复杂度: O(n)
// 空间复杂度：O(n)
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	return max(robWithoutCircle(nums[:len(nums)-1]), robWithoutCircle(nums[1:]))
}

func robWithoutCircle(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[n-1]
}
