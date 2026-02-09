package greedy

// 跳跃游戏
// 力扣：
//
// 随想录：
//

// 时间复杂度：O(n)
// 空间复杂度：O(1)
func jump(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	temp, cover, step := 0, 0, 0 // temp记录下一个最大值
	for i := 0; i < n; i++ {
		temp = max(temp, i+nums[i])
		if i == cover {
			step++
			cover = temp
			if cover >= n-1 {
				break
			}
		}
	}
	return step
}
