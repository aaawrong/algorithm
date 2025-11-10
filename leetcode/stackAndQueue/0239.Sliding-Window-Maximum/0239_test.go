package stackAndQueue

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "常规情况",
			nums:     []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:        3,
			expected: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name:     "窗口为1",
			nums:     []int{1, -1},
			k:        1,
			expected: []int{1, -1},
		},
		{
			name:     "窗口等于数组长度",
			nums:     []int{1, 3, 1, 2, 0, 5},
			k:        6,
			expected: []int{5},
		},
		{
			name:     "数组为空",
			nums:     []int{},
			k:        3,
			expected: []int{},
		},
		{
			name:     "k大于数组长度",
			nums:     []int{2, 4},
			k:        3,
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSlidingWindow(tt.nums, tt.k)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("maxSlidingWindow(%v, %d) = %v; want %v", tt.nums, tt.k, got, tt.expected)
			}
		})
	}
}
