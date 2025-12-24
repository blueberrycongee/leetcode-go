package slidingwindowmaximum

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "示例1",
			nums:     []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:        3,
			expected: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name:     "示例2",
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
		{
			name:     "k等于数组长度",
			nums:     []int{1, 2, 3},
			k:        3,
			expected: []int{3},
		},
		{
			name:     "递减数组",
			nums:     []int{5, 4, 3, 2, 1},
			k:        2,
			expected: []int{5, 4, 3, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxSlidingWindow(tt.nums, tt.k)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("maxSlidingWindow(%v, %d) = %v, want %v", tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}
