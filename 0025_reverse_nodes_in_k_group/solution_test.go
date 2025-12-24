package reversenodesinkgroup

import "testing"

func sliceToList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, n := range nums {
		curr.Next = &ListNode{Val: n}
		curr = curr.Next
	}
	return dummy.Next
}

func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		name     string
		list     []int
		k        int
		expected []int
	}{
		{
			name:     "示例1",
			list:     []int{1, 2, 3, 4, 5},
			k:        2,
			expected: []int{2, 1, 4, 3, 5},
		},
		{
			name:     "示例2",
			list:     []int{1, 2, 3, 4, 5},
			k:        3,
			expected: []int{3, 2, 1, 4, 5},
		},
		{
			name:     "k等于长度",
			list:     []int{1, 2, 3},
			k:        3,
			expected: []int{3, 2, 1},
		},
		{
			name:     "k大于长度",
			list:     []int{1, 2},
			k:        3,
			expected: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := sliceToList(tt.list)
			result := reverseKGroup(head, tt.k)
			got := listToSlice(result)

			if len(got) != len(tt.expected) {
				t.Errorf("got %v, want %v", got, tt.expected)
				return
			}
			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("got %v, want %v", got, tt.expected)
					return
				}
			}
		})
	}
}
