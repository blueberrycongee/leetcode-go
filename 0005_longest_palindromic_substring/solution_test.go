package longestpalindromicsubstring

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		s        string
		expected []string // 可能有多个正确答案
	}{
		{"babad", []string{"bab", "aba"}},
		{"cbbd", []string{"bb"}},
		{"a", []string{"a"}},
		{"ac", []string{"a", "c"}},
	}

	for _, tt := range tests {
		result := longestPalindrome(tt.s)
		valid := false
		for _, exp := range tt.expected {
			if result == exp {
				valid = true
				break
			}
		}
		if !valid {
			t.Errorf("longestPalindrome(%q) = %q, want one of %v", tt.s, result, tt.expected)
		}
	}
}
