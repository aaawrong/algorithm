package greedy

import "sort"

// 用最少数量的箭引爆气球
// 力扣：
// https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons/description/
// 随想录：
// https://programmercarl.com/0452.%E7%94%A8%E6%9C%80%E5%B0%91%E6%95%B0%E9%87%8F%E7%9A%84%E7%AE%AD%E5%BC%95%E7%88%86%E6%B0%94%E7%90%83.html#%E6%80%9D%E8%B7%AF

// 时间复杂度：O(nlogn)
// 空间复杂度：O(logn) (递归栈空间）
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	// 先按照左边界排序
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	res := 1
	for i := 1; i < len(points); i++ {
		if points[i][0] > points[i-1][1] { // 当前左边界小于上一个有边界，无交集
			res++
		} else { // 修改右边界
			points[i][1] = min(points[i][1], points[i-1][1])
		}
	}

	return res
}
