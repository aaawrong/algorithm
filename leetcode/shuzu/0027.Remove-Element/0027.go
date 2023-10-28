package shuzu

// 双指针
func removeElement(nums []int, val int) int {
	size := len(nums)
	slow := 0
	for fast := 0; fast < size; fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	//nums = nums[:slow]
	return slow
}

// 暴力法
func removeElement2(nums []int, val int) int {
	size := len(nums)
	for i := 0; i < size; i++ {
		if nums[i] == val {
			for j := i + 1; j < size; j++ {
				nums[j-1] = nums[j]
			}
			size--
			i--
		}
	}
	return size
}
