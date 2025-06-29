package array

// 长度最小的子数组
// 力扣：https://leetcode.cn/problems/minimum-size-subarray-sum/description/
// 随想录：https://programmercarl.com/0209.%E9%95%BF%E5%BA%A6%E6%9C%80%E5%B0%8F%E7%9A%84%E5%AD%90%E6%95%B0%E7%BB%84.html

// 双指针
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left, sum, res := 0, 0, len(nums)+1
	for right, val := range nums {
		sum += val
		for sum >= target { // 这里用for是因为如果target=100，那么需要吧前面的都扣减 1,1,1,100
			res = min(res, right-left+1)
			sum -= nums[left]
			left++
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

// 暴力法 双for
