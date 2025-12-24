package findallanagrams

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		p        string
		expected []int
	}{
		{
			name:     "示例1",
			s:        "cbaebabacd",
			p:        "abc",
			expected: []int{0, 6},
		},
		{
			name:     "示例2",
			s:        "abab",
			p:        "ab",
			expected: []int{0, 1, 2},
		},
		{
			name:     "s比p短",
			s:        "a",
			p:        "ab",
			expected: nil,
		},
		{
			name:     "无匹配",
			s:        "aaaa",
			p:        "bc",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findAnagrams(tt.s, tt.p)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("findAnagrams(%q, %q) = %v, want %v", tt.s, tt.p, result, tt.expected)
			}
		})
	}
}
