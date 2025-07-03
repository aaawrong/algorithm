package linked_list

import "algorithm/leetcode/structures"

// 反转链表
// 力扣：https://leetcode.cn/problems/reverse-linked-list/description/
// 随想录：https://programmercarl.com/0206.%E7%BF%BB%E8%BD%AC%E9%93%BE%E8%A1%A8.html#%E6%80%9D%E8%B7%AF

type ListNode = structures.ListNode

// 双指针
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}
	return prev
}

// TODO递归
// 时间复杂度：O(N)
// 空间复杂度：O(N)
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

/*第1步：递归到底
reverseList(1) 调用 reverseList(2)
reverseList(2) 调用 reverseList(3)
reverseList(3) 调用 reverseList(4)
reverseList(4) 返回 4（因为 4.Next == nil）

第2步：从后往前反转
在 reverseList(3) 中：
- newHead = 4
- 3.Next.Next = 3  (让 4 指向 3)
- 3.Next = nil
- 返回 newHead (4)
现在：4 → 3 → null

在 reverseList(2) 中：
- newHead = 4
- 2.Next.Next = 2  (让 3 指向 2)
- 2.Next = nil
- 返回 newHead (4)
现在：4 → 3 → 2 → null

在 reverseList(1) 中：
- newHead = 4
- 1.Next.Next = 1  (让 2 指向 1)
- 1.Next = nil
- 返回 newHead (4)
最终：4 → 3 → 2 → 1 → null*/
