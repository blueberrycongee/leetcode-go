package minimumwindowsubstring

import "testing"

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		target   string
		expected string
	}{
		{
			name:     "示例1",
			s:        "ADOBECODEBANC",
			target:   "ABC",
			expected: "BANC",
		},
		{
			name:     "示例2",
			s:        "a",
			target:   "a",
			expected: "a",
		},
		{
			name:     "示例3-不存在",
			s:        "a",
			target:   "aa",
			expected: "",
		},
		{
			name:     "完全匹配",
			s:        "abc",
			target:   "abc",
			expected: "abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minWindow(tt.s, tt.target)
			if result != tt.expected {
				t.Errorf("minWindow(%q, %q) = %q, want %q", tt.s, tt.target, result, tt.expected)
			}
		})
	}
}
