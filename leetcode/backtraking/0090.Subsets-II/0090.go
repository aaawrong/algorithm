package backtraking

import "sort"

// 组合
// 力扣：
// https://leetcode.cn/problems/subsets-ii/
// 参考：
// https://programmercarl.com/0090.%E5%AD%90%E9%9B%86II.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC

// 时间复杂度: O(n * 2^n)
// 空间复杂度: O(n)
func subsetsWithDup(nums []int) [][]int {
	path, result := []int{}, [][]int{}
	sort.Ints(nums)
	dfs(nums, 0, path, &result)
	return result
}

func dfs(nums []int, startIndex int, path []int, result *[][]int) {
	tmp := make([]int, len(path))
	copy(tmp, path)
	*result = append(*result, tmp)

	for i := startIndex; i < len(nums); i++ {
		if i > startIndex && nums[i] == nums[i-1] {
			continue
		}

		path = append(path, nums[i])
		dfs(nums, i+1, path, result)
		path = path[:len(path)-1]
	}
}
