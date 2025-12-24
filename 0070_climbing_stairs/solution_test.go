package climbingstairs

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{10, 89},
	}

	for _, tt := range tests {
		result := climbStairs(tt.n)
		if result != tt.expected {
			t.Errorf("climbStairs(%d) = %d, want %d", tt.n, result, tt.expected)
		}
	}
}
