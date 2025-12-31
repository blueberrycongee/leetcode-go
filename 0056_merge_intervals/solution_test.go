package mergeintervals

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  [][]int
	}{
		{
			name:      "示例1",
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expected:  [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "示例2",
			intervals: [][]int{{1, 4}, {4, 5}},
			expected:  [][]int{{1, 5}},
		},
		{
			name:      "单区间",
			intervals: [][]int{{1, 2}},
			expected:  [][]int{{1, 2}},
		},
		{
			name:      "完全重叠",
			intervals: [][]int{{1, 4}, {2, 3}},
			expected:  [][]int{{1, 4}},
		},
		{
			name:      "无重叠",
			intervals: [][]int{{1, 2}, {3, 4}, {5, 6}},
			expected:  [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := merge(tt.intervals)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
