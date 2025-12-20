package movezeroes

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "example 1",
			nums: []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "example 2",
			nums: []int{0},
			want: []int{0},
		},
		{
			name: "no zeroes",
			nums: []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "all zeroes",
			nums: []int{0, 0, 0},
			want: []int{0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// make a copy to avoid modifying the original test case data if we were reusing it,
			// though here it's fine.
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)

			moveZeroes(numsCopy)

			if !reflect.DeepEqual(numsCopy, tt.want) {
				t.Errorf("moveZeroes() result %v, want %v", numsCopy, tt.want)
			}
		})
	}
}
