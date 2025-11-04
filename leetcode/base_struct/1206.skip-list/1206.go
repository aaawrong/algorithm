package base_struct

import (
	"fmt"
	"math/rand"
)

// 实现跳表
// 力扣：https://leetcode.cn/problems/design-skiplist/
// 参考：https://zhuanlan.zhihu.com/p/620291031

// Level 3:  1 ----------> 9
// Level 2:  1 ----> 5 --> 9
// Level 1:  1 -> 3 -> 5 -> 7 -> 9
// Level 0:  1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9
// 每一层是一条有序链表；
// 所有层共用同一个头节点 head；
// 每层的指针都保存在 head.nexts[i] 中。
// node7，它出现在 Level 2、Level 1、Level 0 三层：
//
//	node7.nexts = []*node{
//		 node8, // Level 0 下一节点是 8
//		 node9, // Level 1 下一节点是 9
//		 node12 // Level 2 下一节点是 12
//	}
//
// nexts[0] → 0 层的右边节点
// nexts[1] → 1 层的右边节点
// nexts[2] → 2 层的右边节点
type Skiplist struct {
	head *node
}

type node struct {
	nexts    []*node
	key, val int
}

// 根据 key 读取 val，第二个 bool flag 反映 key 在 skiplist 中是否存在
func (s *Skiplist) Get(key int) (int, bool) {
	// 根据 key 尝试检索对应的 node，如果 node 存在，则返回对应的 val
	if _node := s.search(key); _node != nil {
		return _node.val, true
	}

	return -1, false
}

// 从跳表中检索 key 对应的 node
func (s *Skiplist) search(key int) *node {
	// 每次检索从头部出发
	move := s.head
	// 每次检索从最大高度出发，直到来到首层
	for level := len(s.head.nexts) - 1; level >= 0; level-- {
		// 在每一层中持续向右遍历，直到下一个节点不存在或者 key 值大于等于 key
		for move.nexts[level] != nil && move.nexts[level].key < key {
			move = move.nexts[level]
		}
		// 如果 key 值相等，则找到了目标直接返回
		if move.nexts[level] != nil && move.nexts[level].key == key {
			return move.nexts[level]
		}

		// 当前层没找到目标，则层数减 1，继续向下
	}

	// 遍历完所有层数，都没有找到目标，返回 nil
	return nil
}

// roll 骰子，决定一个待插入的新节点在 skiplist 中最高层对应的 index
// 例如 level=2 → 该节点出现在 Level 2, 1, 0
func (s *Skiplist) roll() int {
	level := 0
	// 每次 50% 的概率升一层
	for rand.Float64() < 0.5 {
		level++
	}
	return level
}

// 将 key-val 对加入 skiplist
func (s *Skiplist) Put(key, val int) {
	// 假如 kv对已存在，则直接对值进行更新并返回
	if _node := s.search(key); _node != nil {
		_node.val = val
		return
	}

	// roll 出新节点的高度
	level := s.roll()

	// 新节点高度超出跳表最大高度，则需要对高度进行补齐
	for len(s.head.nexts)-1 < level {
		s.head.nexts = append(s.head.nexts, nil)
	}

	// 创建出新的节点
	newNode := node{
		key:   key,
		val:   val,
		nexts: make([]*node, level+1),
	}

	// 从头节点的最高层出发
	move := s.head
	for _level := level; _level >= 0; _level-- {
		// 向右遍历，直到右侧节点不存在或者 key 值大于 key
		// 前面已经判断过=的情况
		for move.nexts[_level] != nil && move.nexts[_level].key < key {
			move = move.nexts[_level]
		}

		// 调整指针关系，完成新节点的插入
		newNode.nexts[_level] = move.nexts[_level]
		move.nexts[_level] = &newNode
	}
}

// 根据 key 从跳表中删除对应的节点
func (s *Skiplist) Del(key int) {
	// 如果 kv 对不存在，则无需删除直接返回
	if _node := s.search(key); _node == nil {
		return
	}

	// 从头节点的最高层出发
	move := s.head
	for level := len(s.head.nexts) - 1; level >= 0; level-- {
		// 向右遍历，直到右侧节点不存在或者 key 值大于等于 key
		for move.nexts[level] != nil && move.nexts[level].key < key {
			move = move.nexts[level]
		}

		// 右侧节点不存在或者 key 值大于 target，则直接跳过
		if move.nexts[level] == nil || move.nexts[level].key > key {
			continue
		}

		// 走到此处意味着右侧节点的 key 值必然等于 key，则调整指针引用
		move.nexts[level] = move.nexts[level].nexts[level]
	}

	// 对跳表的最大高度进行更新
	var dif int
	// 倘若某一层已经不存在数据节点，高度需要递减
	for level := len(s.head.nexts) - 1; level > 0 && s.head.nexts[level] == nil; level-- {
		dif++
	}
	s.head.nexts = s.head.nexts[:len(s.head.nexts)-dif]
}

// 找到 skiplist 当中 ≥ start，且 ≤ end 的 kv 对
func (s *Skiplist) Range(start, end int) [][2]int {
	// 首先通过 ceiling 方法，找到 skiplist 中 key 值大于等于 start 且最接近于 start 的节点 ceilNode
	ceilNode := s.ceiling(start)
	// 如果不存在，直接返回
	if ceilNode == nil {
		return [][2]int{}
	}

	// 从 ceilNode 首层出发向右遍历，把所有位于 [start,end] 区间内的节点统统返回
	var res [][2]int
	for move := ceilNode; move != nil && move.key <= end; move = move.nexts[0] {
		res = append(res, [2]int{move.key, move.val})
	}
	return res
}

// 找到 key 值大于等于 target 且 key 值最接近于 target 的节点
func (s *Skiplist) ceiling(target int) *node {
	move := s.head

	// 自上而下，找到 key 值小于 target 且最接近 target 的 kv 对
	for level := len(s.head.nexts) - 1; level >= 0; level-- {
		for move.nexts[level] != nil && move.nexts[level].key < target {
			move = move.nexts[level]
		}
		// 如果 key 值等于 targe 的 kv 对存在，则直接返回
		if move.nexts[level] != nil && move.nexts[level].key == target {
			return move.nexts[level]
		}
	}

	// 此时 move 已经对应于在首层 key 值小于 key 且最接近于 key 的节点，其右侧第一个节点即为所寻找的目标节点
	return move.nexts[0]
}

// 找到 skiplist 中，key 值大于等于 target 且最接近于 target 的 key-value 对
func (s *Skiplist) Ceiling(target int) ([2]int, bool) {
	if ceilNode := s.ceiling(target); ceilNode != nil {
		return [2]int{ceilNode.key, ceilNode.val}, true
	}
	return [2]int{}, false
}

// 找到 skiplist 中，key 值小于等于 target 且最接近于 target 的 key-value 对
func (s *Skiplist) Floor(target int) ([2]int, bool) {
	// 引用 floor 方法，取 floorNode 值进行返回
	if floorNode := s.floor(target); floorNode != nil {
		return [2]int{floorNode.key, floorNode.val}, true
	}

	return [2]int{}, false
}

// 找到 key 值小于等于 target 且 key 值最接近于 target 的节点
func (s *Skiplist) floor(target int) *node {
	move := s.head

	// 自上而下，找到 key 值小于 target 且最接近 target 的 kv 对
	for level := len(s.head.nexts) - 1; level >= 0; level-- {
		for move.nexts[level] != nil && move.nexts[level].key < target {
			move = move.nexts[level]
		}
		// 如果 key 值等于 targe 的 kv对存在，则直接返回
		if move.nexts[level] != nil && move.nexts[level].key == target {
			return move.nexts[level]
		}
	}

	// move 是首层中 key 值小于 target 且最接近 target 的节点，直接返回 move 即可
	return move
}

func (s *Skiplist) Print() {
	fmt.Println("====== Skiplist ======")
	for level := len(s.head.nexts) - 1; level >= 0; level-- {
		fmt.Printf("Level %d: ", level)
		_node := s.head.nexts[level]
		for _node != nil {
			fmt.Printf("%d ", _node.key)
			_node = _node.nexts[level]
		}
		fmt.Println()
	}
	fmt.Println("======================")
}
