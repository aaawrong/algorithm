package greedy

// 跳跃游戏
// 力扣：
// https://leetcode.cn/problems/jump-game/description/
// 随想录：
// https://programmercarl.com/0055.%E8%B7%B3%E8%B7%83%E6%B8%B8%E6%88%8F.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC

// 时间复杂度：O(n)
// 空间复杂度：O(1)
func canJump(nums []int) bool {
	if len(nums) <= 1 { // 长度为1的情况要放在这里，不能放在循环，比如[0]
		return true
	}

	cover := 0
	for i := 0; i <= cover; i++ {
		cover = max(cover, i+nums[i])
		if cover >= len(nums)-1 { // 注意是len-1
			return true
		}
	}
	return false
}
