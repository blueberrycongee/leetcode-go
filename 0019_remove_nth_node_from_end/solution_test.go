package removenthnodefromend

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

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name     string
		list     []int
		n        int
		expected []int
	}{
		{
			name:     "示例1",
			list:     []int{1, 2, 3, 4, 5},
			n:        2,
			expected: []int{1, 2, 3, 5},
		},
		{
			name:     "示例2-单节点",
			list:     []int{1},
			n:        1,
			expected: []int{},
		},
		{
			name:     "示例3",
			list:     []int{1, 2},
			n:        1,
			expected: []int{1},
		},
		{
			name:     "删除头节点",
			list:     []int{1, 2},
			n:        2,
			expected: []int{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := sliceToList(tt.list)
			result := removeNthFromEnd(head, tt.n)
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
