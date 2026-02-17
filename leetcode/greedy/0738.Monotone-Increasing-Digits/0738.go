package greedy

import "strconv"

// 单调递增的数字
// 力扣：
// https://leetcode.cn/problems/monotone-increasing-digits/description/
// 随想录：
// https://programmercarl.com/0738.%E5%8D%95%E8%B0%83%E9%80%92%E5%A2%9E%E7%9A%84%E6%95%B0%E5%AD%97.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC

// 时间复杂度：O(k) = O(log n)
// 空间复杂度：O(k)
func monotoneIncreasingDigits(n int) int {
	s := []byte(strconv.Itoa(n))
	for i := len(s) - 2; i >= 0; i-- { // 从后开始遍历，找到第一个不符合单调递增的
		if s[i] > s[i+1] {
			s[i]--
			for j := i + 1; j < len(s); j++ { // 后面全部置为9
				s[j] = '9'
			}
		}
	}

	res, _ := strconv.Atoi(string(s))
	return res
}
