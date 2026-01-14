package backtraking

import "sort"

// 组合
// 力扣：
// https://leetcode.cn/problems/permutations-ii/description/
// 参考：
// https://programmercarl.com/0047.%E5%85%A8%E6%8E%92%E5%88%97II.html#%E6%80%9D%E8%B7%AF

// 时间复杂度：O(n * n!)
// 空间复杂度：O(n)
func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	used, path, result := make([]bool, len(nums)), []int{}, [][]int{}
	sort.Ints(nums)
	genPermutations(0, path, nums, &result, &used)
	return result
}

func genPermutations(index int, path, nums []int, result *[][]int, used *[]bool) {
	if index == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			if i > 0 && (nums[i-1] == nums[i]) && !(*used)[i-1] {
				continue
			}
			(*used)[i] = true
			path = append(path, nums[i])
			genPermutations(index+1, path, nums, result, used)
			(*used)[i] = false
			path = path[:len(path)-1]
		}
	}
}
