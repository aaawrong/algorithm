package greedy

// 最大子序和
// 力扣：
// https://leetcode.cn/problems/maximum-subarray/
// 随想录：
// https://programmercarl.com/0053.%E6%9C%80%E5%A4%A7%E5%AD%90%E5%BA%8F%E5%92%8C.html#%E7%AE%97%E6%B3%95%E5%85%AC%E5%BC%80%E8%AF%BE

// 时间复杂度：O(n)
// 空间复杂度：O(1)
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	count := 0
	maxC := nums[0]
	for _, n := range nums {
		count += n
		if count > maxC {
			maxC = count
		}
		if count < 0 {
			count = 0
		}
	}

	return maxC
}
