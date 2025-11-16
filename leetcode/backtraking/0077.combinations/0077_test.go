package base_struct

import (
	"reflect"
	"testing"
)

func TestCombine(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		k        int
		expected [][]int
	}{
		{
			name:     "基础案例-从4个数取2个",
			n:        4,
			k:        2,
			expected: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			name:     "k=1的情况-每个数单独成组",
			n:        3,
			k:        1,
			expected: [][]int{{1}, {2}, {3}},
		},
		{
			name:     "k=n的情况-唯一组合",
			n:        3,
			k:        3,
			expected: [][]int{{1, 2, 3}},
		},
		{
			name: "较大n和k",
			n:    5,
			k:    3,
			expected: [][]int{
				{1, 2, 3}, {1, 2, 4}, {1, 2, 5}, {1, 3, 4}, {1, 3, 5},
				{1, 4, 5}, {2, 3, 4}, {2, 3, 5}, {2, 4, 5}, {3, 4, 5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := combine(tt.n, tt.k)

			if len(result) != len(tt.expected) {
				t.Errorf("组合数量不正确: 期望 %d, 实际 %d", len(tt.expected), len(result))
				return
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("组合内容不匹配\n期望: %v\n实际: %v", tt.expected, result)
			}
		})
	}
}
