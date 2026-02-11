package greedy

// 加油站
// 力扣：
// https://leetcode.cn/problems/gas-station/description/
// 随想录：
// https://programmercarl.com/0134.%E5%8A%A0%E6%B2%B9%E7%AB%99.html

// 时间复杂度：O(n)
// 空间复杂度：O(1)
func canCompleteCircuit(gas []int, cost []int) int {
	// curSum 到当前站能剩下的汽油
	// totalSum 总汽油
	curSum, totalSum, start := 0, 0, 0

	for i := range gas {
		curSum += gas[i] - cost[i]
		totalSum += gas[i] - cost[i]
		if curSum < 0 {
			curSum = 0
			start = i + 1 // 从下一站开始为起点
		}
	}

	if totalSum < 0 {
		return -1
	}
	return start
}

// 暴力解法
func canCompleteCircuit2(gas []int, cost []int) int {
	n := len(gas)
	for start := 0; start < n; start++ {
		tank := 0
		ok := true
		for i := 0; i < n; i++ {
			idx := (start + i) % n // 因为是环，直接index会越界
			tank += gas[idx] - cost[idx]
			if tank < 0 {
				ok = false
				break
			}
		}
		if ok {
			return start
		}
	}
	return -1

}
