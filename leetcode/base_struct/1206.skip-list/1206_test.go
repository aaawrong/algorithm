package base_struct

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 初始化一个空的 Skiplist
func newSkiplist() *Skiplist {
	return &Skiplist{
		head: &node{
			nexts: []*node{nil},
		},
	}
}
func TestSkiplistBasic(t *testing.T) {
	s := newSkiplist()

	// 插入若干元素
	data := map[int]int{
		3: 30,
		1: 10,
		5: 50,
		2: 20,
		4: 40,
	}

	for k, v := range data {
		s.Put(k, v)
	}
	s.Print()
	// 测试 Get
	for k, v := range data {
		if val, ok := s.Get(k); !ok || val != v {
			t.Fatalf("Get(%d) = (%d, %v), want (%d, true)", k, val, ok, v)
		}
	}

	// 测试不存在的 key
	if _, ok := s.Get(100); ok {
		t.Fatalf("Get(100) should not exist")
	}

	// 测试更新已存在 key
	s.Put(3, 99)
	if val, ok := s.Get(3); !ok || val != 99 {
		t.Fatalf("update failed: Get(3) = %d, want 99", val)
	}

	// 测试 Floor
	if kv, ok := s.Floor(4); !ok || kv[0] != 4 {
		t.Fatalf("Floor(4) = %v, want key=4", kv)
	}
	if kv, ok := s.Floor(6); !ok || kv[0] != 5 {
		t.Fatalf("Floor(6) = %v, want key=5", kv)
	}

	// 测试 Ceiling
	if kv, ok := s.Ceiling(2); !ok || kv[0] != 2 {
		t.Fatalf("Ceiling(2) = %v, want key=2", kv)
	}
	if kv, ok := s.Ceiling(0); !ok || kv[0] != 1 {
		t.Fatalf("Ceiling(0) = %v, want key=1", kv)
	}

	// 测试 Range
	rangeRes := s.Range(2, 4)
	expect := [][2]int{{2, 20}, {3, 99}, {4, 40}}
	if len(rangeRes) != len(expect) {
		t.Fatalf("Range length = %d, want %d", len(rangeRes), len(expect))
	}
	for i := range expect {
		if rangeRes[i] != expect[i] {
			t.Fatalf("Range mismatch at %d: got %v, want %v", i, rangeRes[i], expect[i])
		}
	}

	// 测试删除
	s.Del(3)
	if _, ok := s.Get(3); ok {
		t.Fatalf("Del(3) failed: key still exists")
	}

	// 删除不存在 key
	s.Del(100) // 应该安全返回

	// 验证删除后不影响其他节点
	if val, ok := s.Get(4); !ok || val != 40 {
		t.Fatalf("After Del(3), Get(4) = (%d,%v), want (40,true)", val, ok)
	}
}

func TestSkiplistEmpty(t *testing.T) {
	s := newSkiplist()
	if _, ok := s.Get(1); ok {
		t.Fatalf("Get on empty skiplist should return not found")
	}
	if res := s.Range(1, 5); len(res) != 0 {
		t.Fatalf("Range on empty skiplist should return empty slice")
	}
	if _, ok := s.Ceiling(10); ok {
		t.Fatalf("Ceiling on empty skiplist should return not found")
	}
	if _, ok := s.Floor(10); ok {
		t.Fatalf("Floor on empty skiplist should return not found")
	}
}
