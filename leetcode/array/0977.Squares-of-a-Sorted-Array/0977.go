package shuzu

import "sort"

// 有序数组的平方
// 力扣：https://leetcode.cn/problems/squares-of-a-sorted-array/submissions/639775832/
// 随想录：https://programmercarl.com/0977.%E6%9C%89%E5%BA%8F%E6%95%B0%E7%BB%84%E7%9A%84%E5%B9%B3%E6%96%B9.html#%E7%AE%97%E6%B3%95%E5%85%AC%E5%BC%80%E8%AF%BE

// 双指针
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func sortedSquares(nums []int) []int {
	results := make([]int, len(nums))
	for i, j, k := 0, len(nums)-1, len(nums)-1; i <= j; k-- {
		if nums[i]*nums[i] > nums[j]*nums[j] {
			results[k] = nums[i] * nums[i]
			i++
		} else {
			results[k] = nums[j] * nums[j]
			j--
		}
	}
	return results
}

// 暴力
func sortedSquares2(nums []int) []int {
	for i, value := range nums {
		nums[i] = value * value
	}
	sort.Ints(nums)
	return nums
}
