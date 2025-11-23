package backtraking

// 组合
// 力扣：
// https://leetcode.cn/problems/combination-sum-iii/description/
// 参考：

// 时间复杂度：O(n * 2^n)， 遍历节点 O(2^n)，copy操作O(n)
// 空间复杂度：O(n)

func combinationSum3(k int, n int) [][]int {
	if k == 0 {
		return [][]int{}
	}
	path, res := []int{}, [][]int{}
	dfs(k, n, 1, path, &res)
	return res
}

func dfs(k, target, start int, path []int, res *[][]int) {
	if target == 0 && len(path) == k {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
	}

	for i := start; i < 10; i++ {
		if target >= i {
			path = append(path, i)
			dfs(k, target-i, i+1, path, res)
			path = path[:len(path)-1]
		}
	}
}
