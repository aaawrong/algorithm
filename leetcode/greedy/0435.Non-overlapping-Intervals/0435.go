package greedy

import "sort"

// 无重叠区间
// 力扣：
// https://leetcode.cn/problems/non-overlapping-intervals/description/
// 随想录：
// https://programmercarl.com/0435.%E6%97%A0%E9%87%8D%E5%8F%A0%E5%8C%BA%E9%97%B4.html#%E7%AE%97%E6%B3%95%E5%85%AC%E5%BC%80%E8%AF%BE

// 时间复杂度：O(nlogn)
// 空间复杂度：O(logn) (递归栈空间）

// 找不重叠个数
func eraseOverlapIntervals(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	// 先按照左边界排序
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	res := 1
	for i := 1; i < len(points); i++ {
		if points[i][0] >= points[i-1][1] { // 找出不重叠的个数，题目相邻不算重叠
			res++
		} else { // 修改右边界
			points[i][1] = min(points[i][1], points[i-1][1])
		}
	}

	return len(points) - res
}

// 找重叠个数
func eraseOverlapIntervals2(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	// 先按照左边界排序
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	res := 0
	for i := 1; i < len(points); i++ {
		if points[i][0] < points[i-1][1] { // 找出重叠的个数
			res++
			points[i][1] = min(points[i][1], points[i-1][1])
		}
	}

	return res
}
