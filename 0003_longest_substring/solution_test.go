package longestsubstring

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "示例1",
			s:        "abcabcbb",
			expected: 3,
		},
		{
			name:     "示例2-全相同",
			s:        "bbbbb",
			expected: 1,
		},
		{
			name:     "示例3",
			s:        "pwwkew",
			expected: 3,
		},
		{
			name:     "空字符串",
			s:        "",
			expected: 0,
		},
		{
			name:     "单字符",
			s:        "a",
			expected: 1,
		},
		{
			name:     "无重复",
			s:        "abcdef",
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lengthOfLongestSubstring(tt.s)
			if result != tt.expected {
				t.Errorf("lengthOfLongestSubstring(%q) = %d, want %d", tt.s, result, tt.expected)
			}
		})
	}
}
