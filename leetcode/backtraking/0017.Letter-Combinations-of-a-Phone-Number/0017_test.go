package backtraking

import (
	"reflect"
	"sort"
	"testing"
)

// 你的 letterCombinationsBT 函数放在这里

func TestLetterCombinationsBT(t *testing.T) {
	tests := []struct {
		name     string
		digits   string
		expected []string
	}{
		{
			name:     "基础功能测试-两数字组合",
			digits:   "23",
			expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name:     "单数字输入",
			digits:   "2",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "空输入处理",
			digits:   "",
			expected: []string{},
		},
		{
			name:     "包含4字母的数字7",
			digits:   "7",
			expected: []string{"p", "q", "r", "s"},
		},
		{
			name:     "包含4字母的数字9",
			digits:   "9",
			expected: []string{"w", "x", "y", "z"},
		},
		{
			name:     "重复数字测试",
			digits:   "22",
			expected: []string{"aa", "ab", "ac", "ba", "bb", "bc", "ca", "cb", "cc"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := letterCombinationsBT2(tt.digits)

			// 对结果和期望值进行排序，因为题目不要求特定顺序
			sort.Strings(result)
			sort.Strings(tt.expected)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("%s失败\n输入: %s\n期望: %v\n实际: %v",
					tt.name, tt.digits, tt.expected, result)
			}
		})
	}
}
