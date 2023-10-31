package lianbiao

import "algorithm/leetcode/structures"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = structures.ListNode

// 时间复杂度: O(n)
// 空间复杂度: O(1)
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := dummyHead
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}
