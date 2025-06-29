package lianbiao

import "algorithm/leetcode/structures"

type ListNode = structures.ListNode

// dummyhead -> 1 -> 2 -> 3 -> 4 -> nil
// temp保存1, temp1保存3
//O(n),O(1)
func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	cur := dummyHead
	for cur.Next != nil && cur.Next.Next != nil {
		temp := cur.Next
		temp1 := cur.Next.Next.Next
		cur.Next = cur.Next.Next
		cur.Next.Next = temp
		temp.Next = temp1
		cur = cur.Next.Next
	}
	return dummyHead.Next
}
