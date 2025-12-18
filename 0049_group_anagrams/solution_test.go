package groupanagrams

import (
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected [][]string
	}{
		{
			name:  "示例1",
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{
				{"eat", "tea", "ate"},
				{"tan", "nat"},
				{"bat"},
			},
		},
		{
			name:     "示例2-空字符串",
			input:    []string{""},
			expected: [][]string{{""}},
		},
		{
			name:     "示例3-单字符",
			input:    []string{"a"},
			expected: [][]string{{"a"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := groupAnagrams(tt.input)

			// 因为顺序可能不同，需要排序后比较
			if len(result) != len(tt.expected) {
				t.Errorf("长度不对: got %d, want %d", len(result), len(tt.expected))
				return
			}

			// 简单检查：每组内元素个数是否匹配
			sortGroups := func(groups [][]string) {
				for _, g := range groups {
					sort.Strings(g)
				}
				sort.Slice(groups, func(i, j int) bool {
					return groups[i][0] < groups[j][0]
				})
			}

			sortGroups(result)
			sortGroups(tt.expected)

			for i := range result {
				if len(result[i]) != len(tt.expected[i]) {
					t.Errorf("第%d组长度不对", i)
				}
			}
		})
	}
}
