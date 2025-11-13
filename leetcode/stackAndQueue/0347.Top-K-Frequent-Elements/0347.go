package stackAndQueue

import "container/heap"

// 前 K 个高频元素
// 力扣：
// https://leetcode.cn/problems/top-k-frequent-elements/description/
// 随想录：
//

// 时间复杂度：O(nlogk)
// 空间复杂度：O(n)
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}
	q := PriorityQueue{}
	for key, count := range m {
		if len(q) < k {
			heap.Push(&q, &Item{key: key, count: count})
		} else {
			if q[0].count < count {
				heap.Pop(&q)
				heap.Push(&q, &Item{key: key, count: count})
			}
		}
	}
	result := make([]int, k)
	for i := range k {
		item := heap.Pop(&q).(*Item)
		result[k-i-1] = item.key

	}
	return result
}

// Item define
type Item struct {
	key   int
	count int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// 注意：因为golang中的heap是按最小堆组织的，所以count越大，Less()越小，越靠近堆顶.
	return pq[i].count < pq[j].count
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push define
func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

// Pop define
func (pq *PriorityQueue) Pop() any {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}
