package greedy

import "testing"

func TestWiggleMaxLength(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "empty array",
			nums: []int{},
			want: 0,
		},
		{
			name: "single element",
			nums: []int{5},
			want: 1,
		},
		{
			name: "two equal elements",
			nums: []int{2, 2},
			want: 1,
		},
		{
			name: "two increasing elements",
			nums: []int{1, 2},
			want: 2,
		},
		{
			name: "two decreasing elements",
			nums: []int{2, 1},
			want: 2,
		},
		{
			name: "classic wiggle sequence",
			nums: []int{1, 7, 4, 9, 2, 5},
			want: 6,
		},
		{
			name: "all equal",
			nums: []int{3, 3, 3, 3},
			want: 1,
		},
		{
			name: "monotonic increasing",
			nums: []int{1, 2, 3, 4, 5},
			want: 2,
		},
		{
			name: "monotonic decreasing",
			nums: []int{5, 4, 3, 2, 1},
			want: 2,
		},
		{
			name: "with flat segments",
			nums: []int{1, 1, 2, 2, 1},
			want: 3,
		},
		{
			name: "complex case",
			nums: []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8},
			want: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wiggleMaxLength(tt.nums)
			if got != tt.want {
				t.Fatalf("wiggleMaxLength(%v) = %d, want %d",
					tt.nums, got, tt.want)
			}
		})
	}
}
