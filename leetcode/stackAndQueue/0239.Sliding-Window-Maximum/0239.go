package stackAndQueue

// 滑动窗口最大值
// 力扣：
// https://leetcode.cn/problems/sliding-window-maximum/description/
// 随想录：
// https://programmercarl.com/0239.%E6%BB%91%E5%8A%A8%E7%AA%97%E5%8F%A3%E6%9C%80%E5%A4%A7%E5%80%BC.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC

// 时间复杂度：O(n*k)
// 空间复杂度：O(n)
// 暴力法
func maxSlidingWindow1(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	n := len(nums)
	res := make([]int, 0, n-k+1)
	for i := 0; i <= n-k; i++ {
		_max := nums[i]
		for j := 1; j < k; j++ {
			if _max < nums[i+j] {
				_max = nums[i+j]
			}
		}
		res = append(res, _max)
	}

	return res
}

// 时间复杂度：O(n)
// 空间复杂度：O(k)
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || len(nums) < k {
		return []int{}
	}
	window := make([]int, 0, k) // 存储对应的索引位置
	result := make([]int, 0, len(nums)-k+1)
	for i, v := range nums {
		if i >= k && window[0] <= i-k { // 滑动窗口移除最左侧元素
			window = window[1:]
		}

		for len(window) > 0 && nums[window[len(window)-1]] < v { // 最左侧只保留当前窗口最大值
			window = window[0 : len(window)-1]
		}
		window = append(window, i)

		if i >= k-1 {
			result = append(result, nums[window[0]]) // 当遍历到索引 k-1（即第k个元素）时，才形成了第一个完整的窗口
		}
	}
	return result
}
