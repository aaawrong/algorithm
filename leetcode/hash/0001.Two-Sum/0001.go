package hash

// 两个数组的交集
// 力扣：
// https://leetcode.cn/problems/two-sum
// 随想录：
//

// 时间复杂度：O(n)
// 空间复杂度：O(n)
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if idx, ok := m[target-num]; ok {
			return []int{idx, i}
		}
		m[num] = i
	}
	return nil
}
