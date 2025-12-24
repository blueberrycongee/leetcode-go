package palindromelinkedlist

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

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		list     []int
		expected bool
	}{
		{
			name:     "示例1-回文",
			list:     []int{1, 2, 2, 1},
			expected: true,
		},
		{
			name:     "示例2-非回文",
			list:     []int{1, 2},
			expected: false,
		},
		{
			name:     "奇数长度回文",
			list:     []int{1, 2, 1},
			expected: true,
		},
		{
			name:     "单节点",
			list:     []int{1},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := sliceToList(tt.list)
			result := isPalindrome(head)
			if result != tt.expected {
				t.Errorf("isPalindrome(%v) = %v, want %v", tt.list, result, tt.expected)
			}
		})
	}
}
