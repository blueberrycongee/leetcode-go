package houserobber

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
		{[]int{2, 1}, 2},
		{[]int{1}, 1},
	}

	for _, tt := range tests {
		result := rob(tt.nums)
		if result != tt.expected {
			t.Errorf("rob(%v) = %d, want %d", tt.nums, result, tt.expected)
		}
	}
}
