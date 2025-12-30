package backtraking

import (
	"reflect"
	"sort"
	"testing"
)

func TestRestoreIpAddresses(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "经典案例 25525511135",
			input:    "25525511135",
			expected: []string{"255.255.11.135", "255.255.111.35"},
		},
		{
			name:     "全零情况",
			input:    "0000",
			expected: []string{"0.0.0.0"},
		},
		{
			name:     "简单重复数字",
			input:    "1111",
			expected: []string{"1.1.1.1"},
		},
		{
			name:     "包含前导零的情况",
			input:    "010010",
			expected: []string{"0.10.0.10", "0.100.1.0"},
		},
		{
			name:     "多种分割可能",
			input:    "101023",
			expected: []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"},
		},
		{
			name:     "空字符串",
			input:    "",
			expected: []string{},
		},
		{
			name:     "过短字符串",
			input:    "123",
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := restoreIpAddresses(tt.input)

			// 对结果和期望值进行排序，因为顺序不重要
			sort.Strings(result)
			sort.Strings(tt.expected)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("测试失败 %s\n输入: %s\n期望: %v\n实际: %v", tt.name, tt.input, tt.expected, result)
			}
		})
	}
}
