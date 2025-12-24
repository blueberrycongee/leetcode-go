package longestvalidparentheses

import "testing"

func TestLongestValidParentheses(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{"(()", 2},
		{")()())", 4},
		{"", 0},
		{"()(())", 6},
		{"(()(()", 2},
	}

	for _, tt := range tests {
		result := longestValidParentheses(tt.s)
		if result != tt.expected {
			t.Errorf("longestValidParentheses(%q) = %d, want %d", tt.s, result, tt.expected)
		}
	}
}
