package dynamic_programming

// 分割等和子集
// 力扣：
// https://leetcode.cn/problems/last-stone-weight-ii/description/
// 随想录：
// https://programmercarl.com/1049.%E6%9C%80%E5%90%8E%E4%B8%80%E5%9D%97%E7%9F%B3%E5%A4%B4%E7%9A%84%E9%87%8D%E9%87%8FII.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC

// 时间复杂度：O(N∗Target)
// 空间复杂度：O(Target)
// 一维dp
func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, stone := range stones {
		sum += stone
	}

	target := sum / 2           // 取的是下限，所以sum-dp[tar]
	dp := make([]int, target+1) // 在容量为 j 的情况下，当前能够凑出的“最大石头总重量”。
	for _, stone := range stones {
		for j := target; j >= stone; j-- {
			dp[j] = max(dp[j], dp[j-stone]+stone)
		}
	}

	return sum - 2*dp[target]
}
