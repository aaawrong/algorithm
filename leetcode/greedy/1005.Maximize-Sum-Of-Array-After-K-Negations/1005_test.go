package greedy

import "testing"

func TestLargestSumAfterKNegations(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "基础情况",
			nums: []int{4, 2, 3},
			k:    1,
			want: 5,
		},
		{
			name: "多个负数",
			nums: []int{3, -1, 0, 2},
			k:    3,
			want: 6,
		},
		{
			name: "全负数",
			nums: []int{-2, -3, -1},
			k:    2,
			want: 4,
		},
		{
			name: "k大于长度",
			nums: []int{2, -3, -1, 5, -4},
			k:    10,
			want: 13,
		},
		{
			name: "含有0",
			nums: []int{-1, 0, 2},
			k:    3,
			want: 3,
		},
		{
			name: "全正数且k为奇数",
			nums: []int{1, 2, 3},
			k:    1,
			want: 4,
		},
		{
			name: "全正数且k为偶数",
			nums: []int{1, 2, 3},
			k:    2,
			want: 6,
		},
		{
			name: "单元素负数",
			nums: []int{-5},
			k:    3,
			want: 5,
		},
		{
			name: "单元素正数",
			nums: []int{5},
			k:    3,
			want: -5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := append([]int{}, tt.nums...)
			got := largestSumAfterKNegations(input, tt.k)
			if got != tt.want {
				t.Errorf("largestSumAfterKNegations() = %v, want %v", got, tt.want)
			}
		})
	}

}
