package hash

// 两个数组的交集
// 力扣：
// https://leetcode.cn/problems/happy-number
// 随想录：
// https://programmercarl.com/0001.%E4%B8%A4%E6%95%B0%E4%B9%8B%E5%92%8C.html

// 时间复杂度：O(n)
// 空间复杂度：O(n)
func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]struct{})
	res := make([]int, 0)
	for _, n := range nums1 {
		m[n] = struct{}{}
	}
	for _, n := range nums2 {
		if _, ok := m[n]; ok {
			delete(m, n)
			res = append(res, n)
		}
	}
	return res
}
