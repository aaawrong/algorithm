package backtraking

// 组合
// 力扣：
// https://leetcode.cn/problems/combinations/description/
// 参考：

// 时间复杂度：O(n * 2^n)， 遍历节点 O(2^n)，copy操作O(n)
// 空间复杂度：O(n)

func combine(n int, k int) [][]int {
	if n <= 0 || k <= 0 || k > n {
		return [][]int{}
	}
	path, res := []int{}, [][]int{}
	dfs(n, k, 1, path, &res)
	return res
}

// path表示路径已经存的元素
func dfs(n, k, start int, path []int, res *[][]int) {
	if len(path) == k {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	// k-len(path) 表示还需要存几个元素
	// n-(k-len(path))+1 表示至多从几开始，即上限
	for i := start; i <= n-(k-len(path))+1; i++ {
		path = append(path, i)
		dfs(n, k, i+1, path, res)
		path = path[:len(path)-1] // 回到上个节点
	}
	return
}
