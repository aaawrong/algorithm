package linked_list

import "algorithm/leetcode/structures"

// 两两交换链表中的节点
// 力扣：https://leetcode.cn/problems/swap-nodes-in-pairs
// 随想录：https://programmercarl.com/0024.%E4%B8%A4%E4%B8%A4%E4%BA%A4%E6%8D%A2%E9%93%BE%E8%A1%A8%E4%B8%AD%E7%9A%84%E8%8A%82%E7%82%B9.html#%E7%AE%97%E6%B3%95%E5%85%AC%E5%BC%80%E8%AF%BE

type ListNode = structures.ListNode

// 时间复杂度：O(n)
// 空间复杂度：O(1)
// dummyhead -> 1 -> 2 -> 3 -> 4 -> nil
// temp保存1, temp1保存3
func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	cur := dummyHead
	for cur.Next != nil && cur.Next.Next != nil { // 顺序不能错，不能用||，不能会空指针
		temp := cur.Next
		temp1 := cur.Next.Next.Next
		cur.Next = cur.Next.Next
		cur.Next.Next = temp
		temp.Next = temp1
		cur = cur.Next.Next
	}
	return dummyHead.Next
}
