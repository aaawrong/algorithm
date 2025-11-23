package backtraking

import (
	"reflect"
	"testing"
)

func TestCombinationSum3(t *testing.T) {
	tests := []struct {
		name     string
		k        int
		n        int
		expected [][]int
	}{
		{
			name:     "基础功能 k=3, n=7",
			k:        3,
			n:        7,
			expected: [][]int{{1, 2, 4}},
		},
		{
			name:     "多解情况 k=3, n=9",
			k:        3,
			n:        9,
			expected: [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}},
		},
		{
			name:     "无解情况 k=4, n=1",
			k:        4,
			n:        1,
			expected: [][]int{},
		},
		{
			name:     "无解情况 k=3, n=2",
			k:        3,
			n:        2,
			expected: [][]int{},
		},
		{
			name:     "最大边界 k=9, n=45",
			k:        9,
			n:        45,
			expected: [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		},
		{
			name:     "k=1 的边界",
			k:        1,
			n:        5,
			expected: [][]int{{5}},
		},
		{
			name:     "目标值超范围",
			k:        3,
			n:        60,
			expected: [][]int{},
		},
		{
			name:     "空结果检查",
			k:        2,
			n:        1,
			expected: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := combinationSum3(tt.k, tt.n)

			// 首先检查结果长度
			if len(result) != len(tt.expected) {
				t.Errorf("组合数量不正确: 期望 %d, 实际 %d", len(tt.expected), len(result))
				return
			}

			if len(tt.expected) > 0 {
				// 由于结果顺序可能不同，使用反射的DeepEqual检查整体匹配
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("组合内容不匹配\n期望: %v\n实际: %v", tt.expected, result)
				}
			} else {
				// 期望空结果的情况
				if len(result) != 0 {
					t.Errorf("期望空结果，但得到: %v", result)
				}
			}
		})
	}
}
