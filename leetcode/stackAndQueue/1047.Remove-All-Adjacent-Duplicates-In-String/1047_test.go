package stackAndQueue

import (
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"abbaca", "ca"},
		{"azxxzy", "ay"},
		{"aabcca", "ba"},
		{"a", "a"},
		{"", ""},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := removeDuplicates(tt.input); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
