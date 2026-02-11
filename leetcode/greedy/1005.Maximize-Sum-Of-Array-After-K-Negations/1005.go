package greedy

import (
	"math"
	"sort"
)

// K次取反后最大化的数组和
// 力扣：
// https://leetcode.cn/problems/maximize-sum-of-array-after-k-negations/description/
// 随想录：
// https://programmercarl.com/1005.K%E6%AC%A1%E5%8F%96%E5%8F%8D%E5%90%8E%E6%9C%80%E5%A4%A7%E5%8C%96%E7%9A%84%E6%95%B0%E7%BB%84%E5%92%8C.html#%E6%80%9D%E8%B7%AF

// 时间复杂度：O(nlogn)
// 空间复杂度：O(1)
func largestSumAfterKNegations(nums []int, k int) int {
	// 按照绝对值从大到小排序，因为要先翻转大负数
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})

	for i := range nums {
		if k > 0 && nums[i] < 0 {
			nums[i] = -nums[i]
			k--
		}
	}
	if k%2 == 1 { // 如果k剩余次数是奇数，取最小的翻转
		nums[len(nums)-1] = -nums[len(nums)-1]
	}

	res := 0
	for _, num := range nums {
		res += num
	}

	return res
}
