package maximumproductsubarray

import "testing"

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
		{[]int{-2, 3, -4}, 24},
		{[]int{-2}, -2},
	}

	for _, tt := range tests {
		result := maxProduct(tt.nums)
		if result != tt.expected {
			t.Errorf("maxProduct(%v) = %d, want %d", tt.nums, result, tt.expected)
		}
	}
}
