package shuzu

import "sort"

// 双指针
func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	for i, j, k := 0, len(nums)-1, len(nums)-1; i <= j; k-- {
		if nums[i]*nums[i] > nums[j]*nums[j] {
			res[k] = nums[i] * nums[i]
			i++
		} else {
			res[k] = nums[j] * nums[j]
			j--
		}
	}
	return res
}

// 暴力
func sortedSquares2(nums []int) []int {
	for i, value := range nums {
		nums[i] = value * value
	}
	sort.Ints(nums)
	return nums
}
