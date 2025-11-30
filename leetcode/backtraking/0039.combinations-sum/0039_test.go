package backtraking

import (
	"reflect"
	"testing"
)

// 你的 combinationSum 函数放在这里

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			name:       "基础功能测试-candidates=[2,3,6,7], target=7",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			expected:   [][]int{{2, 2, 3}, {7}},
		},
		{
			name:       "多解情况测试-candidates=[2,3,5], target=8",
			candidates: []int{2, 3, 5},
			target:     8,
			expected:   [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name:       "无解情况-candidates=[2], target=1",
			candidates: []int{2},
			target:     1,
			expected:   [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := combinationSum(tt.candidates, tt.target)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("测试失败 %s\n输入: candidates=%v, target=%d\n期望: %v\n实际: %v",
					tt.name, tt.candidates, tt.target, tt.expected, result)
			}
		})
	}
}
