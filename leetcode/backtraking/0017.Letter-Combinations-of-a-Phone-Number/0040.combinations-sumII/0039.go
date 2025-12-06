package backtraking

import "sort"

// 组合
// 力扣：
// https://leetcode.cn/problems/combination-sum/description/
// 参考：

// 时间复杂度：O(n * 2^n)
// 空间复杂度：O(target) ，递归栈
func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	path, result := []int{}, [][]int{}
	sort.Ints(candidates)
	dfs(candidates, target, 0, path, &result)
	return result
}

func dfs(nums []int, target int, index int, path []int, result *[][]int) {
	if target == 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
		return
	}

	for i := index; i < len(nums); i++ {
		if nums[i] > target {
			break
		}

		path = append(path, nums[i])
		dfs(nums, target-nums[i], i, path, result) // 注意这里迭代的时候 index 依旧不变，因为一个元素可以取多次
		path = path[:len(path)-1]
	}
}
