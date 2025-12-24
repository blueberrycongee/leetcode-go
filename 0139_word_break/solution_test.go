package wordbreak

import "testing"

func TestWordBreak(t *testing.T) {
	tests := []struct {
		s        string
		wordDict []string
		expected bool
	}{
		{"leetcode", []string{"leet", "code"}, true},
		{"applepenapple", []string{"apple", "pen"}, true},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
	}

	for _, tt := range tests {
		result := wordBreak(tt.s, tt.wordDict)
		if result != tt.expected {
			t.Errorf("wordBreak(%q, %v) = %v, want %v", tt.s, tt.wordDict, result, tt.expected)
		}
	}
}
