package uniquepaths

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		m, n     int
		expected int
	}{
		{3, 7, 28},
		{3, 2, 3},
		{1, 1, 1},
		{3, 3, 6},
	}

	for _, tt := range tests {
		result := uniquePaths(tt.m, tt.n)
		if result != tt.expected {
			t.Errorf("uniquePaths(%d, %d) = %d, want %d", tt.m, tt.n, result, tt.expected)
		}
	}
}
