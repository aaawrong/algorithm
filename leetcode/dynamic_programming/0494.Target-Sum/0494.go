package dynamic_programming

import "math"

// 目标和
// 力扣：
// https://leetcode.cn/problems/target-sum/
// 随想录：
// https://programmercarl.com/0494.%E7%9B%AE%E6%A0%87%E5%92%8C.html
// 设某一部分和和为left，另一部分和和为right
// 既然为target，那么就一定有 left组合 - right组合 = target。
// left + right = sum，而sum是固定的。right = sum - left
// left - (sum - left) = target 推导出 left = (target + sum)/2 。
// target是固定的，sum是固定的，left就可以求出来。
// 此时问题就是在集合nums中找出和为left的组合

// 时间复杂度：O(N∗Target)
// 空间复杂度：O(Target)
// 一维dp
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if math.Abs(float64(target)) > float64(sum) { // 无法组成 target
		return 0
	}

	if (sum+target)%2 != 0 { //  left = (target + sum)/2 公式推出必须是偶数
		return 0
	}

	bag := (target + sum) / 2
	dp := make([]int, bag+1) // 凑出和 j 的方案数
	dp[0] = 1
	for _, num := range nums {
		for j := bag; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}

	return dp[bag]
}

// 解法二 DFS TODO 回顾
func findTargetSumWays1(nums []int, S int) int {
	// sums[i] 存储的是后缀和 nums[i:]，即从 i 到结尾的和
	sums := make([]int, len(nums))
	sums[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i > -1; i-- {
		sums[i] = sums[i+1] + nums[i]
	}
	res := 0
	dfsFindTargetSumWays(nums, 0, 0, S, &res, sums)
	return res
}

func dfsFindTargetSumWays(nums []int, index int, curSum int, S int, res *int, sums []int) {
	if index == len(nums) {
		if curSum == S {
			*(res) = *(res) + 1
		}
		return
	}
	// 剪枝优化：如果 sums[index] 值小于剩下需要正数的值，那么右边就算都是 + 号也无能为力了，所以这里可以剪枝了
	if S-curSum > sums[index] {
		return
	}
	dfsFindTargetSumWays(nums, index+1, curSum+nums[index], S, res, sums)
	dfsFindTargetSumWays(nums, index+1, curSum-nums[index], S, res, sums)
}
