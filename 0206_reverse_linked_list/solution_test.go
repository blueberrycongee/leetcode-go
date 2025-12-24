package reverselinkedlist

import "testing"

// 辅助函数：数组转链表
func sliceToList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, n := range nums {
		curr.Next = &ListNode{Val: n}
		curr = curr.Next
	}
	return dummy.Next
}

// 辅助函数：链表转数组
func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func TestReverseList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "示例1",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "示例2",
			input:    []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "空链表",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "单节点",
			input:    []int{1},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := sliceToList(tt.input)
			result := reverseList(head)
			got := listToSlice(result)

			if len(got) != len(tt.expected) {
				t.Errorf("reverseList(%v) = %v, want %v", tt.input, got, tt.expected)
				return
			}
			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("reverseList(%v) = %v, want %v", tt.input, got, tt.expected)
					return
				}
			}
		})
	}
}
