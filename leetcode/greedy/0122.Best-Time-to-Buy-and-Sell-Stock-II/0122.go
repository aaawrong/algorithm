package greedy

// 买卖股票的最佳时机 II
// 力扣：
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/description/
// 随想录：
// https://programmercarl.com/0122.%E4%B9%B0%E5%8D%96%E8%82%A1%E7%A5%A8%E7%9A%84%E6%9C%80%E4%BD%B3%E6%97%B6%E6%9C%BAII.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC

// 时间复杂度：O(n)
// 空间复杂度：O(1)
func maxProfit(prices []int) int {
	var profit int
	for i := 1; i < len(prices); i++ { // 如果len为1，利润为0
		if p := prices[i] - prices[i-1]; p > 0 {
			profit += p
		}
	}
	return profit
}
