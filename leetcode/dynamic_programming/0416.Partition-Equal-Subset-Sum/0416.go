package dynamic_programming

// 分割等和子集
// 力扣：
// https://leetcode.cn/problems/partition-equal-subset-sum/description/
// 随想录：
// https://programmercarl.com/0416.%E5%88%86%E5%89%B2%E7%AD%89%E5%92%8C%E5%AD%90%E9%9B%86.html#_01%E8%83%8C%E5%8C%85%E9%97%AE%E9%A2%98

// 时间复杂度：O(N∗Target)
// 空间复杂度：O(Target)
// 一维dp
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// 如果 nums 的总和为奇数则不可能平分成两个子集
	if sum%2 == 1 {
		return false
	}

	target := sum / 2
	dp := make([]int, target+1)

	for _, num := range nums {
		for j := target; j >= num; j-- { // 倒序遍历，每个 num 只会用一次，0/1背包，如果正序会用多次
			dp[j] = max(dp[j], dp[j-num]+num)
		}
	}
	return dp[target] == target
}

// 时间复杂度：O(N∗Target)
// 空间复杂度：O(N∗Target)
// 二维dp
func canPartition2(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}

	target := sum / 2
	dp := make([][]int, len(nums)) // dp[i][j]表示从下标 0 到 i 的元素中任选 在容量 j 下能凑出的最大和
	for i := range dp {
		dp[i] = make([]int, target+1)
	}
	for j := nums[0]; j <= target; j++ {
		// nums 为 [3,4,1]
		// | j     | 0 | 1 | 2 | 3 | 4 | 5 | ... |
		// | dp[0] | 0 | 0 | 0 | 3 | 3 | 3 | ... |
		dp[0][j] = nums[0] // 初始化第一行，背包容量为 j 至少要 大于 nums[0] 才能装
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j <= target; j++ {
			if j < nums[i] {
				dp[i][j] = dp[i-1][j] // 当前容量 j 比 nums[i] 小，放不下这个数， 继承上一行的数据
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i]]+nums[i]) // 取上一行的数据和当前行减去nums[i]的数据+nums[i]的最大值
			}
		}
	}

	return dp[len(nums)-1][target] == target
}
