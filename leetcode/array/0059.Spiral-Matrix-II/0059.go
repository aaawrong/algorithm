package array

// 螺旋矩阵II
// 力扣：https://leetcode.cn/problems/spiral-matrix-ii/description/
// 随想录：https://programmercarl.com/0059.%E8%9E%BA%E6%97%8B%E7%9F%A9%E9%98%B5II.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC

// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func generateMatrix(n int) [][]int {
	startx, starty, offset := 0, 0, 1
	count, loop := 1, n/2 // loop控制循环，转一圈走两行

	//初始化二维数组
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}

	for loop > 0 {
		i, j := startx, starty

		for ; j < n-offset; j++ {
			res[i][j] = count
			count++
		}
		for ; i < n-offset; i++ {
			res[i][j] = count
			count++
		}
		for ; j > starty; j-- {
			res[i][j] = count
			count++
		}
		for ; i > startx; i-- {
			res[i][j] = count
			count++
		}
		startx++
		starty++
		offset++
		loop--
	}

	//n为奇数
	if n%2 == 1 {
		res[n/2][n/2] = n * n
	}
	return res
}
