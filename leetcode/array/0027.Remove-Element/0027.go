package array

// 移除元素
// 力扣：https://leetcode.cn/problems/remove-element/description/
// 随想录：https://programmercarl.com/0027.%E7%A7%BB%E9%99%A4%E5%85%83%E7%B4%A0.html#%E7%AE%97%E6%B3%95%E5%85%AC%E5%BC%80%E8%AF%BE

// 双指针
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func removeElement(nums []int, val int) int {
	slow := 0
	//for fast := 0; fast < len(nums); fast++ {
	//	if nums[fast] != val {
	//		nums[slow] = nums[fast]
	//		slow++
	//	}
	//}
	for _, v := range nums {
		if v != val {
			nums[slow] = v
			slow++
		}
	}
	return slow
}

// 暴力法
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
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
