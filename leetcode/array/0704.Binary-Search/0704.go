package array

// 二分查找
// leetcode: https://leetcode.cn/problems/binary-search/description/
// 随想录: https://programmercarl.com/0704.%E4%BA%8C%E5%88%86%E6%9F%A5%E6%89%BE.html#%E6%80%9D%E8%B7%AF

// 主要是要根据left和right是左闭右开还是左闭右闭来确定循环里面是<=还是<

// [left,right]
// O(log n), O(1)
func search704(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// [left,right)
// O(log n), O(1)

func search704_2(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
