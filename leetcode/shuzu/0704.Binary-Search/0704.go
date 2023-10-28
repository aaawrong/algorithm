package shuzu

// [left,right]
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
func search704_2(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := left + (right-left)/2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return -1
}
