package stackAndQueue

import (
	"testing"
)

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		name     string
		tokens   []string
		expected int
	}{
		{
			name:     "基本加法",
			tokens:   []string{"2", "1", "+", "3", "*"},
			expected: 9, // (2 + 1) * 3 = 9
		},
		{
			name:     "包含除法",
			tokens:   []string{"4", "13", "5", "/", "+"},
			expected: 6, // 4 + (13 / 5) = 4 + 2 = 6
		},
		{
			name:     "简单减法",
			tokens:   []string{"3", "4", "-"},
			expected: -1,
		},
		{
			name:     "单个数字",
			tokens:   []string{"42"},
			expected: 42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := evalRPN(tt.tokens)
			if got != tt.expected {
				t.Errorf("evalRPN(%v) = %d; want %d", tt.tokens, got, tt.expected)
			}
		})
	}
}
