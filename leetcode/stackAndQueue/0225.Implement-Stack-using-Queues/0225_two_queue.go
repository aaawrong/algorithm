package stackAndQueue

// 用队列实现栈
// 力扣：
// https://leetcode.cn/problems/implement-stack-using-queues
// 随想录：

// 时间复杂度：O(n)
// 空间复杂度：O(n)
type MyStack2 struct {
	// 两个个队列实现，动图如下
	// https://leetcode.cn/problems/implement-stack-using-queues/solutions/432204/yong-dui-lie-shi-xian-zhan-by-leetcode-solution/
	// queue1负责存储，queue2负责进栈
	queue1, queue2 []int
}

func Constructor2() MyStack2 {
	return MyStack2{}
}

func (ms *MyStack2) Push(x int) {
	ms.queue2 = append(ms.queue2, x)
	for len(ms.queue1) > 0 {
		ms.queue2 = append(ms.queue2, ms.queue1[0])
		ms.queue1 = ms.queue1[1:]
	}
	ms.queue1, ms.queue2 = ms.queue2, ms.queue1
}

func (ms *MyStack2) Pop() int {
	x := ms.queue1[0]
	ms.queue1 = ms.queue1[1:]
	return x
}

func (ms *MyStack2) Top() int {
	return ms.queue1[0]
}

func (ms *MyStack2) Empty() bool {
	return len(ms.queue1) == 0
}
