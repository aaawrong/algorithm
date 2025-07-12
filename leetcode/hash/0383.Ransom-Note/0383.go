package hash

// 两个数组的交集
// 力扣：
// https://leetcode.cn/problems/ransom-note
// 随想录：
// https://programmercarl.com/0383.%E8%B5%8E%E9%87%91%E4%BF%A1.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC

// 时间复杂度：O(n)
// 空间复杂度：O(n)
func canConstruct(ransomNote string, magazine string) bool {
	record := make([]int, 26)
	for _, v := range magazine {
		record[v-'a']++
	}
	for _, v := range ransomNote {
		record[v-'a']--
		if record[v-'a'] < 0 {
			return false
		}
	}
	return true
}
