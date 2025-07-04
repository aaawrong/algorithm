package linked_list

import "algorithm/leetcode/structures"

// 两两交换链表中的节点
// 力扣：https://leetcode.cn/problems/remove-nth-node-from-end-of-list
// 随想录：
// https://programmercarl.com/0019.%E5%88%A0%E9%99%A4%E9%93%BE%E8%A1%A8%E7%9A%84%E5%80%92%E6%95%B0%E7%AC%ACN%E4%B8%AA%E8%8A%82%E7%82%B9.html#%E7%AE%97%E6%B3%95%E5%85%AC%E5%BC%80%E8%AF%BE
type ListNode = structures.ListNode

// 时间复杂度：O(n)
// 空间复杂度：O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	slow, fast := dummyHead, dummyHead
	for i := 0; i <= n; i++ {
		if fast == nil { // 链表长度大于 n
			return head.Next
		}
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummyHead.Next
}
