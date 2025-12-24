package trappingrainwater

import "testing"

func TestTrap(t *testing.T) {
	tests := []struct {
		name     string
		height   []int
		expected int
	}{
		{
			name:     "示例1",
			height:   []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			expected: 6,
		},
		{
			name:     "示例2",
			height:   []int{4, 2, 0, 3, 2, 5},
			expected: 9,
		},
		{
			name:     "空数组",
			height:   []int{},
			expected: 0,
		},
		{
			name:     "单调递增",
			height:   []int{1, 2, 3, 4, 5},
			expected: 0,
		},
		{
			name:     "单调递减",
			height:   []int{5, 4, 3, 2, 1},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := trap(tt.height)
			if result != tt.expected {
				t.Errorf("trap(%v) = %d, want %d", tt.height, result, tt.expected)
			}
		})
	}
}
