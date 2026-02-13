package greedy

import "sort"

// 根据身高重建队列
// 力扣：
// https://leetcode.cn/problems/queue-reconstruction-by-height/description/
// 随想录：
// https://programmercarl.com/0406.%E6%A0%B9%E6%8D%AE%E8%BA%AB%E9%AB%98%E9%87%8D%E5%BB%BA%E9%98%9F%E5%88%97.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC

// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] { // 身高相同
			return people[i][1] < people[j][1] // 将K按照从小到大排序
		}
		return people[i][0] > people[j][0] // 身高倒序
	})

	for i, p := range people {
		copy(people[p[1]+1:i+1], people[p[1]:i+1]) // 向后移位，题目数据确保队列可以被重建，不需要担心越界问题
		people[p[1]] = p
	}

	return people
}
