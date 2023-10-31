package shuzu

// O(n^2)
// O(1)
func generateMatrix(n int) [][]int {
	startx, starty, offset := 0, 0, 1
	count, loop := 1, n/2 //loop控制循环

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
