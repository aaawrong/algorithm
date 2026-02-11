package greedy

// 分发糖果
// 力扣：
// https://leetcode.cn/problems/candy/description/
// 随想录：
// https://programmercarl.com/0135.%E5%88%86%E5%8F%91%E7%B3%96%E6%9E%9C.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC

// 时间复杂度：O(n)
// 空间复杂度：O(n)
func candy(ratings []int) int {
	candies := make([]int, len(ratings))

	// 从左到右，如果 左边 > 右边 则 左边的糖果+1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// 从右到左，如果 左边 > 右边 则 再判断 糖果是否<=右边的糖，因为需要把左到右的照顾到
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}

	total := len(candies) // 默认每个小孩分的1个糖果
	for _, c := range candies {
		total += c
	}
	return total
}
