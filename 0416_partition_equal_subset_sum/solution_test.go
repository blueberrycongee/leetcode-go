package partitionequalsubsetsum

import "testing"

func TestCanPartition(t *testing.T) {
	tests := []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 5, 11, 5}, true},
		{[]int{1, 2, 3, 5}, false},
		{[]int{1, 2, 5}, false},
		{[]int{2, 2, 1, 1}, true},
	}

	for _, tt := range tests {
		result := canPartition(tt.nums)
		if result != tt.expected {
			t.Errorf("canPartition(%v) = %v, want %v", tt.nums, result, tt.expected)
		}
	}
}
