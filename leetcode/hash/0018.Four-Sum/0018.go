package hash

import "sort"

// 四数之和
// 力扣：
//
// 随想录：
// https://programmercarl.com/0018.%E5%9B%9B%E6%95%B0%E4%B9%8B%E5%92%8C.html#%E6%80%9D%E8%B7%AF

// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-3; i++ {
		n1 := nums[i]
		if i > 0 && n1 == nums[i-1] { // 从第二个开始和前面一位对比去重
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			n2 := nums[j]
			if j > i+1 && n2 == nums[j-1] { // 去重
				continue
			}
			left, right := j+1, len(nums)-1
			for left < right {
				n3, n4 := nums[left], nums[right]
				sum := n1 + n2 + n3 + n4
				if sum == target {
					res = append(res, []int{n1, n2, n3, n4})
					for left < right && nums[left] == n3 {
						left++
					}
					for left < right && nums[right] == n4 {
						right--
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return res
}
