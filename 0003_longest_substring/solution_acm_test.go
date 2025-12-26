package main

import "testing"

func TestACMMode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "示例1",
			input:    "abcabcbb",
			expected: 3,
		},
		{
			name:     "示例2",
			input:    "bbbbb",
			expected: 1,
		},
		{
			name:     "示例3",
			input:    "pwwkew",
			expected: 3,
		},
		{
			name:     "空字符串",
			input:    "",
			expected: 0,
		},
		{
			name:     "单字符",
			input:    "a",
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lengthOfLongestSubstring(tt.input)
			if result != tt.expected {
				t.Errorf("lengthOfLongestSubstring(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}
