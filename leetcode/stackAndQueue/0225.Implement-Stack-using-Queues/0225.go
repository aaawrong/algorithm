package stackAndQueue

// 用队列实现栈
// 力扣：
// https://leetcode.cn/problems/implement-stack-using-queues
// 随想录：

// 时间复杂度：O(n)
// 空间复杂度：O(n)
type MyStack struct {
	// 单个队列实现，动图如下
	// https://leetcode.cn/problems/implement-stack-using-queues/solutions/432204/yong-dui-lie-shi-xian-zhan-by-leetcode-solution/
	queue []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (ms *MyStack) Push(x int) {
	n := len(ms.queue)
	ms.queue = append(ms.queue, x)
	for ; n > 0; n-- {
		ms.queue = append(ms.queue, ms.queue[0])
		ms.queue = ms.queue[1:]
	}
}

func (ms *MyStack) Pop() int {
	x := ms.queue[0]
	ms.queue = ms.queue[1:]
	return x
}

func (ms *MyStack) Top() int {
	return ms.queue[0]
}

func (ms *MyStack) Empty() bool {
	return len(ms.queue) == 0
}
