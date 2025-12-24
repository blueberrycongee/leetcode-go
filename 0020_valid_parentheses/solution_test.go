package validparentheses

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		s        string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"", true},
	}

	for _, tt := range tests {
		result := isValid(tt.s)
		if result != tt.expected {
			t.Errorf("isValid(%q) = %v, want %v", tt.s, result, tt.expected)
		}
	}
}
