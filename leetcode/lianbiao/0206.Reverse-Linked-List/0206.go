package lianbiao

import "algorithm/leetcode/structures"

type ListNode = structures.ListNode

// 双指针
// O(N) , O(1)
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}

// 递归
// O(N), O(N)
func reverseList2(head *ListNode) *ListNode {
	return recur(nil, head)
}

func recur(pre, head *ListNode) *ListNode {
	cur := head
	if cur == nil {
		return pre
	}
	temp := cur.Next
	cur.Next = pre
	return recur(cur, temp)
}
