package longestincreasingsubsequence

import "testing"

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{7, 7, 7, 7, 7, 7, 7}, 1},
		{[]int{1, 3, 6, 7, 9, 4, 10, 5, 6}, 6},
	}

	for _, tt := range tests {
		result := lengthOfLIS(tt.nums)
		if result != tt.expected {
			t.Errorf("lengthOfLIS(%v) = %d, want %d", tt.nums, result, tt.expected)
		}
	}
}
