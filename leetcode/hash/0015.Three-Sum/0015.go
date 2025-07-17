package hash

import "sort"

// 三数之和
// 力扣：
// https://leetcode.cn/problems/3sum
// 随想录：
// https://programmercarl.com/0015.%E4%B8%89%E6%95%B0%E4%B9%8B%E5%92%8C.html#%E6%80%9D%E8%B7%AF

// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		if n1 > 0 { // 因为是排好序的，第一个数不为0，则直接返回
			break
		}
		if i > 0 && n1 == nums[i-1] { // a去重
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			n2, n3 := nums[left], nums[right]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				for left < right && nums[left] == n2 { // 继续往后找到不重复的
					left++
				}
				for left < right && nums[right] == n3 { // 继续往前找到不重复的
					right--
				}
			} else if n1+n2+n3 < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}
