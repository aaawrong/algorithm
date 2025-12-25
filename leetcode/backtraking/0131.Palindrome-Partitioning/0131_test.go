package backtraking

import (
	"reflect"
	"testing"
)

func TestPartition_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected [][]string
	}{
		{
			name:  "basic",
			input: "aab",
			expected: [][]string{
				{"a", "a", "b"},
				{"aa", "b"},
			},
		},
		{
			name:  "empty",
			input: "",
			expected: [][]string{
				{},
			},
		},
		{
			name:  "single",
			input: "a",
			expected: [][]string{
				{"a"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := partition(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("组合内容不匹配\n期望: %v\n实际: %v", tt.expected, result)
			}
		})
	}

}
