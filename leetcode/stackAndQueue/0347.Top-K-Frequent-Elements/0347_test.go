package stackAndQueue

import (
	"reflect"
	"testing"
)

func TestTopKFrequent_Basic(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "示例 1: 基础案例",
			nums:     []int{1, 1, 1, 2, 2, 3},
			k:        2,
			expected: []int{1, 2}, // 1出现3次，2出现2次
		},
		{
			name:     "示例 2: 单个元素",
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
		{
			name: "多个元素频率相同",
			nums: []int{1, 1, 2, 2, 3, 3},
			k:    2,
			// 期望结果应为频率最高的2个元素，具体取决于堆排序的稳定性
			// 可能是 [1,2], [1,3] 或 [2,3]
		},
		{
			name: "所有元素唯一",
			nums: []int{1, 2, 3, 4, 5},
			k:    3,
			// 所有元素频率均为1，结果应为任意3个元素
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := topKFrequent(tt.nums, tt.k)

			// 验证结果长度是否正确
			if len(result) != tt.k {
				t.Errorf("topKFrequent() 返回长度 = %d, 期望长度 %d", len(result), tt.k)
				return
			}

			// 如果有确定的预期结果，进行验证
			if tt.expected != nil && !reflect.DeepEqual(result, tt.expected) {
				// 对于频率相同的情况，可能需要更灵活的验证
				if tt.name != "多个元素频率相同" && tt.name != "所有元素唯一" {
					t.Errorf("topKFrequent() = %v, 期望 %v", result, tt.expected)
				}
			}

			// 可以添加额外验证：确保结果中的元素确实在原数组中出现频率最高
		})
	}
}
