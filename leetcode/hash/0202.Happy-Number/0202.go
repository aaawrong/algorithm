package hash

// 快乐数
// 力扣：
// https://leetcode.cn/problems/happy-number/description/
// 随想录：
//

// 哈希表法
// 时间复杂度：O(logn)，解释n = 1000 → log1000 = 3，位数 d=4（实际是 4 位）,所以位数d = O(logn)
// 空间复杂度：O(logn)，
// 解释 n 最大= 999（3位数）→ 下一值 81 × 3 = 243（3位数）
// 设当前值为 n，其位数 d = ⌊log₁₀n⌋ + 1 → 下一值 ≤ 81d = O(log n)。
func isHappy(n int) bool {
	seen := make(map[int]bool)
	for n != 1 && !seen[n] {
		seen[n] = true
		n = getNext(n)
	}
	return n == 1
}

// 快慢指针法
// 时间复杂度：O(logn)
// 空间复杂度：O(1)
func isHappy2(n int) bool {
	slow, fast := n, getNext(n)
	for slow != 1 && slow != fast {
		slow = getNext(slow)
		fast = getNext(getNext(fast))
	}
	return slow == 1
}

func getNext(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}
