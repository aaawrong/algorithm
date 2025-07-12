package hash

// 两个数组的交集
// 力扣：
// https://leetcode.cn/problems/4sum-ii/
// 随想录：
// https://programmercarl.com/0454.%E5%9B%9B%E6%95%B0%E7%9B%B8%E5%8A%A0II.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC

// 时间复杂度：O(n^2)
// 空间复杂度：O(n^2)
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	count := 0
	m := make(map[int]int, len(nums1)*len(nums2))
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			m[n1+n2]++
		}
	}
	for _, n3 := range nums3 {
		for _, n4 := range nums4 {
			count += m[-n3-n4]
		}
	}
	return count
}
