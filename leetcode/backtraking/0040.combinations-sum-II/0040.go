package backtraking

import "sort"

// 组合
// 力扣：
// https://leetcode.cn/problems/combination-sum-ii/description/
// 参考：
// https://programmercarl.com/0040.%E7%BB%84%E5%90%88%E6%80%BB%E5%92%8CII.html

// 时间复杂度：O(n * 2^n)
// 空间复杂度：O(target) ，递归栈
func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	path, result := []int{}, [][]int{}
	sort.Ints(candidates) //  这里是去重的关键逻辑，需要先排序
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
		if i > index && nums[i] == nums[i-1] { // 层去重，如果当前数和左边相邻的相等，则跳过，否则会有重复集合
			continue
		}
		if nums[i] > target {
			break
		}

		path = append(path, nums[i])
		dfs(nums, target-nums[i], i+1, path, result) // 这里是去重的关键逻辑,本次不取重复数字，下次循环可能会取重复数字
		path = path[:len(path)-1]
	}
}
