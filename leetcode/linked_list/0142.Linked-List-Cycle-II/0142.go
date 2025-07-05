package linked_list

import "algorithm/leetcode/structures"

// 环形链表II
// 力扣：
// https://leetcode.cn/problems/linked-list-cycle-ii/description
// 随想录：
// https://programmercarl.com/0142.%E7%8E%AF%E5%BD%A2%E9%93%BE%E8%A1%A8II.html#%E6%80%9D%E8%B7%AF

type ListNode = structures.ListNode

// 遍历链表中的每个节点，并将它记录下来；一旦遇到了此前遍历过的节点，就可以判定链表中存在环。借助哈希表可以很方便地实现。
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func detectCycle(head *ListNode) *ListNode {
	seen := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := seen[head]; ok {
			return head
		}
		seen[head] = struct{}{}
		head = head.Next
	}
	return nil
}

// 快慢双指针，看随想录解释
// (x + y) * 2 = x + y + n (y + z)  ->  x = (n - 1) (y + z) + z
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func detectCycle2(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast { // 当节点相遇，head和slow同时开始移动，相等即是入口
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return head
		}
	}
	return nil
}
